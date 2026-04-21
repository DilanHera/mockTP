package serviceprovisioning

import (
	"encoding/json"

	"github.com/DilanHera/mockTP/internal/app"
)

var (
	resources = []string{"lockNumberByCriteriaPrepaid", "lockNumberByCriteriaPostpaid", "lockNumberByMobilePrepaid",
		"lockNumberByMobilePostpaid", "clearNumberPreparationPrepaid", "clearNumberPreparationPostpaid", "querySimInfo",
		"requestPrepNoPrepaid", "requestPrepNoPostpaid", "confirmPreparationPrepaid", "confirmPreparationPostpaid"}

	ResourceStates = make(map[string]string)

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

func GetResourceState(resourceName string) string {
	return ResourceStates[resourceName]
}

func ToggleResourceState(resourceName string, app *app.App) {
	if _, ok := ResourceStates[resourceName]; !ok {
		return
	}
	ResourceStates[resourceName] = app.Helper.ToggleApiState(GetResourceState(resourceName))
}

func InitResources(app *app.App) {
	for _, resourceName := range resources {
		ResourceStates[resourceName] = "S"
	}
	results, err := app.CustomRespStore.GetMany(resources)
	if err != nil {
		return
	}
	for _, res := range *results {
		respBytes := []byte(res.Resp)
		if len(respBytes) == 0 {
			continue
		}
		switch res.Name {
		case "lockNumberByCriteriaPrepaid", "lockNumberByCriteriaPostpaid":
			var r LockNumberByCriteriaResponse
			err := json.Unmarshal(respBytes, &r)
			if err != nil {
				break
			}
			if res.Name == "lockNumberByCriteriaPrepaid" {
				UserLockNumberByCriteriaPrepaid = &r
			} else {
				UserLockNumberByCriteriaPostpaid = &r
			}
		case "lockNumberByMobilePrepaid", "lockNumberByMobilePostpaid":
			var r LockNumberByMobileResponse
			err := json.Unmarshal(respBytes, &r)
			if err != nil {
				break
			}
			if res.Name == "lockNumberByMobilePrepaid" {
				UserLockNumberByMobilePrepaid = &r
			} else {
				UserLockNumberByMobilePostpaid = &r
			}
		case "confirmPreparationPrepaid", "confirmPreparationPostpaid":
			var r ConfirmPreparationResponse
			err := json.Unmarshal(respBytes, &r)
			if err != nil {
				break
			}
			if res.Name == "confirmPreparationPrepaid" {
				UserConfirmPreparationPrepaid = &r
			} else {
				UserConfirmPreparationPostpaid = &r
			}
		case "querySimInfo":
			var r QuerySimInfoResponse
			err := json.Unmarshal(respBytes, &r)
			if err != nil {
				break
			}
			UserQuerySimInfo = &r
		case "clearNumberPreparationPrepaid", "clearNumberPreparationPostpaid":
			var r ClearNumberPreparationResponse
			err := json.Unmarshal(respBytes, &r)
			if err != nil {
				break
			}
			if res.Name == "clearNumberPreparationPrepaid" {
				UserClearNumberPreparationPrepaid = &r
			} else {
				UserClearNumberPreparationPostpaid = &r
			}
		case "requestPrepNoPrepaid", "requestPrepNoPostpaid":
			var r RequestPrepNoResponse
			err := json.Unmarshal(respBytes, &r)
			if err != nil {
				break
			}
			if res.Name == "requestPrepNoPrepaid" {
				UserRequestPrepNoPrepaid = &r
			} else {
				UserRequestPrepNoPostpaid = &r
			}
		}
	}
}
