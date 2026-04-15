package im

import (
	"encoding/json"

	"github.com/DilanHera/mockTP/internal/app"
)

type IM interface {
	SendSimSerialNo(input *SendSimSerialNoRequest) (*SendSimSerialNoResponse, error)
	SetUserSendSimSerialNo(jsonData json.RawMessage) error
}

type im struct {
	app *app.App
}

func NewIM(app *app.App) IM {
	return &im{app: app}
}
