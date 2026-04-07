package phx

import (
	"encoding/json"
	"fmt"
)

var UserRequestESIM *RequestESIMResponse
var UserNewRegistration *NewRegistrationResponse

func (p *phx) SetUserRequestESIM(jsonData json.RawMessage) error {
	response := RequestESIMResponse{}
	err := json.Unmarshal(jsonData, &response)
	if err != nil {
		return fmt.Errorf("failed to unmarshal: %w", err)
	}
	UserRequestESIM = &response
	return nil
}

func (p *phx) SetUserNewRegistration(jsonData json.RawMessage) error {
	response := NewRegistrationResponse{}
	err := json.Unmarshal(jsonData, &response)
	if err != nil {
		return fmt.Errorf("failed to unmarshal: %w", err)
	}
	UserNewRegistration = &response
	return nil
}