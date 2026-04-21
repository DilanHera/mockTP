package tui

import (
	"encoding/json"
	"fmt"

	"github.com/DilanHera/mockTP/internal/services/pgzinv/serviceprovisioning"
	"github.com/DilanHera/mockTP/internal/services/phx"
)

func IndexOf[T comparable](arr []T, target T) int {
	for i, v := range arr {
		if v == target {
			return i
		}
	}
	return -1
}

func (m *model) MarshalJSONForPlaceholder(name string) string {
	result, err := m.app.CustomRespStore.Get(name)
	if err != nil {
		return ""
	}
	// jsonData, err := json.MarshalIndent(result, "", "  ")
	// if err != nil {
	// 	return ""
	// }
	// if len(jsonData) == 0 || string(jsonData) == "null" {
	// 	return ""
	// }
	return result.Resp
}

func (m *model) SetCustomResponse(resourceName string, jsonData json.RawMessage) error {
	if jsonData == nil || string(jsonData) == "" {
		m.app.CustomRespStore.CreateOrUpdate(resourceName, "")
		switch resourceName {
		case "lockNumberByCriteriaPrepaid":
			serviceprovisioning.UserLockNumberByCriteriaPrepaid = nil
		case "lockNumberByCriteriaPostpaid":
			serviceprovisioning.UserLockNumberByCriteriaPostpaid = nil
		case "lockNumberByMobilePrepaid":
			serviceprovisioning.UserLockNumberByMobilePrepaid = nil
		case "lockNumberByMobilePostpaid":
			serviceprovisioning.UserLockNumberByMobilePostpaid = nil
		case "confirmPreparationPrepaid":
			serviceprovisioning.UserConfirmPreparationPrepaid = nil
		case "confirmPreparationPostpaid":
			serviceprovisioning.UserConfirmPreparationPostpaid = nil
		case "querySimInfo":
			serviceprovisioning.UserQuerySimInfo = nil
		case "requestPrepNoPrepaid":
			serviceprovisioning.UserRequestPrepNoPrepaid = nil
		case "requestPrepNoPostpaid":
			serviceprovisioning.UserRequestPrepNoPostpaid = nil
		case "clearNumberPreparationPrepaid":
			serviceprovisioning.UserClearNumberPreparationPrepaid = nil
		case "clearNumberPreparationPostpaid":
			serviceprovisioning.UserClearNumberPreparationPostpaid = nil
		case "requestESIM":
			phx.UserRequestESIM = nil
		case "newRegistration":
			phx.UserNewRegistration = nil
		}
		return nil
	}
	switch resourceName {
	case "lockNumberByCriteriaPrepaid", "lockNumberByCriteriaPostpaid":
		var r serviceprovisioning.LockNumberByCriteriaResponse
		err := json.Unmarshal(jsonData, &r)
		if err != nil {
			return fmt.Errorf("failed to unmarshal: %w", err)
		}
		if len(r.ResourceItemList) > 0 && r.ResourceItemList[0].ResourceName != resourceName {
			return fmt.Errorf("wrong resource name, expected: %s", resourceName)
		}
		if resourceName == "lockNumberByCriteriaPrepaid" {
			serviceprovisioning.UserLockNumberByCriteriaPrepaid = &r
		} else {
			serviceprovisioning.UserLockNumberByCriteriaPostpaid = &r
		}
	case "lockNumberByMobilePrepaid", "lockNumberByMobilePostpaid":
		var r serviceprovisioning.LockNumberByMobileResponse
		err := json.Unmarshal(jsonData, &r)
		if err != nil {
			return fmt.Errorf("failed to unmarshal: %w", err)
		}
		if len(r.ResourceItemList) > 0 && r.ResourceItemList[0].ResourceName != resourceName {
			return fmt.Errorf("wrong resource name, expected: %s", resourceName)
		}
		if resourceName == "lockNumberByMobilePrepaid" {
			serviceprovisioning.UserLockNumberByMobilePrepaid = &r
		} else {
			serviceprovisioning.UserLockNumberByMobilePostpaid = &r
		}
	case "confirmPreparationPrepaid", "confirmPreparationPostpaid":
		var r serviceprovisioning.ConfirmPreparationResponse
		err := json.Unmarshal(jsonData, &r)
		if err != nil {
			return fmt.Errorf("failed to unmarshal: %w", err)
		}
		if len(r.ResourceItemList) > 0 && r.ResourceItemList[0].ResourceName != resourceName {
			return fmt.Errorf("wrong resource name, expected: %s", resourceName)
		}
		if resourceName == "confirmPreparationPrepaid" {
			serviceprovisioning.UserConfirmPreparationPrepaid = &r
		} else {
			serviceprovisioning.UserConfirmPreparationPostpaid = &r
		}
	case "querySimInfo":
		var r serviceprovisioning.QuerySimInfoResponse
		err := json.Unmarshal(jsonData, &r)
		if err != nil {
			return fmt.Errorf("failed to unmarshal: %w", err)
		}
		if len(r.ResourceItemList) > 0 && r.ResourceItemList[0].ResourceName != resourceName {
			return fmt.Errorf("wrong resource name, expected: %s", resourceName)
		}
		serviceprovisioning.UserQuerySimInfo = &r
	case "clearNumberPreparationPrepaid", "clearNumberPreparationPostpaid":
		var r serviceprovisioning.ClearNumberPreparationResponse
		err := json.Unmarshal(jsonData, &r)
		if err != nil {
			return fmt.Errorf("failed to unmarshal: %w", err)
		}
		if len(r.ResourceItemList) > 0 && r.ResourceItemList[0].ResourceName != resourceName {
			return fmt.Errorf("wrong resource name, expected: %s", resourceName)
		}
		if resourceName == "clearNumberPreparationPrepaid" {
			serviceprovisioning.UserClearNumberPreparationPrepaid = &r
		} else {
			serviceprovisioning.UserClearNumberPreparationPostpaid = &r
		}
	case "requestPrepNoPrepaid", "requestPrepNoPostpaid":
		var r serviceprovisioning.RequestPrepNoResponse
		err := json.Unmarshal(jsonData, &r)
		if err != nil {
			return fmt.Errorf("failed to unmarshal: %w", err)
		}
		if len(r.ResourceItemList) > 0 && r.ResourceItemList[0].ResourceName != resourceName {
			return fmt.Errorf("wrong resource name, expected: %s", resourceName)
		}
		if resourceName == "requestPrepNoPrepaid" {
			serviceprovisioning.UserRequestPrepNoPrepaid = &r
		} else {
			serviceprovisioning.UserRequestPrepNoPostpaid = &r
		}
	case "requestEsimPrepaid", "requestEsimPostpaid":
		var r phx.RequestESIMResponse
		err := json.Unmarshal(jsonData, &r)
		if err != nil {
			return fmt.Errorf("failed to unmarshal: %w", err)
		}
		phx.UserRequestESIM = &r
	case "newRegistration":
		var r phx.NewRegistrationResponse
		err := json.Unmarshal(jsonData, &r)
		if err != nil {
			return fmt.Errorf("failed to unmarshal: %w", err)
		}
		phx.UserNewRegistration = &r
	default:
		return fmt.Errorf("unknown resource name: %s", resourceName)
	}
	m.app.CustomRespStore.CreateOrUpdate(resourceName, string(jsonData))
	return nil
}
