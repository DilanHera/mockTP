package esb

import (
	"encoding/json"

	"github.com/DilanHera/mockTP/internal/app"
)

type ESB interface {
	OauthToken(input *json.RawMessage) (json.RawMessage, error)
	CreateFreightOrder(input *json.RawMessage) (json.RawMessage, error)
	DOCreation(input *json.RawMessage) (json.RawMessage, error)
	LegoupdateOrderStatus(input *json.RawMessage) (json.RawMessage, error)
	Persosim(input *json.RawMessage) (json.RawMessage, error)
	SerialNumberExpirationDate(input *json.RawMessage) (json.RawMessage, error)
}

type esb struct {
	app *app.App
}

func NewESB(app *app.App) ESB {
	return &esb{app: app}
}

