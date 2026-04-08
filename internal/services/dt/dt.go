package dt

import (
	"encoding/json"
	"github.com/DilanHera/mockTP/internal/app"
)

type DT interface {
	ListOrderNoByDono(input *ListOrderNoByDonoRequest) (*ListOrderNoByDonoResponse, error)
	PickingDocument(input *PickingDocumentRequest) (*PickingDocumentResponse, error)

	// Set mock response from user
	SetUserListOrderNoByDono(jsonData json.RawMessage) error
	SetUserPickingDocument(jsonData json.RawMessage) error
}

type dt struct {
	app *app.App
}

func NewDT(app *app.App) DT {
	return &dt{app: app}
}