package im

import (
	"encoding/json"
	"fmt"
)

var UserSendSimSerialNo *SendSimSerialNoResponse

func (i *im) SetUserSendSimSerialNo(jsonData json.RawMessage) error {
	if jsonData == nil || string(jsonData) == "" {
		UserSendSimSerialNo = nil
		return nil
	}
	response := SendSimSerialNoResponse{}
	err := json.Unmarshal(jsonData, &response)
	if err != nil {
		return fmt.Errorf("failed to unmarshal: %w", err)
	}
	UserSendSimSerialNo = &response
	return nil
}
