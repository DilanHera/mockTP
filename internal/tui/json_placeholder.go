package tui

import (
	"encoding/json"

	"github.com/DilanHera/mockTP/internal/services/esb"
)

// func(m *model) ServiceProvisioningMockPlaceholder(resourceName string) string {
// 	switch resourceName {
// 	case "lockNumberByCriteriaPrepaid":
// 		return m.MarshalJSONForPlaceholder(serviceprovisioning.UserLockNumberByCriteriaPrepaid)
// 	case "lockNumberByCriteriaPostpaid":
// 		return m.MarshalJSONForPlaceholder(serviceprovisioning.UserLockNumberByCriteriaPostpaid)
// 	case "lockNumberByMobilePrepaid":
// 		return MarshalJSONForPlaceholder(serviceprovisioning.UserLockNumberByMobilePrepaid)
// 	case "lockNumberByMobilePostpaid":
// 		return MarshalJSONForPlaceholder(serviceprovisioning.UserLockNumberByMobilePostpaid)
// 	case "clearNumberPreparationPrepaid":
// 		return MarshalJSONForPlaceholder(serviceprovisioning.UserClearNumberPreparationPrepaid)
// 	case "clearNumberPreparationPostpaid":
// 		return MarshalJSONForPlaceholder(serviceprovisioning.UserClearNumberPreparationPostpaid)
// 	case "querySimInfo":
// 		return MarshalJSONForPlaceholder(serviceprovisioning.UserQuerySimInfo)
// 	case "requestPrepNoPrepaid":
// 		return MarshalJSONForPlaceholder(serviceprovisioning.UserRequestPrepNoPrepaid)
// 	case "requestPrepNoPostpaid":
// 		return MarshalJSONForPlaceholder(serviceprovisioning.UserRequestPrepNoPostpaid)
// 	case "confirmPreparationPrepaid":
// 		return MarshalJSONForPlaceholder(serviceprovisioning.UserConfirmPreparationPrepaid)
// 	case "confirmPreparationPostpaid":
// 		return MarshalJSONForPlaceholder(serviceprovisioning.UserConfirmPreparationPostpaid)
// 	default:
// 		return "{}"
// 	}
// }

// func PHXMockPlaceholder(apiName string) string {
// 	switch apiName {
// 	case "requestESIM":
// 		return MarshalJSONForPlaceholder(phx.UserRequestESIM)
// 	case "newRegistration":
// 		return MarshalJSONForPlaceholder(phx.UserNewRegistration)
// 	default:
// 		return "{}"
// 	}
// }

func EsbMockPlaceholder(apiName string) string {
	switch apiName {
	case "oauthToken":
		res, _ := json.Marshal(esb.OauthTokenResponse{
			AccessToken: "mock-token",
			TokenType:   "bearer",
			ExpiresIn:   3600,
			Error:       "remove for success case",
		})
		return string(res)
	default:
		return "{}"
	}
}
