package internal

import (
	"context"

	"github.com/peske/lsp-srv/lsp/protocol"
)

// serverWrapper implements protocol.Server and wraps the actual server.
type serverWrapper struct {
	inner protocol.Server
}

func NewServerWrapper(inner protocol.Server) protocol.Server {
	return &serverWrapper{inner: inner}
}

func (s *serverWrapper) Progress(ctx context.Context, params *protocol.ProgressParams) error {
	return s.inner.Progress(ctx, params)
}

func (s *serverWrapper) SetTrace(ctx context.Context, params *protocol.SetTraceParams) error {
	return s.inner.SetTrace(ctx, params)
}

func (s *serverWrapper) IncomingCalls(ctx context.Context, params *protocol.CallHierarchyIncomingCallsParams) ([]protocol.CallHierarchyIncomingCall, error) {
	return s.inner.IncomingCalls(ctx, params)
}

func (s *serverWrapper) OutgoingCalls(ctx context.Context, params *protocol.CallHierarchyOutgoingCallsParams) ([]protocol.CallHierarchyOutgoingCall, error) {
	return s.inner.OutgoingCalls(ctx, params)
}

func (s *serverWrapper) ResolveCodeAction(ctx context.Context, params *protocol.CodeAction) (*protocol.CodeAction, error) {
	return s.inner.ResolveCodeAction(ctx, params)
}

func (s *serverWrapper) ResolveCodeLens(ctx context.Context, params *protocol.CodeLens) (*protocol.CodeLens, error) {
	return s.inner.ResolveCodeLens(ctx, params)
}

func (s *serverWrapper) ResolveCompletionItem(ctx context.Context, params *protocol.CompletionItem) (*protocol.CompletionItem, error) {
	return s.inner.ResolveCompletionItem(ctx, params)
}

func (s *serverWrapper) ResolveDocumentLink(ctx context.Context, params *protocol.DocumentLink) (*protocol.DocumentLink, error) {
	return s.inner.ResolveDocumentLink(ctx, params)
}

func (s *serverWrapper) Exit(ctx context.Context) error {
	return s.inner.Exit(ctx)
}

func (s *serverWrapper) Initialize(ctx context.Context, params *protocol.ParamInitialize) (*protocol.InitializeResult, error) {
	return s.inner.Initialize(ctx, params)
}

func (s *serverWrapper) Initialized(ctx context.Context, params *protocol.InitializedParams) error {
	return s.inner.Initialized(ctx, params)
}

func (s *serverWrapper) Resolve(ctx context.Context, params *protocol.InlayHint) (*protocol.InlayHint, error) {
	return s.inner.Resolve(ctx, params)
}

func (s *serverWrapper) DidChangeNotebookDocument(ctx context.Context, params *protocol.DidChangeNotebookDocumentParams) error {
	return s.inner.DidChangeNotebookDocument(ctx, params)
}

func (s *serverWrapper) DidCloseNotebookDocument(ctx context.Context, params *protocol.DidCloseNotebookDocumentParams) error {
	return s.inner.DidCloseNotebookDocument(ctx, params)
}

func (s *serverWrapper) DidOpenNotebookDocument(ctx context.Context, params *protocol.DidOpenNotebookDocumentParams) error {
	return s.inner.DidOpenNotebookDocument(ctx, params)
}

func (s *serverWrapper) DidSaveNotebookDocument(ctx context.Context, params *protocol.DidSaveNotebookDocumentParams) error {
	return s.inner.DidSaveNotebookDocument(ctx, params)
}

func (s *serverWrapper) Shutdown(ctx context.Context) error {
	return s.inner.Shutdown(ctx)
}

func (s *serverWrapper) CodeAction(ctx context.Context, params *protocol.CodeActionParams) ([]protocol.CodeAction, error) {
	return s.inner.CodeAction(ctx, params)
}

