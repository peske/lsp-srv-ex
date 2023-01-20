package lsp_srv_ex

import (
	"context"

	"github.com/peske/lsp-srv/lsp/protocol"
	"go.uber.org/zap"
)

// serverWrapper implements protocol.Server and wraps the actual server.
type serverWrapper struct {
	inner  protocol.Server
	helper *Helper
	cfg    *Config

	logger *zap.Logger
}

func NewServerWrapper(inner protocol.Server, helper *Helper, cfg *Config,
	lgr *zap.Logger) protocol.Server {
	if cfg == nil {
		cfg = &Config{}
	}
	return &serverWrapper{
		inner:  inner,
		helper: helper,
		cfg:    cfg,
		logger: lgr,
	}
}

func (s *serverWrapper) Progress(ctx context.Context, params *protocol.ProgressParams) error {
	s.logger.Debug("Progress", zap.Any("params", params))
	return s.inner.Progress(ctx, params)
}

func (s *serverWrapper) SetTrace(ctx context.Context, params *protocol.SetTraceParams) error {
	s.logger.Debug("SetTrace", zap.Any("params", params))
	return s.inner.SetTrace(ctx, params)
}

func (s *serverWrapper) IncomingCalls(ctx context.Context, params *protocol.CallHierarchyIncomingCallsParams) ([]protocol.CallHierarchyIncomingCall, error) {
	s.logger.Debug("IncomingCalls", zap.Any("params", params))
	return s.inner.IncomingCalls(ctx, params)
}

func (s *serverWrapper) OutgoingCalls(ctx context.Context, params *protocol.CallHierarchyOutgoingCallsParams) ([]protocol.CallHierarchyOutgoingCall, error) {
	s.logger.Debug("OutgoingCalls", zap.Any("params", params))
	return s.inner.OutgoingCalls(ctx, params)
}

func (s *serverWrapper) ResolveCodeAction(ctx context.Context, params *protocol.CodeAction) (*protocol.CodeAction, error) {
	s.logger.Debug("ResolveCodeAction", zap.Any("params", params))
	return s.inner.ResolveCodeAction(ctx, params)
}

func (s *serverWrapper) ResolveCodeLens(ctx context.Context, params *protocol.CodeLens) (*protocol.CodeLens, error) {
	s.logger.Debug("ResolveCodeLens", zap.Any("params", params))
	return s.inner.ResolveCodeLens(ctx, params)
}

func (s *serverWrapper) ResolveCompletionItem(ctx context.Context, params *protocol.CompletionItem) (*protocol.CompletionItem, error) {
	s.logger.Debug("ResolveCompletionItem", zap.Any("params", params))
	return s.inner.ResolveCompletionItem(ctx, params)
}

func (s *serverWrapper) ResolveDocumentLink(ctx context.Context, params *protocol.DocumentLink) (*protocol.DocumentLink, error) {
	s.logger.Debug("ResolveDocumentLink", zap.Any("params", params))
	return s.inner.ResolveDocumentLink(ctx, params)
}

func (s *serverWrapper) Exit(ctx context.Context) error {
	s.logger.Debug("Exit")
	return s.inner.Exit(ctx)
}

func (s *serverWrapper) Initialize(ctx context.Context, params *protocol.ParamInitialize) (*protocol.InitializeResult,
	error) {
	s.logger.Debug("Initialize", zap.Any("params", params))
	if err := s.helper.setStatus(Initializing); err != nil {
		return nil, err
	}
	res, err := s.inner.Initialize(ctx, params)
	if s.helper.Cache == nil || err != nil {
		return res, err
	}
	return s.helper.Cache.initialize(params, res)
}

func (s *serverWrapper) Initialized(ctx context.Context, params *protocol.InitializedParams) error {
	s.logger.Debug("Initialized", zap.Any("params", params))
	if err := s.helper.setStatus(Initialized); err != nil {
		return err
	}
	return s.inner.Initialized(ctx, params)
}

func (s *serverWrapper) Resolve(ctx context.Context, params *protocol.InlayHint) (*protocol.InlayHint, error) {
	s.logger.Debug("Resolve", zap.Any("params", params))
	return s.inner.Resolve(ctx, params)
}

func (s *serverWrapper) DidChangeNotebookDocument(ctx context.Context, params *protocol.DidChangeNotebookDocumentParams) error {
	s.logger.Debug("DidChangeNotebookDocument", zap.Any("params", params))
	return s.inner.DidChangeNotebookDocument(ctx, params)
}

