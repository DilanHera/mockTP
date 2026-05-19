package im

import (
	"github.com/DilanHera/mockTP/internal/app"
)

var apiNames = []string{"sendSimSerialNo"}

type IM interface {
	SendSimSerialNo(input *SendSimSerialNoRequest) (*SendSimSerialNoResponse, error)
	// SetUserSendSimSerialNo(jsonData json.RawMessage) error
}

type im struct {
	app *app.App
}

func NewIM(app *app.App) IM {
	app.Service.InitServiceStore(apiNames)
	return &im{app: app}
}
