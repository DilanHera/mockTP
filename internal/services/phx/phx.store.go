package phx

import (
	"encoding/json"
	"fmt"
)

var (
	APIErrorStates = map[string]bool{
		"requestESIM":     false,
		"newRegistration": false,
	}

	UserRequestESIM     *RequestESIMResponse
	UserNewRegistration *NewRegistrationResponse
)

func IsAPIErrorState(apiName string) bool {
	return APIErrorStates[apiName]
}

func HasCustomAPIResponse(apiName string) bool {
	switch apiName {
	case "requestESIM":
		return UserRequestESIM != nil
	case "newRegistration":
		return UserNewRegistration != nil
	default:
		return false
	}
}

func ToggleAPIErrorState(apiName string) {
	if _, ok := APIErrorStates[apiName]; !ok {
		return
	}
	APIErrorStates[apiName] = !APIErrorStates[apiName]
}

func (p *phx) SetUserRequestESIM(jsonData json.RawMessage) error {
	if jsonData == nil || string(jsonData) == "" {
		UserRequestESIM = nil
		return nil
	}
	response := RequestESIMResponse{}
	err := json.Unmarshal(jsonData, &response)
	if err != nil {
		return fmt.Errorf("failed to unmarshal: %w", err)
	}
	UserRequestESIM = &response
	return nil
}

func (p *phx) SetUserNewRegistration(jsonData json.RawMessage) error {
	if jsonData == nil || string(jsonData) == "" {
		UserNewRegistration = nil
		return nil
	}
	response := NewRegistrationResponse{}
	err := json.Unmarshal(jsonData, &response)
	if err != nil {
		return fmt.Errorf("failed to unmarshal: %w", err)
	}
	UserNewRegistration = &response
	return nil
}
