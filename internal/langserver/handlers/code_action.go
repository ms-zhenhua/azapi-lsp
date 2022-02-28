package handlers

import (
	"context"
	"fmt"

	lsctx "github.com/Azure/azapi-lsp/internal/context"
	ilsp "github.com/Azure/azapi-lsp/internal/lsp"
	lsp "github.com/Azure/azapi-lsp/internal/protocol"
)

func (h *logHandler) TextDocumentCodeAction(ctx context.Context, params lsp.CodeActionParams) []lsp.CodeAction {
	ca, err := h.textDocumentCodeAction(ctx, params)
	if err != nil {
		h.logger.Printf("code action failed: %s", err)
	}

	return ca
}

func (h *logHandler) textDocumentCodeAction(ctx context.Context, params lsp.CodeActionParams) ([]lsp.CodeAction, error) {
	var ca []lsp.CodeAction

	// For action definitions, refer to https://code.visualstudio.com/api/references/vscode-api#CodeActionKind
	// We only support format type code actions at the moment, and do not want to format without the client asking for
	// them, so exit early here if nothing is requested.
	if len(params.Context.Only) == 0 {
		h.logger.Printf("No code action requested, exiting")
		return ca, nil
	}

	for _, o := range params.Context.Only {
		h.logger.Printf("Code actions requested: %q", o)
	}

	wantedCodeActions := ilsp.SupportedCodeActions.Only(params.Context.Only)
	if len(wantedCodeActions) == 0 {
		return nil, fmt.Errorf("could not find a supported code action to execute for %s, wanted %v",
			params.TextDocument.URI, params.Context.Only)
	}

	h.logger.Printf("Code actions supported: %v", wantedCodeActions)

	fh := ilsp.FileHandlerFromDocumentURI(params.TextDocument.URI)

	fs, err := lsctx.DocumentStorage(ctx)
	if err != nil {
		return ca, err
	}
	file, err := fs.GetDocument(fh)
	if err != nil {
		return ca, err
	}
	original, err := file.Text()
	if err != nil {
		return ca, err
	}
	_ = original
	for action := range wantedCodeActions {
		switch action {
		case ilsp.SourceFormatAllTerraform:

			ca = append(ca, lsp.CodeAction{
				Title: "Format Document",
				Kind:  action,
				Edit: lsp.WorkspaceEdit{
					Changes: map[string][]lsp.TextEdit{
						//	string(fh.URI()): edits,
					},
				},
			})
		}
	}

	return ca, nil
}
