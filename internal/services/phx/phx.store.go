package phx

import (
	"encoding/json"

	"github.com/DilanHera/mockTP/internal/app"
)

var (
	apiNames  = []string{"requestESIM", "newRegistration"}
	ApiStates = make(map[string]string)

	UserRequestESIM     *RequestESIMResponse
	UserNewRegistration *NewRegistrationResponse
)

func GetApiState(apiName string) string {
	return ApiStates[apiName]
}

func ToggleApiState(apiName string, app *app.App) {
	if _, ok := ApiStates[apiName]; !ok {
		return
	}
	ApiStates[apiName] = app.Helper.ToggleApiState(GetApiState(apiName))
}

func InitApis(app *app.App) {
	for _, apiName := range apiNames {
		ApiStates[apiName] = "S"
	}
	results, err := app.CustomRespStore.GetMany(apiNames)
	if err != nil {
		return
	}
	for _, res := range *results {
		respBytes := []byte(res.Resp)
		if len(respBytes) == 0 {
			continue
		}
		switch res.Name {
		case "requestESIM":
			var r RequestESIMResponse
			err := json.Unmarshal(respBytes, &r)
			if err != nil {
				break
			}
			UserRequestESIM = &r
		case "newRegistration":
			var r NewRegistrationResponse
			err := json.Unmarshal(respBytes, &r)
			if err != nil {
				break
			}
			UserNewRegistration = &r
		}
	}
}
