package serviceprovisioning

import (
	"encoding/json"
	"fmt"
)

var (
	ResourceErrorStates = map[string]bool{
		"lockNumberByCriteriaPrepaid":    false,
		"lockNumberByCriteriaPostpaid":   false,
		"lockNumberByMobilePrepaid":      false,
		"lockNumberByMobilePostpaid":     false,
		"clearNumberPreparationPrepaid":  false,
		"clearNumberPreparationPostpaid": false,
		"querySimInfo":                   false,
		"requestPrepNoPrepaid":           false,
		"requestPrepNoPostpaid":          false,
		"confirmPreparationPrepaid":      false,
		"confirmPreparationPostpaid":     false,
	}

	UserLockNumberByCriteriaPrepaid    *LockNumberByCriteriaResponse
	UserLockNumberByCriteriaPostpaid   *LockNumberByCriteriaResponse
	UserLockNumberByMobilePrepaid      *LockNumberByMobileResponse
	UserLockNumberByMobilePostpaid     *LockNumberByMobileResponse
	UserClearNumberPreparationPrepaid  *ClearNumberPreparationResponse
	UserClearNumberPreparationPostpaid *ClearNumberPreparationResponse
	UserQuerySimInfo                   *QuerySimInfoResponse
	UserRequestPrepNoPrepaid           *RequestPrepNoResponse
	UserRequestPrepNoPostpaid          *RequestPrepNoResponse
	UserConfirmPreparationPrepaid      *ConfirmPreparationResponse
	UserConfirmPreparationPostpaid     *ConfirmPreparationResponse
)

func IsResourceErrorState(resourceName string) bool {
	return ResourceErrorStates[resourceName]
}

func HasCustomResourceResponse(resourceName string) bool {
	switch resourceName {
	case "lockNumberByCriteriaPrepaid":
		return UserLockNumberByCriteriaPrepaid != nil
	case "lockNumberByCriteriaPostpaid":
		return UserLockNumberByCriteriaPostpaid != nil
	case "lockNumberByMobilePrepaid":
		return UserLockNumberByMobilePrepaid != nil
	case "lockNumberByMobilePostpaid":
		return UserLockNumberByMobilePostpaid != nil
	case "clearNumberPreparationPrepaid":
		return UserClearNumberPreparationPrepaid != nil
	case "clearNumberPreparationPostpaid":
		return UserClearNumberPreparationPostpaid != nil
	case "querySimInfo":
		return UserQuerySimInfo != nil
	case "requestPrepNoPrepaid":
		return UserRequestPrepNoPrepaid != nil
	case "requestPrepNoPostpaid":
		return UserRequestPrepNoPostpaid != nil
	case "confirmPreparationPrepaid":
		return UserConfirmPreparationPrepaid != nil
	case "confirmPreparationPostpaid":
		return UserConfirmPreparationPostpaid != nil
	default:
		return false
	}
}

func ToggleResourceErrorState(resourceName string) {
	if _, ok := ResourceErrorStates[resourceName]; !ok {
		return
	}
	ResourceErrorStates[resourceName] = !ResourceErrorStates[resourceName]
}

func (s *serviceProvisioning) SetUserLockNumberByCriteriaPrepaid(jsonData json.RawMessage) error {
	if jsonData == nil || string(jsonData) == "" {
		UserLockNumberByCriteriaPrepaid = nil
		return nil
	}
	response := LockNumberByCriteriaResponse{}
	err := json.Unmarshal(jsonData, &response)
	if err != nil {
		return fmt.Errorf("failed to unmarshal: %w", err)
	}
	if len(response.ResourceItemList) > 0 && response.ResourceItemList[0].ResourceName != "lockNumberByCriteriaPrepaid" {
		return fmt.Errorf("wrong resource name, expected: lockNumberByCriteriaPrepaid")
	}
	UserLockNumberByCriteriaPrepaid = &response
	return nil
}

func (s *serviceProvisioning) SetUserLockNumberByCriteriaPostpaid(jsonData json.RawMessage) error {
	if jsonData == nil || string(jsonData) == "" {
		UserLockNumberByCriteriaPostpaid = nil
		return nil
	}
	response := LockNumberByCriteriaResponse{}
	err := json.Unmarshal(jsonData, &response)
	if err != nil {
		return fmt.Errorf("failed to unmarshal: %w", err)
	}
	if len(response.ResourceItemList) > 0 && response.ResourceItemList[0].ResourceName != "lockNumberByCriteriaPostpaid" {
		return fmt.Errorf("wrong resource name, expected: lockNumberByCriteriaPostpaid")
	}
	UserLockNumberByCriteriaPostpaid = &response
	return nil
}

func (s *serviceProvisioning) SetUserLockNumberByMobilePrepaid(jsonData json.RawMessage) error {
	if jsonData == nil || string(jsonData) == "" {
		UserLockNumberByMobilePrepaid = nil
		return nil
	}
	response := LockNumberByMobileResponse{}
	err := json.Unmarshal(jsonData, &response)
	if err != nil {
		return fmt.Errorf("failed to unmarshal: %w", err)
	}
	if len(response.ResourceItemList) > 0 && response.ResourceItemList[0].ResourceName != "lockNumberByMobilePrepaid" {
		return fmt.Errorf("wrong resource name, expected: lockNumberByMobilePrepaid")
	}
	UserLockNumberByMobilePrepaid = &response
	return nil
}