func (s *serverWrapper) DidCloseNotebookDocument(ctx context.Context, params *protocol.DidCloseNotebookDocumentParams) error {
	s.logger.Debug("DidCloseNotebookDocument", zap.Any("params", params))
	return s.inner.DidCloseNotebookDocument(ctx, params)
}

func (s *serverWrapper) DidOpenNotebookDocument(ctx context.Context, params *protocol.DidOpenNotebookDocumentParams) error {
	s.logger.Debug("DidOpenNotebookDocument", zap.Any("params", params))
	return s.inner.DidOpenNotebookDocument(ctx, params)
}

func (s *serverWrapper) DidSaveNotebookDocument(ctx context.Context, params *protocol.DidSaveNotebookDocumentParams) error {
	s.logger.Debug("DidSaveNotebookDocument", zap.Any("params", params))
	return s.inner.DidSaveNotebookDocument(ctx, params)
}

func (s *serverWrapper) Shutdown(ctx context.Context) error {
	s.logger.Debug("Shutdown")
	if err := s.helper.setStatus(Shutdown); err != nil {
		return err
	}
	return s.inner.Shutdown(ctx)
}

func (s *serverWrapper) CodeAction(ctx context.Context, params *protocol.CodeActionParams) ([]protocol.CodeAction, error) {
	s.logger.Debug("CodeAction", zap.Any("params", params))
	return s.inner.CodeAction(ctx, params)
}

func (s *serverWrapper) CodeLens(ctx context.Context, params *protocol.CodeLensParams) ([]protocol.CodeLens, error) {
	s.logger.Debug("CodeLens", zap.Any("params", params))
	return s.inner.CodeLens(ctx, params)
}

func (s *serverWrapper) ColorPresentation(ctx context.Context, params *protocol.ColorPresentationParams) ([]protocol.ColorPresentation, error) {
	s.logger.Debug("ColorPresentation", zap.Any("params", params))
	return s.inner.ColorPresentation(ctx, params)
}

func (s *serverWrapper) Completion(ctx context.Context, params *protocol.CompletionParams) (*protocol.CompletionList, error) {
	s.logger.Debug("Completion", zap.Any("params", params))
	return s.inner.Completion(ctx, params)
}

func (s *serverWrapper) Declaration(ctx context.Context, params *protocol.DeclarationParams) (*protocol.Or_textDocument_declaration, error) {
	s.logger.Debug("Declaration", zap.Any("params", params))
	return s.inner.Declaration(ctx, params)
}

func (s *serverWrapper) Definition(ctx context.Context, params *protocol.DefinitionParams) ([]protocol.Location, error) {
	s.logger.Debug("Definition", zap.Any("params", params))
	return s.inner.Definition(ctx, params)
}

func (s *serverWrapper) Diagnostic(ctx context.Context, params *string) (*string, error) {
	s.logger.Debug("Diagnostic", zap.Any("params", params))
	return s.inner.Diagnostic(ctx, params)
}

func (s *serverWrapper) DidChange(ctx context.Context, params *protocol.DidChangeTextDocumentParams) error {
	s.logger.Debug("DidChange", zap.Any("params", params))
	if s.helper.Cache != nil {
		if err := s.helper.Cache.didChange(params); err != nil {
			return err
		}
	}
	return s.inner.DidChange(ctx, params)
}

func (s *serverWrapper) DidClose(ctx context.Context, params *protocol.DidCloseTextDocumentParams) error {
	s.logger.Debug("DidClose", zap.Any("params", params))
	if s.helper.Cache != nil {
		if err := s.helper.Cache.didClose(params); err != nil {
			return err
		}
	}
	return s.inner.DidClose(ctx, params)
}

func (s *serverWrapper) DidOpen(ctx context.Context, params *protocol.DidOpenTextDocumentParams) error {
	s.logger.Debug("DidOpen", zap.Any("params", params))
	if s.helper.Cache != nil {
		if err := s.helper.Cache.didOpen(params); err != nil {
			return err
		}
	}
	return s.inner.DidOpen(ctx, params)
}

func (s *serverWrapper) DidSave(ctx context.Context, params *protocol.DidSaveTextDocumentParams) error {
	s.logger.Debug("DidSave", zap.Any("params", params))
	if s.helper.Cache != nil {
		if err := s.helper.Cache.didSave(params); err != nil {
			return err
		}
	}
	return s.inner.DidSave(ctx, params)
}

