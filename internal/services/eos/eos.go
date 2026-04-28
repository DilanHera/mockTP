package eos

import "github.com/DilanHera/mockTP/internal/app"

var apiNames = []string{
	"updateSimSerialNo",
}

type Eos interface {
	UpdateSimSerialNo(input *UpdateSimSerialNoRequest) (UpdateSimSerialNoResponse, error)
}

type eos struct {
	app *app.App
}

func NewEos(app *app.App) Eos {
	app.Service.InitServiceStore(apiNames)
	return &eos{app: app}
}
