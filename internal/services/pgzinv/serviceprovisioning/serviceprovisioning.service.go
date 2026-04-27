package serviceprovisioning

import (
	"github.com/DilanHera/mockTP/internal/app"
)

var resources = []string{"lockNumberByCriteriaPrepaid", "lockNumberByCriteriaPostpaid", "lockNumberByMobilePrepaid",
	"lockNumberByMobilePostpaid", "clearNumberPreparationPrepaid", "clearNumberPreparationPostpaid", "querySimInfo",
	"requestPrepNoPrepaid", "requestPrepNoPostpaid", "confirmPreparationPrepaid", "confirmPreparationPostpaid"}

func InitServiceProvisioningStore(app *app.App) {
	for _, resourceName := range resources {
		app.ApiInfoStore.Create(resourceName, "", "S", 200)
	}
}