func (s *serverWrapper) DocumentColor(ctx context.Context, params *protocol.DocumentColorParams) ([]protocol.ColorInformation, error) {
	s.logger.Debug("DocumentColor", zap.Any("params", params))
	return s.inner.DocumentColor(ctx, params)
}

func (s *serverWrapper) DocumentHighlight(ctx context.Context, params *protocol.DocumentHighlightParams) ([]protocol.DocumentHighlight, error) {
	s.logger.Debug("DocumentHighlight", zap.Any("params", params))
	return s.inner.DocumentHighlight(ctx, params)
}

func (s *serverWrapper) DocumentLink(ctx context.Context, params *protocol.DocumentLinkParams) ([]protocol.DocumentLink, error) {
	s.logger.Debug("DocumentLink", zap.Any("params", params))
	return s.inner.DocumentLink(ctx, params)
}

func (s *serverWrapper) DocumentSymbol(ctx context.Context, params *protocol.DocumentSymbolParams) ([]interface{}, error) {
	s.logger.Debug("DocumentSymbol", zap.Any("params", params))
	return s.inner.DocumentSymbol(ctx, params)
}

func (s *serverWrapper) FoldingRange(ctx context.Context, params *protocol.FoldingRangeParams) ([]protocol.FoldingRange, error) {
	s.logger.Debug("FoldingRange", zap.Any("params", params))
	return s.inner.FoldingRange(ctx, params)
}

func (s *serverWrapper) Formatting(ctx context.Context, params *protocol.DocumentFormattingParams) ([]protocol.TextEdit, error) {
	s.logger.Debug("Formatting", zap.Any("params", params))
	return s.inner.Formatting(ctx, params)
}

func (s *serverWrapper) Hover(ctx context.Context, params *protocol.HoverParams) (*protocol.Hover, error) {
	s.logger.Debug("Hover", zap.Any("params", params))
	return s.inner.Hover(ctx, params)
}

func (s *serverWrapper) Implementation(ctx context.Context, params *protocol.ImplementationParams) ([]protocol.Location, error) {
	s.logger.Debug("Implementation", zap.Any("params", params))
	return s.inner.Implementation(ctx, params)
}

func (s *serverWrapper) InlayHint(ctx context.Context, params *protocol.InlayHintParams) ([]protocol.InlayHint, error) {
	s.logger.Debug("InlayHint", zap.Any("params", params))
	return s.inner.InlayHint(ctx, params)
}

func (s *serverWrapper) InlineValue(ctx context.Context, params *protocol.InlineValueParams) ([]protocol.InlineValue, error) {
	s.logger.Debug("InlineValue", zap.Any("params", params))
	return s.inner.InlineValue(ctx, params)
}

func (s *serverWrapper) LinkedEditingRange(ctx context.Context, params *protocol.LinkedEditingRangeParams) (*protocol.LinkedEditingRanges, error) {
	s.logger.Debug("LinkedEditingRange", zap.Any("params", params))
	return s.inner.LinkedEditingRange(ctx, params)
}

func (s *serverWrapper) Moniker(ctx context.Context, params *protocol.MonikerParams) ([]protocol.Moniker, error) {
	s.logger.Debug("Moniker", zap.Any("params", params))
	return s.inner.Moniker(ctx, params)
}

func (s *serverWrapper) OnTypeFormatting(ctx context.Context, params *protocol.DocumentOnTypeFormattingParams) ([]protocol.TextEdit, error) {
	s.logger.Debug("OnTypeFormatting", zap.Any("params", params))
	return s.inner.OnTypeFormatting(ctx, params)
}

func (s *serverWrapper) PrepareCallHierarchy(ctx context.Context, params *protocol.CallHierarchyPrepareParams) ([]protocol.CallHierarchyItem, error) {
	s.logger.Debug("PrepareCallHierarchy", zap.Any("params", params))
	return s.inner.PrepareCallHierarchy(ctx, params)
}

func (s *serverWrapper) PrepareRename(ctx context.Context, params *protocol.PrepareRenameParams) (*protocol.PrepareRename2Gn, error) {
	s.logger.Debug("PrepareRename", zap.Any("params", params))
	return s.inner.PrepareRename(ctx, params)
}

func (s *serverWrapper) PrepareTypeHierarchy(ctx context.Context, params *protocol.TypeHierarchyPrepareParams) ([]protocol.TypeHierarchyItem, error) {
	s.logger.Debug("PrepareTypeHierarchy", zap.Any("params", params))
	return s.inner.PrepareTypeHierarchy(ctx, params)
}

