package internal

import (
	"context"

	"github.com/peske/lsp-srv/lsp/protocol"
)

// clientWrapper implements protocol.ClientCloser interface and wraps the actual client.
type clientWrapper struct {
	inner protocol.ClientCloser
}

func NewClientWrapper(inner protocol.ClientCloser) protocol.ClientCloser {
	return &clientWrapper{inner: inner}
}

func (c *clientWrapper) LogTrace(ctx context.Context, params *protocol.LogTraceParams) error {
	return c.inner.LogTrace(ctx, params)
}

func (c *clientWrapper) Progress(ctx context.Context, params *protocol.ProgressParams) error {
	return c.inner.Progress(ctx, params)
}

func (c *clientWrapper) RegisterCapability(ctx context.Context, params *protocol.RegistrationParams) error {
	return c.inner.RegisterCapability(ctx, params)
}

func (c *clientWrapper) UnregisterCapability(ctx context.Context, params *protocol.UnregistrationParams) error {
	return c.inner.UnregisterCapability(ctx, params)
}

func (c *clientWrapper) Event(ctx context.Context, params *interface{}) error {
	return c.inner.Event(ctx, params)
}

func (c *clientWrapper) PublishDiagnostics(ctx context.Context, params *protocol.PublishDiagnosticsParams) error {
	return c.inner.PublishDiagnostics(ctx, params)
}

func (c *clientWrapper) LogMessage(ctx context.Context, params *protocol.LogMessageParams) error {
	return c.inner.LogMessage(ctx, params)
}

func (c *clientWrapper) ShowDocument(ctx context.Context, params *protocol.ShowDocumentParams) (*protocol.ShowDocumentResult, error) {
	return c.inner.ShowDocument(ctx, params)
}

func (c *clientWrapper) ShowMessage(ctx context.Context, params *protocol.ShowMessageParams) error {
	return c.inner.ShowMessage(ctx, params)
}

func (c *clientWrapper) ShowMessageRequest(ctx context.Context, params *protocol.ShowMessageRequestParams) (*protocol.MessageActionItem, error) {
	return c.inner.ShowMessageRequest(ctx, params)
}

func (c *clientWrapper) WorkDoneProgressCreate(ctx context.Context, params *protocol.WorkDoneProgressCreateParams) error {
	return c.inner.WorkDoneProgressCreate(ctx, params)
}

func (c *clientWrapper) ApplyEdit(ctx context.Context, params *protocol.ApplyWorkspaceEditParams) (*protocol.ApplyWorkspaceEditResult, error) {
	return c.inner.ApplyEdit(ctx, params)
}

func (c *clientWrapper) CodeLensRefresh(ctx context.Context) error {
	return c.inner.CodeLensRefresh(ctx)
}

func (c *clientWrapper) Configuration(ctx context.Context, params *protocol.ParamConfiguration) ([]protocol.LSPAny, error) {
	return c.inner.Configuration(ctx, params)
}

func (c *clientWrapper) WorkspaceFolders(ctx context.Context) ([]protocol.WorkspaceFolder, error) {
	return c.inner.WorkspaceFolders(ctx)
}

func (c *clientWrapper) Close() error {
	return c.inner.Close()
}
