package esb

import (
	"github.com/DilanHera/mockTP/internal/app"
)

type ESB interface {
	OauthToken(input *OauthTokenRequest) (OauthTokenResponse, error)
	CreateFreightOrder(input *CreateFreightOrderRequest) (*CreateFreightOrderResponse, error)
	DOCreation(input *DOCreationRequest) (*DOCreationResponse, error)
	LegoupdateOrderStatus(input *LegoupdateOrderStatusRequest) (*LegoupdateOrderStatusResponse, error)
	Persosim(input *PersosimRequest) (*PersosimResponse, error)
	SerialNumberExpirationDate(input *SerialNumberExpirationDateRequest) (*SerialNumberExpirationDateResponse, error)
}

type esb struct {
	app *app.App
}

func NewESB(app *app.App) ESB {
	return &esb{app: app}
}
