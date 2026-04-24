package dt

import (
	"encoding/json"
	"fmt"

	"github.com/DilanHera/mockTP/internal/app"
	"github.com/DilanHera/mockTP/internal/store"
)

var (
	apiNames = []string{"listOrderNoByDono", "pickingDocument", "queryPrint", "queryStockImeiMyStore", "reprintReceiptForm", "updateSimSerialPerso", "authenticate"}

	UserListOrderNoByDono     *ListOrderNoByDonoResponse
	UserPickingDocument       *PickingDocumentResponse
	UserQueryPrint            *QueryPrintResponse
	UserQueryStockImeiMyStore *QueryStockImeiMyStoreResponse
	UserReprintReceiptForm    *ReprintReceiptFormResponse
	UserUpdateSimSerialPerso  *UpdateSimSerialPersoResponse
	UserAuthenticate          *AuthenticateResponse
)

func (d *dt) GetApiInfo(apiName string) store.ApiInfo {
	res, err := d.app.AppInfoStore.Get(apiName)
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
	case "listOrderNoByDono":
		var r ListOrderNoByDonoResponse
		err := json.Unmarshal(resp, &r)
		if err != nil {
			break
		}
		UserListOrderNoByDono = &r
	case "pickingDocument":
		var r PickingDocumentResponse
		err := json.Unmarshal(resp, &r)
		if err != nil {
			break
		}
		UserPickingDocument = &r
	case "queryPrint":
		var r QueryPrintResponse
		err := json.Unmarshal(resp, &r)
		if err != nil {
			break
		}
		UserQueryPrint = &r
	case "queryStockImeiMyStore":
		var r QueryStockImeiMyStoreResponse
		err := json.Unmarshal(resp, &r)
		if err != nil {
			break
		}
		UserQueryStockImeiMyStore = &r
	case "reprintReceiptForm":
		var r ReprintReceiptFormResponse
		err := json.Unmarshal(resp, &r)
		if err != nil {
			break
		}
		UserReprintReceiptForm = &r
	case "updateSimSerialPerso":
		var r UpdateSimSerialPersoResponse
		err := json.Unmarshal(resp, &r)
		if err != nil {
			break
		}
		UserUpdateSimSerialPerso = &r
	case "authenticate":
		var r AuthenticateResponse
		err := json.Unmarshal(resp, &r)
		if err != nil {
			break
		}
		UserAuthenticate = &r
	}
}

func InitDTStore(app *app.App) {
	for _, apiName := range apiNames {
		app.AppInfoStore.Create(apiName, "", "S")
	}
}

func (d *dt) SetUserListOrderNoByDono(jsonData json.RawMessage) error {
	if jsonData == nil || string(jsonData) == "" {
		UserListOrderNoByDono = nil
		return nil
	}
	response := ListOrderNoByDonoResponse{}
	err := json.Unmarshal(jsonData, &response)
	if err != nil {
		return fmt.Errorf("failed to unmarshal: %w", err)
	}
	UserListOrderNoByDono = &response
	return nil
}

func (d *dt) SetUserPickingDocument(jsonData json.RawMessage) error {
	if jsonData == nil || string(jsonData) == "" {
		UserPickingDocument = nil
		return nil
	}
	response := PickingDocumentResponse{}
	err := json.Unmarshal(jsonData, &response)
	if err != nil {
		return fmt.Errorf("failed to unmarshal: %w", err)
	}
	UserPickingDocument = &response
	return nil
}

func (d *dt) SetUserQueryPrint(jsonData json.RawMessage) error {
	if jsonData == nil || string(jsonData) == "" {
		UserQueryPrint = nil
		return nil
	}
	response := QueryPrintResponse{}
	err := json.Unmarshal(jsonData, &response)
	if err != nil {
		return fmt.Errorf("failed to unmarshal: %w", err)
	}
	UserQueryPrint = &response
	return nil
}

func (d *dt) SetUserQueryStockImeiMyStore(jsonData json.RawMessage) error {
	if jsonData == nil || string(jsonData) == "" {
		UserQueryStockImeiMyStore = nil
		return nil
	}
	response := QueryStockImeiMyStoreResponse{}
	err := json.Unmarshal(jsonData, &response)
	if err != nil {
		return fmt.Errorf("failed to unmarshal: %w", err)
	}
	UserQueryStockImeiMyStore = &response
	return nil
}

func (d *dt) SetUserReprintReceiptForm(jsonData json.RawMessage) error {
	if jsonData == nil || string(jsonData) == "" {
		UserReprintReceiptForm = nil
		return nil
	}
	response := ReprintReceiptFormResponse{}
	err := json.Unmarshal(jsonData, &response)
	if err != nil {
		return fmt.Errorf("failed to unmarshal: %w", err)
	}
	UserReprintReceiptForm = &response
	return nil
}

func (d *dt) SetUserUpdateSimSerialPerso(jsonData json.RawMessage) error {
	if jsonData == nil || string(jsonData) == "" {
		UserUpdateSimSerialPerso = nil
		return nil
	}
	response := UpdateSimSerialPersoResponse{}
	err := json.Unmarshal(jsonData, &response)
	if err != nil {
		return fmt.Errorf("failed to unmarshal: %w", err)
	}
	UserUpdateSimSerialPerso = &response
	return nil
}

func (d *dt) SetUserAuthenticate(jsonData json.RawMessage) error {
	if jsonData == nil || string(jsonData) == "" {
		UserAuthenticate = nil
		return nil
	}
	response := AuthenticateResponse{}
	err := json.Unmarshal(jsonData, &response)
	if err != nil {
		return fmt.Errorf("failed to unmarshal: %w", err)
	}
	UserAuthenticate = &response
	return nil
}
