package tui

import (
	"encoding/json"
	"fmt"

	"github.com/DilanHera/mockTP/internal/app"
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
	result, err := m.app.AppInfoStore.Get(name)
	if err != nil {
		return ""
	}
	return result.Resp
}

func InitApiStates(app *app.App) {
	result, err := app.AppInfoStore.GetAll()
	if err != nil {
		panic(err)
	}
	for _, r := range *result {
		ApiStates[r.Name] = r.State
	}
}

func (m *model) SetCustomResponse(resourceName string, jsonData json.RawMessage) error {
	if jsonData == nil || string(jsonData) == "" {
		m.app.AppInfoStore.UpdateResp(resourceName, "")
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
	m.app.AppInfoStore.UpdateResp(resourceName, string(jsonData))
	return nil
}

func ToggleApiState(resourceName string, app *app.App) {
	ApiStates[resourceName] = app.Helper.ToggleApiState(ApiStates[resourceName])
	app.AppInfoStore.UpdateState(resourceName, ApiStates[resourceName])
}
