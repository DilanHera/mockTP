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

	UserOauthToken                 *json.RawMessage
	UserCreateFreightOrder         *json.RawMessage
	UserDOCreation                 *json.RawMessage
	UserLegoupdateOrderStatus      *json.RawMessage
	UserPersosim                   *json.RawMessage
	UserSerialNumberExpirationDate *json.RawMessage
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
	copyBytes := func(b []byte) json.RawMessage {
		dst := make([]byte, len(b))
		copy(dst, b)
		return json.RawMessage(dst)
	}

	switch name {
	case "oauthToken":
		r := copyBytes(resp)
		UserOauthToken = &r
	case "createFreightOrder":
		r := copyBytes(resp)
		UserCreateFreightOrder = &r
	case "doCreation":
		r := copyBytes(resp)
		UserDOCreation = &r
	case "legoUpdateOrderStatus":
		r := copyBytes(resp)
		UserLegoupdateOrderStatus = &r
	case "persoSim":
		r := copyBytes(resp)
		UserPersosim = &r
	case "serialNumberExpirationDate":
		r := copyBytes(resp)
		UserSerialNumberExpirationDate = &r
	}
}

func InitESBStore(app *app.App) {
	for _, apiName := range apiNames {
		app.AppInfoStore.Create(apiName, "", "S")
	}
}

