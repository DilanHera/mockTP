package serviceprovisioning

import (
	"encoding/json"
	"fmt"
)

var UserLockNumberByCriteriaPrepaid *LockNumberByCriteriaResponse
var UserLockNumberByCriteriaPostpaid *LockNumberByCriteriaResponse

var UserLockNumberByMobilePrepaid *LockNumberByMobileResponse
var UserLockNumberByMobilePostpaid *LockNumberByMobileResponse

var UserClearNumberPreparationPrepaid *ClearNumberPreparationResponse
var UserClearNumberPreparationPostpaid *ClearNumberPreparationResponse

var UserQuerySimInfo *QuerySimInfoResponse

var UserRequestPrepNoPrepaid *RequestPrepNoResponse
var UserRequestPrepNoPostpaid *RequestPrepNoResponse

var UserConfirmPreparationPrepaid *ConfirmPreparationResponse
var UserConfirmPreparationPostpaid *ConfirmPreparationResponse

func (s *serviceProvisioning) SetUserLockNumberByCriteriaPrepaid(jsonData json.RawMessage) error {
	response := LockNumberByCriteriaResponse{}
	err := json.Unmarshal(jsonData, &response)
	if err != nil {
		return fmt.Errorf("failed to unmarshal: %w", err)
	}
	if len(response.ResourceItemList) > 0 && response.ResourceItemList[0].ResourceName != "lockNumberByCriteriaPrepaid" {
		return fmt.Errorf("wrong resource name, expected: lockNumberByCriteriaPrepaid")
	}
	// err := s.app.Helper.UnmarshalAndValidate(jsonData, &response)
	// if err != nil {
	// 	return fmt.Errorf("failed to validate: %w", err)
	// }
	UserLockNumberByCriteriaPrepaid = &response
	return nil
}

func (s *serviceProvisioning) SetUserLockNumberByCriteriaPostpaid(jsonData json.RawMessage) error {
	response := LockNumberByCriteriaResponse{}
	err := json.Unmarshal(jsonData, &response)
	if err != nil {
		return fmt.Errorf("failed to unmarshal: %w", err)
	}
	if len(response.ResourceItemList) > 0 && response.ResourceItemList[0].ResourceName != "lockNumberByCriteriaPostpaid" {
		return fmt.Errorf("wrong resource name, expected: lockNumberByCriteriaPostpaid")
	}
	// err := s.app.Helper.UnmarshalAndValidate(jsonData, &response)
	// if err != nil {
	// 	return fmt.Errorf("failed to validate: %w", err)
	// }
	UserLockNumberByCriteriaPostpaid = &response
	return nil
}

func (s *serviceProvisioning) SetUserLockNumberByMobilePrepaid(jsonData json.RawMessage) error {
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
