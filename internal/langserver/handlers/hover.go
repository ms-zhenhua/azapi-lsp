package handlers

import (
	"context"

	"github.com/ms-henglu/azurerm-restapi-lsp/internal/langserver/handlers/hover"

	lsctx "github.com/ms-henglu/azurerm-restapi-lsp/internal/context"
	ilsp "github.com/ms-henglu/azurerm-restapi-lsp/internal/lsp"
	lsp "github.com/ms-henglu/azurerm-restapi-lsp/internal/protocol"
)

func (svc *service) TextDocumentHover(ctx context.Context, params lsp.TextDocumentPositionParams) (*lsp.Hover, error) {
	fs, err := lsctx.DocumentStorage(ctx)
	if err != nil {
		return nil, err
	}

	cc, err := ilsp.ClientCapabilities(ctx)
	if err != nil {
		return nil, err
	}

	doc, err := fs.GetDocument(ilsp.FileHandlerFromDocumentURI(params.TextDocument.URI))
	if err != nil {
		return nil, err
	}

	fPos, err := ilsp.FilePositionFromDocumentPosition(params, doc)
	if err != nil {
		return nil, err
	}

	data, err := doc.Text()
	if err != nil {
		return nil, err
	}

	svc.logger.Printf("Looking for hover data at %q -> %#v", doc.Filename(), fPos.Position())
	hoverData := hover.HoverAtPos(data, doc.Filename(), fPos.Position(), svc.logger)
	svc.logger.Printf("received hover data: %#v", hoverData)

	return ilsp.HoverData(hoverData, cc.TextDocument), nil
}
