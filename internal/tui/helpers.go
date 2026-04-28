package tui

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/DilanHera/mockTP/internal/app"
	"github.com/DilanHera/mockTP/internal/services/dt"
	"github.com/DilanHera/mockTP/internal/services/esb"
	"github.com/DilanHera/mockTP/internal/services/im"
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
	result, err := m.app.ApiInfoStore.Get(name)
	if err != nil {
		return ""
	}
	return string(result.Resp)
}

func (m *model) HttpStatusCodePlaceholder(name string) string {
	result, err := m.app.ApiInfoStore.Get(name)
	if err != nil {
		return ""
	}
	return strconv.Itoa(result.HttpCode)
}

func InitApiStates(app *app.App) {
	result, err := app.ApiInfoStore.GetAll()
	if err != nil {
		panic(err)
	}
	for _, r := range *result {
		ApiStates[r.Name] = r.State
	}
}

func (m *model) SetCustomResponse(resourceName string, jsonData json.RawMessage) error {
	if jsonData == nil || string(jsonData) == "" {
		m.app.ApiInfoStore.UpdateResp(resourceName, "")
		return nil
	}
	switch resourceName {
	case "lockNumberByCriteriaPrepaid", "lockNumberByCriteriaPostpaid":
		var r serviceprovisioning.LockNumberByCriteriaResponse
		err := m.app.Helper.DecodeAndValidate(jsonData, &r)
		if err != nil {
			return fmt.Errorf("failed to unmarshal: %w", err)
		}
		if len(r.ResourceItemList) > 0 && r.ResourceItemList[0].ResourceName != resourceName {
			return fmt.Errorf("wrong resource name, expected: %s", resourceName)
		}
	case "lockNumberByMobilePrepaid", "lockNumberByMobilePostpaid":
		var r serviceprovisioning.LockNumberByMobileResponse
		err := m.app.Helper.DecodeAndValidate(jsonData, &r)
		if err != nil {
			return fmt.Errorf("failed to unmarshal: %w", err)
		}
		if len(r.ResourceItemList) > 0 && r.ResourceItemList[0].ResourceName != resourceName {
			return fmt.Errorf("wrong resource name, expected: %s", resourceName)
		}
	case "confirmPreparationPrepaid", "confirmPreparationPostpaid":
		var r serviceprovisioning.ConfirmPreparationResponse
		err := m.app.Helper.DecodeAndValidate(jsonData, &r)
		if err != nil {
			return fmt.Errorf("failed to unmarshal: %w", err)
		}
		if len(r.ResourceItemList) > 0 && r.ResourceItemList[0].ResourceName != resourceName {
			return fmt.Errorf("wrong resource name, expected: %s", resourceName)
		}
	case "querySimInfo":
		var r serviceprovisioning.QuerySimInfoResponse
		err := m.app.Helper.DecodeAndValidate(jsonData, &r)
		if err != nil {
			return fmt.Errorf("failed to unmarshal: %w", err)
		}
		if len(r.ResourceItemList) > 0 && r.ResourceItemList[0].ResourceName != resourceName {
			return fmt.Errorf("wrong resource name, expected: %s", resourceName)
		}
	case "clearNumberPreparationPrepaid", "clearNumberPreparationPostpaid":
		var r serviceprovisioning.ClearNumberPreparationResponse
		err := m.app.Helper.DecodeAndValidate(jsonData, &r)
		if err != nil {
			return fmt.Errorf("failed to unmarshal: %w", err)
		}
		if len(r.ResourceItemList) > 0 && r.ResourceItemList[0].ResourceName != resourceName {
			return fmt.Errorf("wrong resource name, expected: %s", resourceName)
		}
	case "requestPrepNoPrepaid", "requestPrepNoPostpaid":
		var r serviceprovisioning.RequestPrepNoResponse
		err := m.app.Helper.DecodeAndValidate(jsonData, &r)
		if err != nil {
			return fmt.Errorf("failed to unmarshal: %w", err)
		}
		if len(r.ResourceItemList) > 0 && r.ResourceItemList[0].ResourceName != resourceName {
			return fmt.Errorf("wrong resource name, expected: %s", resourceName)
		}
	case "requestESIM":
		var r phx.RequestESIMResponse
		err := m.app.Helper.DecodeAndValidate(jsonData, &r)
		if err != nil {
			return fmt.Errorf("failed to unmarshal: %w", err)
		}
	case "newRegistration":
		var r phx.NewRegistrationResponse
		err := m.app.Helper.DecodeAndValidate(jsonData, &r)
		if err != nil {
			return fmt.Errorf("failed to unmarshal: %w", err)
		}
	case "encryptLib":
		var r phx.EncryptLibResponse
		err := m.app.Helper.DecodeAndValidate(jsonData, &r)
		if err != nil {
			return fmt.Errorf("failed to unmarshal: %w", err)
		}
	case "checkPerso": // continue from here
		var r phx.CheckPersoResponse
		err := m.app.Helper.DecodeAndValidate(jsonData, &r)
		if err != nil {
			return fmt.Errorf("failed to unmarshal: %w", err)
		}
	case "productProvisioning":
		var r phx.ProductProvisioningResponse
		err := m.app.Helper.DecodeAndValidate(jsonData, &r)
		if err != nil {
			return fmt.Errorf("failed to unmarshal: %w", err)
		}
	case "listOrderNoByDono":
		var r dt.ListOrderNoByDonoResponse
		err := m.app.Helper.DecodeAndValidate(jsonData, &r)
		if err != nil {
			return fmt.Errorf("failed to unmarshal: %w", err)
		}
	case "pickingDocument":
		var r dt.PickingDocumentResponse
		err := m.app.Helper.DecodeAndValidate(jsonData, &r)
		if err != nil {
			return fmt.Errorf("failed to unmarshal: %w", err)
		}
	case "queryPrint":
		var r dt.QueryPrintResponse
		err := m.app.Helper.DecodeAndValidate(jsonData, &r)
		if err != nil {
			return fmt.Errorf("failed to unmarshal: %w", err)
		}
	case "queryStockImeiMyStore":
		var r dt.QueryStockImeiMyStoreResponse
		err := m.app.Helper.DecodeAndValidate(jsonData, &r)
		if err != nil {
			return fmt.Errorf("failed to unmarshal: %w", err)
		}
	case "reprintReceiptForm":
		var r dt.ReprintReceiptFormResponse
		err := m.app.Helper.DecodeAndValidate(jsonData, &r)
		if err != nil {
			return fmt.Errorf("failed to unmarshal: %w", err)
		}
	case "updateSimSerialPerso":
		var r dt.UpdateSimSerialPersoResponse
		err := m.app.Helper.DecodeAndValidate(jsonData, &r)
		if err != nil {
			return fmt.Errorf("failed to unmarshal: %w", err)
		}
	case "authenticate":
		var r dt.AuthenticateResponse
		err := m.app.Helper.DecodeAndValidate(jsonData, &r)
		if err != nil {
			return fmt.Errorf("failed to unmarshal: %w", err)
		}
	case "sendSimSerialNo":
		var r im.SendSimSerialNoResponse
		err := m.app.Helper.DecodeAndValidate(jsonData, &r)
		if err != nil {
			return fmt.Errorf("failed to unmarshal: %w", err)
		}
	case "oauthToken":
		var r esb.OauthTokenResponse
		err := m.app.Helper.DecodeAndValidate(jsonData, &r)
		if err != nil {
			return fmt.Errorf("failed to unmarshal: %w", err)
		}
	case "createFreightOrder":
		var r esb.CreateFreightOrderResponse
		err := m.app.Helper.DecodeAndValidate(jsonData, &r)
		if err != nil {
			return fmt.Errorf("failed to unmarshal: %w", err)
		}
	case "doCreation":
		var r esb.DOCreationResponse
		err := m.app.Helper.DecodeAndValidate(jsonData, &r)
		if err != nil {
			return fmt.Errorf("failed to unmarshal: %w", err)
		}
	case "legoUpdateOrderStatus":
		var r esb.LegoupdateOrderStatusResponse
		err := m.app.Helper.DecodeAndValidate(jsonData, &r)
		if err != nil {
			return fmt.Errorf("failed to unmarshal: %w", err)
		}
	case "persoSim":
		var r esb.PersosimResponse
		err := m.app.Helper.DecodeAndValidate(jsonData, &r)
		if err != nil {
			return fmt.Errorf("failed to unmarshal: %w", err)
		}
	case "serialNumberExpirationDate":
		var r esb.SerialNumberExpirationDateResponse
		err := m.app.Helper.DecodeAndValidate(jsonData, &r)
		if err != nil {
			return fmt.Errorf("failed to unmarshal: %w", err)
		}
	default:
		return fmt.Errorf("unknown resource name: %s", resourceName)
	}
	m.app.ApiInfoStore.UpdateResp(resourceName, string(jsonData))
	return nil
}

func ToggleApiState(resourceName string, app *app.App) {
	ApiStates[resourceName] = app.Helper.ToggleApiState(ApiStates[resourceName])
	app.ApiInfoStore.UpdateState(resourceName, ApiStates[resourceName])
}
