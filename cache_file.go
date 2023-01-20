package lsp_srv_ex

import (
	"bytes"
	"fmt"
	"os"
	"sync"

	"github.com/peske/lsp-srv/lsp/protocol"
	"github.com/peske/lsp-srv/span"
	"github.com/peske/x-tools-internal/jsonrpc2"
)

// FileInfo interface defines the basic file info.
type FileInfo interface {
	// URI is the `span.URI` of the file.
	URI() span.URI
	// Path is the local path relative to the root path.
	Path() string
	// IsOpened is `true` if the file is open in the IDE, `false` if it isn't.
	IsOpened() bool
}

// File structure represents a _detached_ file, meaning
// that it isn't updated after the structure is created.
type File struct {
	file    *file
	Content []byte
	Version int32
}

// URI returns `span.URI` of the file.
func (f *File) URI() span.URI {
	return f.file.URI()
}

// Path returns the local path relative to the root path.
func (f *File) Path() string {
	return f.file.Path()
}

// IsOpened returns `true` if the file is open in the IDE, `false` if it isn't.
func (f *File) IsOpened() bool {
	return f.file.IsOpened()
}

// ChangedMeanwhile checks if the original file is changed since
// this detached `File` instance is created.
func (f *File) ChangedMeanwhile() bool {
	f.file.mu.RLock()
	defer f.file.mu.RUnlock()

	return f.file.version != f.Version
}

// GetSavedContent returns the saved content of the file.
// `forceRead`: if set to `true` forces reading the content
// from the disk, even if it's cached.
func (f *File) GetSavedContent(forceRead bool) ([]byte, error) {
	return f.file.getSavedContent(forceRead)
}

// file represents a file in cache.
type file struct {
	parent *Cache

	uri        span.URI
	path       string
	languageID string

	mu           sync.RWMutex // Content lock, protects the following fields
	savedContent []byte
	ideContent   []byte
	version      int32
}

// URI returns `span.URI` of the file.
func (f *file) URI() span.URI {
	return f.uri
}

// Path returns the local path relative to the root path.
func (f *file) Path() string {
	n := f.path
	if n == "" {
		n = f.uri.Filename()
		f.path = n
	}
	return n
}

// IsOpened returns `true` if the file is open in the IDE, `false` if it isn't.
func (f *file) IsOpened() bool {
	f.mu.RLock()
	defer f.mu.RUnlock()
	return f.ideContent != nil
}

func (f *file) detach() *File {
	f.mu.RLock()
	defer f.mu.RUnlock()

	df := &File{
		file:    f,
		Version: f.version,
	}

	if f.ideContent != nil {
		df.Content = make([]byte, len(f.ideContent), len(f.ideContent))
		copy(df.Content, f.ideContent)
	}

	return df
}

func (f *file) resetSavedContent() {
	f.mu.Lock()
	f.savedContent = nil
	f.mu.Unlock()
}

func (f *file) getSavedContent(forceRead bool) ([]byte, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	if forceRead {
		f.savedContent = nil
	}
	c, err := f.getSavedContentLocked()
	if err != nil {
		return nil, err
	}
	cp := make([]byte, len(c), len(c))
	copy(cp, c)
	return cp, nil
}

func (f *file) getSavedContentLocked() ([]byte, error) {
	if f.savedContent == nil {
		if c, err := os.ReadFile(f.Path()); err == nil {
			f.savedContent = c
		} else {
			return nil, err
		}
	}
	return f.savedContent, nil
}

func (f *file) mergeChanges(params *protocol.DidChangeTextDocumentParams) error {
	f.mu.Lock()
	defer f.mu.Unlock()

	content := f.ideContent
	for _, cc := range params.ContentChanges {
		// TODO(adonovan): refactor to use diff.Apply, which is robust w.r.t.
		// out-of-order or overlapping changes---and much more efficient.

		// Make sure to update column mapper along with the content.
		m := protocol.NewMapper(f.uri, content)
		if cc.Range == nil {
			return fmt.Errorf("%w: didChange unexpected nil range for change", jsonrpc2.ErrInternal)
		}
		spn, err := m.RangeSpan(*cc.Range)
		if err != nil {
			return err
		}
		start, end := spn.Start().Offset(), spn.End().Offset()
		if end < start {
			return fmt.Errorf("%w: invalid range for content change", jsonrpc2.ErrInternal)
		}
		var buf bytes.Buffer
		buf.Write(content[:start])
		buf.WriteString(cc.Text)
		buf.Write(content[end:])
		content = buf.Bytes()
	}

	f.ideContent = content
	f.version = params.TextDocument.Version

	return nil
}

func (f *file) setIdeContent(content []byte, version int32) {
	f.mu.Lock()
	defer f.mu.Unlock()

	f.ideContent = content
	f.version = version
}

func (f *file) closed() {
	f.mu.Lock()
	f.ideContent = nil
	f.mu.Unlock()
}
