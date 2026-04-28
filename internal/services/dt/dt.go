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

var apiNames = []string{"listOrderNoByDono", "pickingDocument", "queryPrint", "queryStockImeiMyStore", "reprintReceiptForm", "updateSimSerialPerso", "authenticate"}

func NewDT(app *app.App) DT {
	app.Service.InitServiceStore(apiNames)
	return &dt{app: app}
}
