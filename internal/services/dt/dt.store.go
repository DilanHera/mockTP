package dt

import (
	"encoding/json"
	"fmt"
)

var UserListOrderNoByDono *ListOrderNoByDonoResponse
var UserPickingDocument *PickingDocumentResponse
var UserQueryPrint *QueryPrintResponse
var UserQueryStockImeiMyStore *QueryStockImeiMyStoreResponse
var UserReprintReceiptForm *ReprintReceiptFormResponse
var UserUpdateSimSerialPerso *UpdateSimSerialPersoResponse
var UserAuthenticate *AuthenticateResponse

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
