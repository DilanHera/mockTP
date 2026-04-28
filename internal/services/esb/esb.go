package esb

import (
	"github.com/DilanHera/mockTP/internal/app"
)

var apiNames = []string{
	"oauthToken",
	"createFreightOrder",
	"doCreation",
	"legoUpdateOrderStatus",
	"persoSim",
	"serialNumberExpirationDate",
}

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
	app.Service.InitServiceStore(apiNames)
	return &esb{app: app}
}
