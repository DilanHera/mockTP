package smis

import "github.com/DilanHera/mockTP/internal/app"

var apiNames = []string{"updateSerial"}

type Smis interface {
	UpdateSerial(input *UpdateSerialRequest) (UpdateSerialResponse, error)
}

type smis struct {
	app *app.App
}

func NewSmis(app *app.App) Smis {
	app.Service.InitServiceStore(apiNames)
	return &smis{app: app}
}
