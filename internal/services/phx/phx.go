package phx

import (
	"github.com/DilanHera/mockTP/internal/app"
)

type Phx interface {
	RequestESIM(input *RequestESIMRequest) (*RequestESIMResponse, error)
	NewRegistration(input *NewRegistrationRequest) (*NewRegistrationResponse, error)
}

type phx struct {
	app *app.App
}

func NewPhx(app *app.App) Phx {
	return &phx{app: app}
}