func (s *serverWrapper) CodeLens(ctx context.Context, params *protocol.CodeLensParams) ([]protocol.CodeLens, error) {
	return s.inner.CodeLens(ctx, params)
}

func (s *serverWrapper) ColorPresentation(ctx context.Context, params *protocol.ColorPresentationParams) ([]protocol.ColorPresentation, error) {
	return s.inner.ColorPresentation(ctx, params)
}

func (s *serverWrapper) Completion(ctx context.Context, params *protocol.CompletionParams) (*protocol.CompletionList, error) {
	return s.inner.Completion(ctx, params)
}

func (s *serverWrapper) Declaration(ctx context.Context, params *protocol.DeclarationParams) (*protocol.Or_textDocument_declaration, error) {
	return s.inner.Declaration(ctx, params)
}

func (s *serverWrapper) Definition(ctx context.Context, params *protocol.DefinitionParams) ([]protocol.Location, error) {
	return s.inner.Definition(ctx, params)
}

func (s *serverWrapper) Diagnostic(ctx context.Context, params *string) (*string, error) {
	return s.inner.Diagnostic(ctx, params)
}

func (s *serverWrapper) DidChange(ctx context.Context, params *protocol.DidChangeTextDocumentParams) error {
	return s.inner.DidChange(ctx, params)
}

func (s *serverWrapper) DidClose(ctx context.Context, params *protocol.DidCloseTextDocumentParams) error {
	return s.inner.DidClose(ctx, params)
}

func (s *serverWrapper) DidOpen(ctx context.Context, params *protocol.DidOpenTextDocumentParams) error {
	return s.inner.DidOpen(ctx, params)
}

func (s *serverWrapper) DidSave(ctx context.Context, params *protocol.DidSaveTextDocumentParams) error {
	return s.inner.DidSave(ctx, params)
}

func (s *serverWrapper) DocumentColor(ctx context.Context, params *protocol.DocumentColorParams) ([]protocol.ColorInformation, error) {
	return s.inner.DocumentColor(ctx, params)
}

func (s *serverWrapper) DocumentHighlight(ctx context.Context, params *protocol.DocumentHighlightParams) ([]protocol.DocumentHighlight, error) {
	return s.inner.DocumentHighlight(ctx, params)
}

func (s *serverWrapper) DocumentLink(ctx context.Context, params *protocol.DocumentLinkParams) ([]protocol.DocumentLink, error) {
	return s.inner.DocumentLink(ctx, params)
}

func (s *serverWrapper) DocumentSymbol(ctx context.Context, params *protocol.DocumentSymbolParams) ([]interface{}, error) {
	return s.inner.DocumentSymbol(ctx, params)
}

func (s *serverWrapper) FoldingRange(ctx context.Context, params *protocol.FoldingRangeParams) ([]protocol.FoldingRange, error) {
	return s.inner.FoldingRange(ctx, params)
}

func (s *serverWrapper) Formatting(ctx context.Context, params *protocol.DocumentFormattingParams) ([]protocol.TextEdit, error) {
	return s.inner.Formatting(ctx, params)
}

func (s *serverWrapper) Hover(ctx context.Context, params *protocol.HoverParams) (*protocol.Hover, error) {
	return s.inner.Hover(ctx, params)
}

func (s *serverWrapper) Implementation(ctx context.Context, params *protocol.ImplementationParams) ([]protocol.Location, error) {
	return s.inner.Implementation(ctx, params)
}

func (s *serverWrapper) InlayHint(ctx context.Context, params *protocol.InlayHintParams) ([]protocol.InlayHint, error) {
	return s.inner.InlayHint(ctx, params)
}

func (s *serverWrapper) InlineValue(ctx context.Context, params *protocol.InlineValueParams) ([]protocol.InlineValue, error) {
	return s.inner.InlineValue(ctx, params)
}

func (s *serverWrapper) LinkedEditingRange(ctx context.Context, params *protocol.LinkedEditingRangeParams) (*protocol.LinkedEditingRanges, error) {
	return s.inner.LinkedEditingRange(ctx, params)
}

