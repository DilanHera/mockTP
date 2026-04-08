package dt

import (
	"encoding/json"
	"fmt"
)

var UserListOrderNoByDono *ListOrderNoByDonoResponse

func (d *dt) SetUserListOrderNoByDono(jsonData json.RawMessage) error {
	response := ListOrderNoByDonoResponse{}
	err := json.Unmarshal(jsonData, &response)
	if err != nil {
		return fmt.Errorf("failed to unmarshal: %w", err)
	}
	UserListOrderNoByDono = &response
	return nil
}