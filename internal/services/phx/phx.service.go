package phx

import (
	"encoding/json"

	"github.com/DilanHera/mockTP/internal/app"
	"github.com/DilanHera/mockTP/internal/store"
)

var (
	apiNames = []string{"requestESIM", "newRegistration"}
	// ApiStates = make(map[string]string)

	UserRequestESIM     *RequestESIMResponse
	UserNewRegistration *NewRegistrationResponse
)

func (p *phx) GetApiInfo(apiName string) store.ApiInfo {
	res, err := p.app.AppInfoStore.Get(apiName)
	if err != nil {
		return store.ApiInfo{}
	}
	if res.Resp != "" {
		CreateResponse([]byte(res.Resp), apiName)
	}
	return *res
}

func CreateResponse(resp []byte, name string) {
	switch name {
	case "requestESIM":
		var r RequestESIMResponse
		err := json.Unmarshal(resp, &r)
		if err != nil {
			break
		}
		UserRequestESIM = &r
	case "newRegistration":
		var r NewRegistrationResponse
		err := json.Unmarshal(resp, &r)
		if err != nil {
			break
		}
		UserNewRegistration = &r
	}
}

func InitPhxStore(app *app.App) {
	for _, apiName := range apiNames {
		app.AppInfoStore.Create(apiName, "", "S")
	}
}