func (s *serverWrapper) Moniker(ctx context.Context, params *protocol.MonikerParams) ([]protocol.Moniker, error) {
	return s.inner.Moniker(ctx, params)
}

func (s *serverWrapper) OnTypeFormatting(ctx context.Context, params *protocol.DocumentOnTypeFormattingParams) ([]protocol.TextEdit, error) {
	return s.inner.OnTypeFormatting(ctx, params)
}

func (s *serverWrapper) PrepareCallHierarchy(ctx context.Context, params *protocol.CallHierarchyPrepareParams) ([]protocol.CallHierarchyItem, error) {
	return s.inner.PrepareCallHierarchy(ctx, params)
}

func (s *serverWrapper) PrepareRename(ctx context.Context, params *protocol.PrepareRenameParams) (*protocol.PrepareRename2Gn, error) {
	return s.inner.PrepareRename(ctx, params)
}

func (s *serverWrapper) PrepareTypeHierarchy(ctx context.Context, params *protocol.TypeHierarchyPrepareParams) ([]protocol.TypeHierarchyItem, error) {
	return s.inner.PrepareTypeHierarchy(ctx, params)
}

func (s *serverWrapper) RangeFormatting(ctx context.Context, params *protocol.DocumentRangeFormattingParams) ([]protocol.TextEdit, error) {
	return s.inner.RangeFormatting(ctx, params)
}

func (s *serverWrapper) References(ctx context.Context, params *protocol.ReferenceParams) ([]protocol.Location, error) {
	return s.inner.References(ctx, params)
}

func (s *serverWrapper) Rename(ctx context.Context, params *protocol.RenameParams) (*protocol.WorkspaceEdit, error) {
	return s.inner.Rename(ctx, params)
}

func (s *serverWrapper) SelectionRange(ctx context.Context, params *protocol.SelectionRangeParams) ([]protocol.SelectionRange, error) {
	return s.inner.SelectionRange(ctx, params)
}

func (s *serverWrapper) SemanticTokensFull(ctx context.Context, params *protocol.SemanticTokensParams) (*protocol.SemanticTokens, error) {
	return s.inner.SemanticTokensFull(ctx, params)
}

func (s *serverWrapper) SemanticTokensFullDelta(ctx context.Context, params *protocol.SemanticTokensDeltaParams) (interface{}, error) {
	return s.inner.SemanticTokensFullDelta(ctx, params)
}

func (s *serverWrapper) SemanticTokensRange(ctx context.Context, params *protocol.SemanticTokensRangeParams) (*protocol.SemanticTokens, error) {
	return s.inner.SemanticTokensRange(ctx, params)
}

func (s *serverWrapper) SignatureHelp(ctx context.Context, params *protocol.SignatureHelpParams) (*protocol.SignatureHelp, error) {
	return s.inner.SignatureHelp(ctx, params)
}

func (s *serverWrapper) TypeDefinition(ctx context.Context, params *protocol.TypeDefinitionParams) ([]protocol.Location, error) {
	return s.inner.TypeDefinition(ctx, params)
}

func (s *serverWrapper) WillSave(ctx context.Context, params *protocol.WillSaveTextDocumentParams) error {
	return s.inner.WillSave(ctx, params)
}

func (s *serverWrapper) WillSaveWaitUntil(ctx context.Context, params *protocol.WillSaveTextDocumentParams) ([]protocol.TextEdit, error) {
	return s.inner.WillSaveWaitUntil(ctx, params)
}

func (s *serverWrapper) Subtypes(ctx context.Context, params *protocol.TypeHierarchySubtypesParams) ([]protocol.TypeHierarchyItem, error) {
	return s.inner.Subtypes(ctx, params)
}

