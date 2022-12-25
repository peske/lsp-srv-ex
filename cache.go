package lsp_srv_ex

import (
	"bytes"
	"fmt"
	"sync"

	"github.com/peske/lsp-srv/lsp/protocol"
	"github.com/peske/lsp-srv/span"
	"github.com/peske/x-tools-internal/jsonrpc2"
	"go.uber.org/zap"
)

type File struct {
	uri        span.URI
	name       string
	content    []byte
	version    int32
	languageID string

	buffer []*protocol.TextDocumentContentChangeEvent
}

func (f *File) URI() span.URI {
	return f.uri
}

func (f *File) Name() string {
	return f.name
}

type Cache struct {
	logger *zap.Logger

	mu    sync.Mutex
	files map[string]*File
}

func (c *Cache) didOpen(params *protocol.DidOpenTextDocumentParams) (err error) {
	if params == nil {
		err = fmt.Errorf("%w: didOpen params == nil", jsonrpc2.ErrInvalidParams)
		c.logger.Error("didOpen", zap.Error(err))
		return
	}

	uri := params.TextDocument.URI.SpanURI()
	if !uri.IsFile() {
		return
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	c.files[string(uri)] = &File{
		uri:        uri,
		content:    []byte(params.TextDocument.Text),
		version:    params.TextDocument.Version,
		languageID: params.TextDocument.LanguageID,
	}
	return
}

func (c *Cache) didClose(params *protocol.DidCloseTextDocumentParams) (err error) {
	if params == nil {
		err = fmt.Errorf("%w: didClose params == nil", jsonrpc2.ErrInvalidParams)
		c.logger.Error("didClose", zap.Error(err))
		return
	}

	uri := params.TextDocument.URI.SpanURI()

	c.mu.Lock()
	delete(c.files, string(uri))
	c.mu.Unlock()
	return
}

func (c *Cache) didChange(params *protocol.DidChangeTextDocumentParams) (err error) {
	defer func() {
		if err != nil {
			c.logger.Error("didChange", zap.Error(err))
		}
	}()

	if params == nil {
		err = fmt.Errorf("%w: didChange params == nil", jsonrpc2.ErrInvalidParams)
		return
	}

	uri := params.TextDocument.URI.SpanURI()
	if !uri.IsFile() {
		return
	}

	if len(params.ContentChanges) == 0 {
		err = fmt.Errorf("%w: didChange no content changes provided", jsonrpc2.ErrInternal)
		return
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	f := c.files[string(uri)]

	// Check if the client sent the full content of the file.
	// We accept a full content change even if the server expected incremental changes.
	if len(params.ContentChanges) == 1 && params.ContentChanges[0].Range == nil &&
		params.ContentChanges[0].RangeLength == 0 {
		languageID := ""
		if f == nil {
			c.logger.Warn(fmt.Sprintf("didChange '%s' full content received.", uri))
		} else {
			languageID = f.languageID
			c.logger.Warn(fmt.Sprintf("didChange '%s' full content received although the file exists.", uri))
		}
		c.files[string(uri)] = &File{
			uri:        uri,
			content:    []byte(params.ContentChanges[0].Text),
			version:    params.TextDocument.Version,
			languageID: languageID,
		}
		return
	}

	if f == nil {
		err = fmt.Errorf("%w: file not found", jsonrpc2.ErrInternal)
		return
	}

	content := f.content
	for _, cc := range params.ContentChanges {
		// TODO(adonovan): refactor to use diff.Apply, which is robust w.r.t.
		// out-of-order or overlapping changes---and much more efficient.

		// Make sure to update column mapper along with the content.
		m := protocol.NewColumnMapper(uri, content)
		if cc.Range == nil {
			err = fmt.Errorf("%w: didChange unexpected nil range for change", jsonrpc2.ErrInternal)
			return
		}
		var spn span.Span
		spn, err = m.RangeSpan(*cc.Range)
		if err != nil {
			return
		}
		start, end := spn.Start().Offset(), spn.End().Offset()
		if end < start {
			err = fmt.Errorf("%w: invalid range for content change", jsonrpc2.ErrInternal)
			return
		}
		var buf bytes.Buffer
		buf.Write(content[:start])
		buf.WriteString(cc.Text)
		buf.Write(content[end:])
		content = buf.Bytes()
	}

	f.content = content
	f.version = params.TextDocument.Version

	return
}

// GetFiles returns the list of files kept in the cache.
func (c *Cache) GetFiles() []*File {
	c.mu.Lock()
	defer c.mu.Unlock()

	fs := make([]*File, 0, len(c.files))
	for _, f := range c.files {
		fs = append(fs, f)
	}
	return fs
}

func (c *Cache) GetFileContent(uri span.URI) (string, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	f := c.files[string(uri)]
	if f == nil {
		return "", false
	}

	return string(f.content), true
}
