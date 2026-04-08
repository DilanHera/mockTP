package dt

import (
	"encoding/json"
	"fmt"
)

var UserListOrderNoByDono *ListOrderNoByDonoResponse
var UserPickingDocument *PickingDocumentResponse

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
