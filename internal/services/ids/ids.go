package ids

import "github.com/DilanHera/mockTP/internal/app"

var apiNames = []string{"userInfo"}

type Ids interface {
	Authen(req *AuthenRequest) (AuthenResponse, error)
	UserInfo(req *UserInfoRequest) (UserInfoResponse, error)
}

type ids struct {
	app *app.App
	ids Ids
}

func NewIds(app *app.App) Ids {
	app.Service.InitServiceStore(apiNames)
	return &ids{app: app}
}
