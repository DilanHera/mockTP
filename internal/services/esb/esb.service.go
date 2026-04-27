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

func InitESBStore(app *app.App) {
	for _, apiName := range apiNames {
		app.ApiInfoStore.Create(apiName, "", "S", 200)
	}
}
