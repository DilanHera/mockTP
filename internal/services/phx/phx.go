package phx

import (
	"github.com/DilanHera/mockTP/internal/app"
)

var apiNames = []string{"requestESIM", "newRegistration", "encryptLib", "checkPerso", "productProvisioning"}

type Phx interface {
	RequestESIM(input *RequestESIMRequest) (*RequestESIMResponse, error)
	NewRegistration(input *NewRegistrationRequest) (*NewRegistrationResponse, error)
	EncryptLib(input *EncryptLibRequest) (*EncryptLibResponse, error)
	CheckPerso(input *CheckPersoRequest) (*CheckPersoResponse, error)
	ProductProvisioning(input *ProductProvisioningRequest) (*ProductProvisioningResponse, error)
}

type phx struct {
	app *app.App
}

func NewPhx(app *app.App) Phx {
	app.Service.InitServiceStore(apiNames)
	return &phx{app: app}
}
