package dt

import (
	"github.com/DilanHera/mockTP/internal/app"
)

type DT interface {
	ListOrderNoByDono(input []string) (*ListOrderNoByDonoResponse, error)
	PickingDocument(input *PickingDocumentRequest) (*PickingDocumentResponse, error)
	QueryPrint(input *QueryPrintRequest) (*QueryPrintResponse, error)
	QueryStockImeiMyStore(input *QueryStockImeiMyStoreRequest) (*QueryStockImeiMyStoreResponse, error)
	ReprintReceiptForm(input *ReprintReceiptFormRequest) (*ReprintReceiptFormResponse, error)
	UpdateSimSerialPerso(input *UpdateSimSerialPersoRequest) (*UpdateSimSerialPersoResponse, error)
	Authenticate(input *AuthenticateRequest) (*AuthenticateResponse, error)
}

type dt struct {
	app *app.App
}

func NewDT(app *app.App) DT {
	return &dt{app: app}
}
