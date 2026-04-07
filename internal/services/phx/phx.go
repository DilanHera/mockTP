package phx

import (
	"encoding/json"

	"github.com/DilanHera/mockTP/internal/app"
)

type Phx interface {
	RequestESIM(input *RequestESIMRequest) (*RequestESIMResponse, error)
	NewRegistration(input *NewRegistrationRequest) (*NewRegistrationResponse, error)
	// Set mock response from user
	SetUserRequestESIM(json.RawMessage) error
	SetUserNewRegistration(json.RawMessage) error
}

type phx struct {
	app *app.App
}

func NewPhx(app *app.App) Phx {
	return &phx{app: app}
}