func (s *serverWrapper) Supertypes(ctx context.Context, params *protocol.TypeHierarchySupertypesParams) ([]protocol.TypeHierarchyItem, error) {
	return s.inner.Supertypes(ctx, params)
}

func (s *serverWrapper) WorkDoneProgressCancel(ctx context.Context, params *protocol.WorkDoneProgressCancelParams) error {
	return s.inner.WorkDoneProgressCancel(ctx, params)
}

func (s *serverWrapper) DiagnosticWorkspace(ctx context.Context, params *protocol.WorkspaceDiagnosticParams) (*protocol.WorkspaceDiagnosticReport, error) {
	return s.inner.DiagnosticWorkspace(ctx, params)
}

func (s *serverWrapper) DiagnosticRefresh(ctx context.Context) error {
	return s.inner.DiagnosticRefresh(ctx)
}

func (s *serverWrapper) DidChangeConfiguration(ctx context.Context, params *protocol.DidChangeConfigurationParams) error {
	return s.inner.DidChangeConfiguration(ctx, params)
}

func (s *serverWrapper) DidChangeWatchedFiles(ctx context.Context, params *protocol.DidChangeWatchedFilesParams) error {
	return s.inner.DidChangeWatchedFiles(ctx, params)
}

func (s *serverWrapper) DidChangeWorkspaceFolders(ctx context.Context, params *protocol.DidChangeWorkspaceFoldersParams) error {
	return s.inner.DidChangeWorkspaceFolders(ctx, params)
}

func (s *serverWrapper) DidCreateFiles(ctx context.Context, params *protocol.CreateFilesParams) error {
	return s.inner.DidCreateFiles(ctx, params)
}

func (s *serverWrapper) DidDeleteFiles(ctx context.Context, params *protocol.DeleteFilesParams) error {
	return s.inner.DidDeleteFiles(ctx, params)
}

func (s *serverWrapper) DidRenameFiles(ctx context.Context, params *protocol.RenameFilesParams) error {
	return s.inner.DidRenameFiles(ctx, params)
}

func (s *serverWrapper) ExecuteCommand(ctx context.Context, params *protocol.ExecuteCommandParams) (interface{}, error) {
	return s.inner.ExecuteCommand(ctx, params)
}

func (s *serverWrapper) InlayHintRefresh(ctx context.Context) error {
	return s.inner.InlayHintRefresh(ctx)
}

func (s *serverWrapper) InlineValueRefresh(ctx context.Context) error {
	return s.inner.InlineValueRefresh(ctx)
}

func (s *serverWrapper) SemanticTokensRefresh(ctx context.Context) error {
	return s.inner.SemanticTokensRefresh(ctx)
}

func (s *serverWrapper) Symbol(ctx context.Context, params *protocol.WorkspaceSymbolParams) ([]protocol.SymbolInformation, error) {
	return s.inner.Symbol(ctx, params)
}

func (s *serverWrapper) WillCreateFiles(ctx context.Context, params *protocol.CreateFilesParams) (*protocol.WorkspaceEdit, error) {
	return s.inner.WillCreateFiles(ctx, params)
}

func (s *serverWrapper) WillDeleteFiles(ctx context.Context, params *protocol.DeleteFilesParams) (*protocol.WorkspaceEdit, error) {
	return s.inner.WillDeleteFiles(ctx, params)
}

func (s *serverWrapper) WillRenameFiles(ctx context.Context, params *protocol.RenameFilesParams) (*protocol.WorkspaceEdit, error) {
	return s.inner.WillRenameFiles(ctx, params)
}

func (s *serverWrapper) ResolveWorkspaceSymbol(ctx context.Context, params *protocol.WorkspaceSymbol) (*protocol.WorkspaceSymbol, error) {
	return s.inner.ResolveWorkspaceSymbol(ctx, params)
}

func (s *serverWrapper) NonstandardRequest(ctx context.Context, method string, params interface{}) (interface{}, error) {
	return s.inner.NonstandardRequest(ctx, method, params)
}
