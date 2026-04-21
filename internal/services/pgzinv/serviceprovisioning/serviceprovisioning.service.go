package serviceprovisioning

import (
	"encoding/json"
	"fmt"

	"github.com/DilanHera/mockTP/internal/app"
	"github.com/DilanHera/mockTP/internal/store"
)

var (
	resources = []string{"lockNumberByCriteriaPrepaid", "lockNumberByCriteriaPostpaid", "lockNumberByMobilePrepaid",
		"lockNumberByMobilePostpaid", "clearNumberPreparationPrepaid", "clearNumberPreparationPostpaid", "querySimInfo",
		"requestPrepNoPrepaid", "requestPrepNoPostpaid", "confirmPreparationPrepaid", "confirmPreparationPostpaid"}

	// ResourceStates = make(map[string]string)

	UserLockNumberByCriteriaPrepaid    *LockNumberByCriteriaResponse
	UserLockNumberByCriteriaPostpaid   *LockNumberByCriteriaResponse
	UserLockNumberByMobilePrepaid      *LockNumberByMobileResponse
	UserLockNumberByMobilePostpaid     *LockNumberByMobileResponse
	UserClearNumberPreparationPrepaid  *ClearNumberPreparationResponse
	UserClearNumberPreparationPostpaid *ClearNumberPreparationResponse
	UserQuerySimInfo                   *QuerySimInfoResponse
	UserRequestPrepNoPrepaid           *RequestPrepNoResponse
	UserRequestPrepNoPostpaid          *RequestPrepNoResponse
	UserConfirmPreparationPrepaid      *ConfirmPreparationResponse
	UserConfirmPreparationPostpaid     *ConfirmPreparationResponse
)

func (s *serviceProvisioning) GetApiInfo(resourceName string) store.ApiInfo {
	res, err := s.app.AppInfoStore.Get(resourceName)
	if err != nil {
		fmt.Println(err)
		return store.ApiInfo{}
	}
	if res.Resp != "" {
		CreateResponse([]byte(res.Resp), resourceName)
	}
	return *res
}

func CreateResponse(resp []byte, name string) {
	switch name {
	case "lockNumberByCriteriaPrepaid", "lockNumberByCriteriaPostpaid":
		var r LockNumberByCriteriaResponse
		err := json.Unmarshal(resp, &r)
		if err != nil {
			break
		}
		if name == "lockNumberByCriteriaPrepaid" {
			UserLockNumberByCriteriaPrepaid = &r
		} else {
			UserLockNumberByCriteriaPostpaid = &r
		}
	case "lockNumberByMobilePrepaid", "lockNumberByMobilePostpaid":
		var r LockNumberByMobileResponse
		err := json.Unmarshal(resp, &r)
		if err != nil {
			break
		}
		if name == "lockNumberByMobilePrepaid" {
			UserLockNumberByMobilePrepaid = &r
		} else {
			UserLockNumberByMobilePostpaid = &r
		}
	case "confirmPreparationPrepaid", "confirmPreparationPostpaid":
		var r ConfirmPreparationResponse
		err := json.Unmarshal(resp, &r)
		if err != nil {
			break
		}
		if name == "confirmPreparationPrepaid" {
			UserConfirmPreparationPrepaid = &r
		} else {
			UserConfirmPreparationPostpaid = &r
		}
	case "querySimInfo":
		var r QuerySimInfoResponse
		err := json.Unmarshal(resp, &r)
		if err != nil {
			break
		}
		UserQuerySimInfo = &r
	case "clearNumberPreparationPrepaid", "clearNumberPreparationPostpaid":
		var r ClearNumberPreparationResponse
		err := json.Unmarshal(resp, &r)
		if err != nil {
			break
		}
		if name == "clearNumberPreparationPrepaid" {
			UserClearNumberPreparationPrepaid = &r
		} else {
			UserClearNumberPreparationPostpaid = &r
		}
	case "requestPrepNoPrepaid", "requestPrepNoPostpaid":
		var r RequestPrepNoResponse
		err := json.Unmarshal(resp, &r)
		if err != nil {
			break
		}
		if name == "requestPrepNoPrepaid" {
			UserRequestPrepNoPrepaid = &r
		} else {
			UserRequestPrepNoPostpaid = &r
		}
	}
}

func InitServiceProvisioningStore(app *app.App) {
	for _, resourceName := range resources {
		app.AppInfoStore.Create(resourceName, "", "S")
	}
}
