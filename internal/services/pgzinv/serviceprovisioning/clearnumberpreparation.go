package serviceprovisioning

import pgzinvmodel "github.com/DilanHera/mockTP/internal/services/pgzinv/model"

type ClearNumberPreparationRequestResourceItem struct {
	ResourceName     string `json:"resourceName" validate:"required"`
	OperationType    string `json:"operationType" validate:"required"`
	UserId           string `json:"userId" validate:"required"`
	NumberStatusFrom string `json:"numberStatusFrom" validate:"required"`
	NumberStatusTo   string `json:"numberStatusTo" validate:"required"`
	Key              string `json:"key" validate:"required"`
	MobileNo         string `json:"mobileNo" validate:"required"`
	Date             string `json:"date" validate:"required"`
}

type ClearNumberPreparationResponse struct {
	ResponseHeader   pgzinvmodel.ResponseHeader         `json:"responseHeader" validate:"required"`
	ResourceItemList []pgzinvmodel.ResourceItemListBase `json:"resourceItemList" validate:"required,dive"`
}

func (s *serviceProvisioning) ClearNumberPreparation(input *ClearNumberPreparationRequestResourceItem, requestHeader pgzinvmodel.HeaderServiceProvisioning) (*ClearNumberPreparationResponse, error) {
	if UserClearNumberPreparationPrepaid != nil && input.ResourceName == "clearNumberPreparationPrepaid" {
		return UserClearNumberPreparationPrepaid, nil
	}
	if UserClearNumberPreparationPostpaid != nil && input.ResourceName == "clearNumberPreparationPostpaid" {
		return UserClearNumberPreparationPostpaid, nil
	}
	response := &ClearNumberPreparationResponse{
		ResponseHeader: pgzinvmodel.ResponseHeader{
			ResourceGroupId:  requestHeader.ResourceGroupId,
			ResourceOrderId:  "DBSIPGSA001G-PGZINV-202303171437060271",
			ReTransmit:       "0",
			UserSys:          requestHeader.UserSys,
			DeveloperMessage: "",
			ResultCode:       "20000",
			ResultDesc:       "Success",
		},
	}
	return response, nil
}
