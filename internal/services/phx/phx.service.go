package phx

import (
	"github.com/DilanHera/mockTP/internal/app"
)

var apiNames = []string{"requestESIM", "newRegistration", "encryptLib", "checkPerso", "productProvisioning"}

func InitPhxStore(app *app.App) {
	for _, apiName := range apiNames {
		app.ApiInfoStore.Create(apiName, "", "S", 200)
	}
}
