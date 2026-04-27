package dt

import (
	"github.com/DilanHera/mockTP/internal/app"
)

var apiNames = []string{"listOrderNoByDono", "pickingDocument", "queryPrint", "queryStockImeiMyStore", "reprintReceiptForm", "updateSimSerialPerso", "authenticate"}

func InitDTStore(app *app.App) {
	for _, apiName := range apiNames {
		app.ApiInfoStore.Create(apiName, "", "S", 200)
	}
}
