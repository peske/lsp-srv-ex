package lsp_srv_ex

import (
	"context"

	"github.com/peske/lsp-srv/lsp/protocol"
	"go.uber.org/zap"
)

// clientWrapper implements protocol.ClientCloser interface and wraps the actual client.
type clientWrapper struct {
	inner  protocol.ClientCloser
	helper *Helper
	logger *zap.Logger
}

func (c *clientWrapper) DiagnosticRefresh(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func (c *clientWrapper) InlayHintRefresh(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func (c *clientWrapper) InlineValueRefresh(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func (c *clientWrapper) SemanticTokensRefresh(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func NewClientWrapper(inner protocol.ClientCloser, helper *Helper, lgr *zap.Logger) protocol.ClientCloser {
	return &clientWrapper{
		inner:  inner,
		helper: helper,
		logger: lgr,
	}
}

func (c *clientWrapper) LogTrace(ctx context.Context, params *protocol.LogTraceParams) error {
	c.logger.Debug("LogTrace", zap.Any("params", params))
	return c.inner.LogTrace(ctx, params)
}

func (c *clientWrapper) Progress(ctx context.Context, params *protocol.ProgressParams) error {
	c.logger.Debug("Progress", zap.Any("params", params))
	return c.inner.Progress(ctx, params)
}

func (c *clientWrapper) RegisterCapability(ctx context.Context, params *protocol.RegistrationParams) error {
	c.logger.Debug("RegisterCapability", zap.Any("params", params))
	return c.inner.RegisterCapability(ctx, params)
}

func (c *clientWrapper) UnregisterCapability(ctx context.Context, params *protocol.UnregistrationParams) error {
	c.logger.Debug("UnregisterCapability", zap.Any("params", params))
	return c.inner.UnregisterCapability(ctx, params)
}

func (c *clientWrapper) Event(ctx context.Context, params *interface{}) error {
	c.logger.Debug("Event", zap.Any("params", params))
	return c.inner.Event(ctx, params)
}

func (c *clientWrapper) PublishDiagnostics(ctx context.Context, params *protocol.PublishDiagnosticsParams) error {
	c.logger.Debug("PublishDiagnostics", zap.Any("params", params))
	return c.inner.PublishDiagnostics(ctx, params)
}

func (c *clientWrapper) LogMessage(ctx context.Context, params *protocol.LogMessageParams) error {
	c.logger.Debug("LogMessage", zap.Any("params", params))
	return c.inner.LogMessage(ctx, params)
}

func (c *clientWrapper) ShowDocument(ctx context.Context, params *protocol.ShowDocumentParams) (*protocol.ShowDocumentResult, error) {
	c.logger.Debug("ShowDocument", zap.Any("params", params))
	return c.inner.ShowDocument(ctx, params)
}

func (c *clientWrapper) ShowMessage(ctx context.Context, params *protocol.ShowMessageParams) error {
	c.logger.Debug("ShowMessage", zap.Any("params", params))
	return c.inner.ShowMessage(ctx, params)
}

func (c *clientWrapper) ShowMessageRequest(ctx context.Context, params *protocol.ShowMessageRequestParams) (*protocol.MessageActionItem, error) {
	c.logger.Debug("ShowMessageRequest", zap.Any("params", params))
	return c.inner.ShowMessageRequest(ctx, params)
}

func (c *clientWrapper) WorkDoneProgressCreate(ctx context.Context, params *protocol.WorkDoneProgressCreateParams) error {
	c.logger.Debug("WorkDoneProgressCreate", zap.Any("params", params))
	return c.inner.WorkDoneProgressCreate(ctx, params)
}

func (c *clientWrapper) ApplyEdit(ctx context.Context, params *protocol.ApplyWorkspaceEditParams) (*protocol.ApplyWorkspaceEditResult, error) {
	c.logger.Debug("ApplyEdit", zap.Any("params", params))
	return c.inner.ApplyEdit(ctx, params)
}

func (c *clientWrapper) CodeLensRefresh(ctx context.Context) error {
	c.logger.Debug("CodeLensRefresh")
	return c.inner.CodeLensRefresh(ctx)
}

func (c *clientWrapper) Configuration(ctx context.Context, params *protocol.ParamConfiguration) ([]protocol.LSPAny, error) {
	c.logger.Debug("Configuration", zap.Any("params", params))
	return c.inner.Configuration(ctx, params)
}

func (c *clientWrapper) WorkspaceFolders(ctx context.Context) ([]protocol.WorkspaceFolder, error) {
	c.logger.Debug("WorkspaceFolders")
	return c.inner.WorkspaceFolders(ctx)
}

func (c *clientWrapper) Close() error {
	c.logger.Debug("Close")
	return c.inner.Close()
}
