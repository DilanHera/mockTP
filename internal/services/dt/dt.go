package dt

import (
	"encoding/json"

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
	// Set mock response from user
	SetUserListOrderNoByDono(jsonData json.RawMessage) error
	SetUserPickingDocument(jsonData json.RawMessage) error
	SetUserQueryPrint(jsonData json.RawMessage) error
	SetUserQueryStockImeiMyStore(jsonData json.RawMessage) error
	SetUserReprintReceiptForm(jsonData json.RawMessage) error
	SetUserUpdateSimSerialPerso(jsonData json.RawMessage) error
	SetUserAuthenticate(jsonData json.RawMessage) error
}

type dt struct {
	app *app.App
}

func NewDT(app *app.App) DT {
	return &dt{app: app}
}