func (s *serverWrapper) RangeFormatting(ctx context.Context, params *protocol.DocumentRangeFormattingParams) ([]protocol.TextEdit, error) {
	s.logger.Debug("RangeFormatting", zap.Any("params", params))
	return s.inner.RangeFormatting(ctx, params)
}

func (s *serverWrapper) References(ctx context.Context, params *protocol.ReferenceParams) ([]protocol.Location, error) {
	s.logger.Debug("References", zap.Any("params", params))
	return s.inner.References(ctx, params)
}

func (s *serverWrapper) Rename(ctx context.Context, params *protocol.RenameParams) (*protocol.WorkspaceEdit, error) {
	s.logger.Debug("Rename", zap.Any("params", params))
	return s.inner.Rename(ctx, params)
}

func (s *serverWrapper) SelectionRange(ctx context.Context, params *protocol.SelectionRangeParams) ([]protocol.SelectionRange, error) {
	s.logger.Debug("SelectionRange", zap.Any("params", params))
	return s.inner.SelectionRange(ctx, params)
}

func (s *serverWrapper) SemanticTokensFull(ctx context.Context, params *protocol.SemanticTokensParams) (*protocol.SemanticTokens, error) {
	s.logger.Debug("SemanticTokensFull", zap.Any("params", params))
	return s.inner.SemanticTokensFull(ctx, params)
}

func (s *serverWrapper) SemanticTokensFullDelta(ctx context.Context, params *protocol.SemanticTokensDeltaParams) (interface{}, error) {
	s.logger.Debug("SemanticTokensFullDelta", zap.Any("params", params))
	return s.inner.SemanticTokensFullDelta(ctx, params)
}

func (s *serverWrapper) SemanticTokensRange(ctx context.Context, params *protocol.SemanticTokensRangeParams) (*protocol.SemanticTokens, error) {
	s.logger.Debug("SemanticTokensRange", zap.Any("params", params))
	return s.inner.SemanticTokensRange(ctx, params)
}

func (s *serverWrapper) SignatureHelp(ctx context.Context, params *protocol.SignatureHelpParams) (*protocol.SignatureHelp, error) {
	s.logger.Debug("SignatureHelp", zap.Any("params", params))
	return s.inner.SignatureHelp(ctx, params)
}

func (s *serverWrapper) TypeDefinition(ctx context.Context, params *protocol.TypeDefinitionParams) ([]protocol.Location, error) {
	s.logger.Debug("TypeDefinition", zap.Any("params", params))
	return s.inner.TypeDefinition(ctx, params)
}

func (s *serverWrapper) WillSave(ctx context.Context, params *protocol.WillSaveTextDocumentParams) error {
	s.logger.Debug("WillSave", zap.Any("params", params))
	return s.inner.WillSave(ctx, params)
}

func (s *serverWrapper) WillSaveWaitUntil(ctx context.Context, params *protocol.WillSaveTextDocumentParams) ([]protocol.TextEdit, error) {
	s.logger.Debug("WillSaveWaitUntil", zap.Any("params", params))
	return s.inner.WillSaveWaitUntil(ctx, params)
}

func (s *serverWrapper) Subtypes(ctx context.Context, params *protocol.TypeHierarchySubtypesParams) ([]protocol.TypeHierarchyItem, error) {
	s.logger.Debug("Subtypes", zap.Any("params", params))
	return s.inner.Subtypes(ctx, params)
}

func (s *serverWrapper) Supertypes(ctx context.Context, params *protocol.TypeHierarchySupertypesParams) ([]protocol.TypeHierarchyItem, error) {
	s.logger.Debug("Supertypes", zap.Any("params", params))
	return s.inner.Supertypes(ctx, params)
}

func (s *serverWrapper) WorkDoneProgressCancel(ctx context.Context, params *protocol.WorkDoneProgressCancelParams) error {
	s.logger.Debug("WorkDoneProgressCancel", zap.Any("params", params))
	return s.inner.WorkDoneProgressCancel(ctx, params)
}

func (s *serverWrapper) DiagnosticWorkspace(ctx context.Context, params *protocol.WorkspaceDiagnosticParams) (*protocol.WorkspaceDiagnosticReport, error) {
	s.logger.Debug("DiagnosticWorkspace", zap.Any("params", params))
	return s.inner.DiagnosticWorkspace(ctx, params)
}