func (s *serviceProvisioning) SetUserLockNumberByMobilePostpaid(jsonData json.RawMessage) error {
	if jsonData == nil || string(jsonData) == "" {
		UserLockNumberByMobilePostpaid = nil
		return nil
	}
	response := LockNumberByMobileResponse{}
	err := json.Unmarshal(jsonData, &response)
	if err != nil {
		return fmt.Errorf("failed to unmarshal: %w", err)
	}
	if len(response.ResourceItemList) > 0 && response.ResourceItemList[0].ResourceName != "lockNumberByMobilePostpaid" {
		return fmt.Errorf("wrong resource name, expected: lockNumberByMobilePostpaid")
	}
	UserLockNumberByMobilePostpaid = &response
	return nil
}

func (s *serviceProvisioning) SetUserClearNumberPreparationPrepaid(jsonData json.RawMessage) error {
	if jsonData == nil || string(jsonData) == "" {
		UserClearNumberPreparationPrepaid = nil
		return nil
	}
	response := ClearNumberPreparationResponse{}
	err := json.Unmarshal(jsonData, &response)
	if err != nil {
		return fmt.Errorf("failed to unmarshal: %w", err)
	}
	if len(response.ResourceItemList) > 0 && response.ResourceItemList[0].ResourceName != "clearNumberPreparationPrepaid" {
		return fmt.Errorf("wrong resource name, expected: clearNumberPreparationPrepaid")
	}
	UserClearNumberPreparationPrepaid = &response
	return nil
}

func (s *serviceProvisioning) SetUserClearNumberPreparationPostpaid(jsonData json.RawMessage) error {
	if jsonData == nil || string(jsonData) == "" {
		UserClearNumberPreparationPostpaid = nil
		return nil
	}
	response := ClearNumberPreparationResponse{}
	err := json.Unmarshal(jsonData, &response)
	if err != nil {
		return fmt.Errorf("failed to unmarshal: %w", err)
	}
	if len(response.ResourceItemList) > 0 && response.ResourceItemList[0].ResourceName != "clearNumberPreparationPostpaid" {
		return fmt.Errorf("wrong resource name, expected: clearNumberPreparationPostpaid")
	}
	UserClearNumberPreparationPostpaid = &response
	return nil
}

func (s *serviceProvisioning) SetUserQuerySimInfo(jsonData json.RawMessage) error {
	if jsonData == nil || string(jsonData) == "" {
		UserQuerySimInfo = nil
		return nil
	}
	response := QuerySimInfoResponse{}
	err := json.Unmarshal(jsonData, &response)
	if err != nil {
		return fmt.Errorf("failed to unmarshal: %w", err)
	}
	if len(response.ResourceItemList) > 0 && response.ResourceItemList[0].ResourceName != "querySimInfo" {
		return fmt.Errorf("wrong resource name, expected: querySimInfo")
	}
	UserQuerySimInfo = &response
	return nil
}

func (s *serviceProvisioning) SetUserRequestPrepNoPrepaid(jsonData json.RawMessage) error {
	if jsonData == nil || string(jsonData) == "" {
		UserRequestPrepNoPrepaid = nil
		return nil
	}
	response := RequestPrepNoResponse{}
	err := json.Unmarshal(jsonData, &response)
	if err != nil {
		return fmt.Errorf("failed to unmarshal: %w", err)
	}
	if len(response.ResourceItemList) > 0 && response.ResourceItemList[0].ResourceName != "requestPrepNoPrepaid" {
		return fmt.Errorf("wrong resource name, expected: requestPrepNoPrepaid")
	}
	UserRequestPrepNoPrepaid = &response
	return nil
}

func (s *serviceProvisioning) SetUserRequestPrepNoPostpaid(jsonData json.RawMessage) error {
	if jsonData == nil || string(jsonData) == "" {
		UserRequestPrepNoPostpaid = nil
		return nil
	}
	response := RequestPrepNoResponse{}
	err := json.Unmarshal(jsonData, &response)
	if err != nil {
		return fmt.Errorf("failed to unmarshal: %w", err)
	}
	if len(response.ResourceItemList) > 0 && response.ResourceItemList[0].ResourceName != "requestPrepNoPostpaid" {
		return fmt.Errorf("wrong resource name, expected: requestPrepNoPostpaid")
	}
	UserRequestPrepNoPostpaid = &response
	return nil
}

func (s *serviceProvisioning) SetUserConfirmPreparationPrepaid(jsonData json.RawMessage) error {
	if jsonData == nil || string(jsonData) == "" {
		UserConfirmPreparationPrepaid = nil
		return nil
	}
	response := ConfirmPreparationResponse{}
	err := json.Unmarshal(jsonData, &response)
	if err != nil {
		return fmt.Errorf("failed to unmarshal: %w", err)
	}
	if len(response.ResourceItemList) > 0 && response.ResourceItemList[0].ResourceName != "confirmPreparationPrepaid" {
		return fmt.Errorf("wrong resource name, expected: confirmPreparationPrepaid")
	}
	UserConfirmPreparationPrepaid = &response
	return nil
}

func (s *serviceProvisioning) SetUserConfirmPreparationPostpaid(jsonData json.RawMessage) error {
	if jsonData == nil || string(jsonData) == "" {
		UserConfirmPreparationPostpaid = nil
		return nil
	}
	response := ConfirmPreparationResponse{}
	err := json.Unmarshal(jsonData, &response)
	if err != nil {
		return fmt.Errorf("failed to unmarshal: %w", err)
	}
	if len(response.ResourceItemList) > 0 && response.ResourceItemList[0].ResourceName != "confirmPreparationPostpaid" {
		return fmt.Errorf("wrong resource name, expected: confirmPreparationPostpaid")
	}
	UserConfirmPreparationPostpaid = &response
	return nil
}
