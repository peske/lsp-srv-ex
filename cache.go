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

// FileInfo interface defines the basic file info.
type FileInfo interface {
	URI() span.URI
	Name() string
}

type File struct {
	file    *file
	Content []byte
	Version int32
}

// URI returns `span.URI` of the file.
func (f *File) URI() span.URI {
	return f.file.URI()
}

// Name returns the name of the file.
func (f *File) Name() string {
	return f.file.Name()
}

// ChangedMeanwhile checks if the original file in the cache is changed since
// this `File` instance is created (and detached).
func (f *File) ChangedMeanwhile() bool {
	f.file.parent.mu.Lock()
	defer f.file.parent.mu.Unlock()

	return f.file.version != f.Version
}

// file represents a file in cache.
type file struct {
	parent *Cache

	uri        span.URI
	name       string
	content    []byte
	version    int32
	languageID string

	buffer []*protocol.TextDocumentContentChangeEvent
}

// URI returns `span.URI` of the file.
func (f *file) URI() span.URI {
	return f.uri
}

// Name returns the name of the file.
func (f *file) Name() string {
	n := f.name
	if n == "" {
		n = f.uri.Filename()
		f.name = n
	}
	return n
}

func (f *file) toDetachedLocked() *File {
	df := &File{
		file:    f,
		Content: make([]byte, len(f.content), len(f.content)),
		Version: f.version,
	}

	copy(df.Content, f.content)

	return df
}

// Cache represents the cache.
type Cache struct {
	logger *zap.Logger

	mu    sync.Mutex
	files map[string]*file
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

	c.files[string(uri)] = &file{
		parent:     c,
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
		if f == nil {
			f = &file{parent: c, uri: uri}
			c.files[string(uri)] = f
			c.logger.Warn(fmt.Sprintf("didChange '%s' full content received.", uri))
		} else {
			c.logger.Warn(fmt.Sprintf("didChange '%s' full content received although the file exists.", uri))
		}
		f.content = []byte(params.ContentChanges[0].Text)
		f.version = params.TextDocument.Version
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
func (c *Cache) GetFiles() []FileInfo {
	c.mu.Lock()
	defer c.mu.Unlock()

	fs := make([]FileInfo, 0, len(c.files))
	for _, f := range c.files {
		fs = append(fs, f)
	}
	return fs
}

// GetFile returns `*File` pointer specified by `uri`.
// Note that the returned instance is detached from the cache, meaning that it
// won't be updated with any changes that may arrive after creating it.
func (c *Cache) GetFile(uri span.URI) *File {
	c.mu.Lock()
	defer c.mu.Unlock()

	if f := c.files[string(uri)]; f == nil {
		return nil
	} else {
		return f.toDetachedLocked()
	}
}
