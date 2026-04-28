package mychannel

import (
	"github.com/DilanHera/mockTP/internal/app"
)

var apiNames = []string{"simSerialNo"}

type MyChannel interface {
	SimSerialNo(input *SimSerialNoRequest) (*SimSerialNoResponse, error)
}

type myChannel struct {
	app *app.App
}

func NewMyChannel(app *app.App) MyChannel {
	app.Service.InitServiceStore(apiNames)
	return &myChannel{app: app}
}