func (s *serverWrapper) DiagnosticRefresh(ctx context.Context) error {
	s.logger.Debug("DiagnosticRefresh")
	return s.inner.DiagnosticRefresh(ctx)
}

func (s *serverWrapper) DidChangeConfiguration(ctx context.Context, params *protocol.DidChangeConfigurationParams) error {
	s.logger.Debug("DidChangeConfiguration", zap.Any("params", params))
	return s.inner.DidChangeConfiguration(ctx, params)
}

func (s *serverWrapper) DidChangeWatchedFiles(ctx context.Context, params *protocol.DidChangeWatchedFilesParams) error {
	s.logger.Debug("DidChangeWatchedFiles", zap.Any("params", params))
	return s.inner.DidChangeWatchedFiles(ctx, params)
}

func (s *serverWrapper) DidChangeWorkspaceFolders(ctx context.Context, params *protocol.DidChangeWorkspaceFoldersParams) error {
	s.logger.Debug("DidChangeWorkspaceFolders", zap.Any("params", params))
	return s.inner.DidChangeWorkspaceFolders(ctx, params)
}

func (s *serverWrapper) DidCreateFiles(ctx context.Context, params *protocol.CreateFilesParams) error {
	s.logger.Debug("DidCreateFiles", zap.Any("params", params))
	return s.inner.DidCreateFiles(ctx, params)
}

func (s *serverWrapper) DidDeleteFiles(ctx context.Context, params *protocol.DeleteFilesParams) error {
	s.logger.Debug("DidDeleteFiles", zap.Any("params", params))
	return s.inner.DidDeleteFiles(ctx, params)
}

func (s *serverWrapper) DidRenameFiles(ctx context.Context, params *protocol.RenameFilesParams) error {
	s.logger.Debug("DidRenameFiles", zap.Any("params", params))
	return s.inner.DidRenameFiles(ctx, params)
}

func (s *serverWrapper) ExecuteCommand(ctx context.Context, params *protocol.ExecuteCommandParams) (interface{}, error) {
	s.logger.Debug("ExecuteCommand", zap.Any("params", params))
	return s.inner.ExecuteCommand(ctx, params)
}

func (s *serverWrapper) InlayHintRefresh(ctx context.Context) error {
	s.logger.Debug("InlayHintRefresh")
	return s.inner.InlayHintRefresh(ctx)
}

func (s *serverWrapper) InlineValueRefresh(ctx context.Context) error {
	s.logger.Debug("InlineValueRefresh")
	return s.inner.InlineValueRefresh(ctx)
}

func (s *serverWrapper) SemanticTokensRefresh(ctx context.Context) error {
	s.logger.Debug("SemanticTokensRefresh")
	return s.inner.SemanticTokensRefresh(ctx)
}

func (s *serverWrapper) Symbol(ctx context.Context, params *protocol.WorkspaceSymbolParams) ([]protocol.SymbolInformation, error) {
	s.logger.Debug("Symbol", zap.Any("params", params))
	return s.inner.Symbol(ctx, params)
}

func (s *serverWrapper) WillCreateFiles(ctx context.Context, params *protocol.CreateFilesParams) (*protocol.WorkspaceEdit, error) {
	s.logger.Debug("WillCreateFiles", zap.Any("params", params))
	return s.inner.WillCreateFiles(ctx, params)
}

func (s *serverWrapper) WillDeleteFiles(ctx context.Context, params *protocol.DeleteFilesParams) (*protocol.WorkspaceEdit, error) {
	s.logger.Debug("WillDeleteFiles", zap.Any("params", params))
	return s.inner.WillDeleteFiles(ctx, params)
}

func (s *serverWrapper) WillRenameFiles(ctx context.Context, params *protocol.RenameFilesParams) (*protocol.WorkspaceEdit, error) {
	s.logger.Debug("WillRenameFiles", zap.Any("params", params))
	return s.inner.WillRenameFiles(ctx, params)
}

func (s *serverWrapper) ResolveWorkspaceSymbol(ctx context.Context, params *protocol.WorkspaceSymbol) (*protocol.WorkspaceSymbol, error) {
	s.logger.Debug("ResolveWorkspaceSymbol", zap.Any("params", params))
	return s.inner.ResolveWorkspaceSymbol(ctx, params)
}

func (s *serverWrapper) NonstandardRequest(ctx context.Context, method string, params interface{}) (interface{}, error) {
	s.logger.Debug("NonstandardRequest", zap.String("method", method), zap.Any("params", params))
	return s.inner.NonstandardRequest(ctx, method, params)
}
