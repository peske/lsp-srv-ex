package lsp_srv_ex

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"github.com/peske/lsp-srv/lsp/protocol"
	"github.com/peske/lsp-srv/span"
	"github.com/peske/x-tools-internal/jsonrpc2"
	"go.uber.org/zap"
)

// Cache represents the cache.
type Cache struct {
	logger *zap.Logger

	rootUri  span.URI
	rootPath string

	mu    sync.RWMutex
	files map[span.URI]*file
}

func (c *Cache) getFile(uri span.URI) *file {
	c.mu.RLock()
	defer c.mu.RUnlock()

	if c.files == nil {
		return nil
	}
	return c.files[uri]
}

func (c *Cache) setFile(f *file) *file {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.files == nil {
		c.files = map[span.URI]*file{f.uri: f}
		return f
	}
	if prev := c.files[f.uri]; prev != nil {
		return prev
	}
	c.files[f.uri] = f
	return f
}

func (c *Cache) initialize(params *protocol.ParamInitialize, res *protocol.InitializeResult) (
	*protocol.InitializeResult, error) {
	c.rootUri = params.RootURI.SpanURI()
	c.rootPath = params.RootPath

	fs := make(map[span.URI]*file)
	c.loadFiles(c.rootPath, fs)

	c.mu.Lock()
	c.files = fs
	c.mu.Unlock()

	if res == nil {
		res = &protocol.InitializeResult{}
	}

	var tds *protocol.TextDocumentSyncOptions
	ok := false
	if res.Capabilities.TextDocumentSync != nil {
		tds, ok = res.Capabilities.TextDocumentSync.(*protocol.TextDocumentSyncOptions)
	}
	if !ok {
		tds = &protocol.TextDocumentSyncOptions{}
		res.Capabilities.TextDocumentSync = tds
	}
	tds.OpenClose = true
	tds.Change = protocol.Incremental
	tds.Save = &protocol.SaveOptions{IncludeText: false}

	return res, nil
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
		c.logger.Warn("didChange for a non-file uri", zap.String("URI", string(uri)))
		return
	}

	if len(params.ContentChanges) == 0 {
		err = fmt.Errorf("%w: didChange no content changes provided", jsonrpc2.ErrInternal)
		return
	}

	f := c.getFile(uri)

	// Check if the client sent the full content of the file.
	// We accept a full content change even if the server expected incremental changes.
	if len(params.ContentChanges) == 1 && params.ContentChanges[0].Range == nil &&
		params.ContentChanges[0].RangeLength == 0 {
		if f == nil {
			f = c.setFile(&file{parent: c, uri: uri})
			c.logger.Warn(fmt.Sprintf("didChange '%s' full content received.", uri))
		} else {
			c.logger.Warn(fmt.Sprintf("didChange '%s' full content received although the file exists.", uri))
		}
		f.setIdeContent([]byte(params.ContentChanges[0].Text), params.TextDocument.Version)
		return
	}

	if f == nil {
		err = fmt.Errorf("%w: file not found", jsonrpc2.ErrInternal)
		return
	}

	err = f.mergeChanges(params)
	return
}

func (c *Cache) didClose(params *protocol.DidCloseTextDocumentParams) (err error) {
	if params == nil {
		err = fmt.Errorf("%w: didClose params == nil", jsonrpc2.ErrInvalidParams)
		c.logger.Error("didClose", zap.Error(err))
		return
	}

	f := c.getFile(params.TextDocument.URI.SpanURI())
	if f == nil {
		c.logger.Warn("didClose unknown file", zap.String("URI", string(params.TextDocument.URI)))
	} else {
		f.closed()
	}
	return
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

	f := c.getFile(uri)
	if f == nil {
		c.logger.Warn("didOpen unknown file", zap.String("URI", string(uri)))
		f = c.setFile(&file{
			parent:     c,
			uri:        uri,
			languageID: params.TextDocument.LanguageID,
		})
	}

	f.setIdeContent([]byte(params.TextDocument.Text), params.TextDocument.Version)
	return
}

func (c *Cache) didSave(params *protocol.DidSaveTextDocumentParams) (err error) {
	if params == nil {
		err = fmt.Errorf("%w: didSave params == nil", jsonrpc2.ErrInvalidParams)
		c.logger.Error("didSave", zap.Error(err))
		return
	}

	f := c.getFile(params.TextDocument.URI.SpanURI())
	if f == nil {
		c.logger.Warn("didSave unknown file", zap.String("URI", string(params.TextDocument.URI)))
	} else {
		f.resetSavedContent()
	}
	return
}

// RootURI returns the `span.URI` of the root directory.
func (c *Cache) RootURI() span.URI {
	return c.rootUri
}

// RootPath returns the local absolute path of the root directory.
func (c *Cache) RootPath() string {
	return c.rootPath
}

// GetFiles returns the list of files kept in the cache.
func (c *Cache) GetFiles() []FileInfo {
	c.mu.RLock()
	defer c.mu.RUnlock()

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
	if f := c.getFile(uri); f != nil {
		return f.detach()
	}
	return nil
}

func (c *Cache) loadFiles(dir string, files map[span.URI]*file) {
	fds, err := os.ReadDir(dir)
	if err != nil {
		c.logger.Warn("loadFiles error", zap.Error(err))
		return
	}
	for _, fd := range fds {
		path := filepath.Join(dir, fd.Name())
		if fd.IsDir() {
			if fd.Name() != ".git" {
				c.loadFiles(path, files)
			}
		} else {
			f := &file{
				parent: c,
				uri:    span.URIFromPath(path),
				path:   path,
			}
			files[f.uri] = f
		}
	}
}
