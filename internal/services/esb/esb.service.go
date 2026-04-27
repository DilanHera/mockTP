package esb

import (
	"encoding/json"

	"github.com/DilanHera/mockTP/internal/app"
	"github.com/DilanHera/mockTP/internal/store"
)

var (
	apiNames = []string{
		"oauthToken",
		"createFreightOrder",
		"doCreation",
		"legoUpdateOrderStatus",
		"persoSim",
		"serialNumberExpirationDate",
	}

	UserOauthToken                 *OauthTokenResponse
	UserCreateFreightOrder         *CreateFreightOrderResponse
	UserDOCreation                 *DOCreationResponse
	UserLegoupdateOrderStatus      *LegoupdateOrderStatusResponse
	UserPersosim                   *PersosimResponse
	UserSerialNumberExpirationDate *SerialNumberExpirationDateResponse
)

func (e *esb) GetApiInfo(apiName string) store.ApiInfo {
	res, err := e.app.AppInfoStore.Get(apiName)
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
	case "oauthToken":
		var r OauthTokenResponse
		err := json.Unmarshal(resp, &r)
		if err != nil {
			break
		}
		UserOauthToken = &r
	case "createFreightOrder":
		var r CreateFreightOrderResponse
		err := json.Unmarshal(resp, &r)
		if err != nil {
			break
		}
		UserCreateFreightOrder = &r
	case "doCreation":
		var r DOCreationResponse
		err := json.Unmarshal(resp, &r)
		if err != nil {
			break
		}
		UserDOCreation = &r
	case "legoUpdateOrderStatus":
		var r LegoupdateOrderStatusResponse
		err := json.Unmarshal(resp, &r)
		if err != nil {
			break
		}
		UserLegoupdateOrderStatus = &r
	case "persoSim":
		var r PersosimResponse
		err := json.Unmarshal(resp, &r)
		if err != nil {
			break
		}
		UserPersosim = &r
	case "serialNumberExpirationDate":
		var r SerialNumberExpirationDateResponse
		err := json.Unmarshal(resp, &r)
		if err != nil {
			break
		}
		UserSerialNumberExpirationDate = &r
	}
}

func InitESBStore(app *app.App) {
	for _, apiName := range apiNames {
		app.AppInfoStore.Create(apiName, "", "S")
	}
}
