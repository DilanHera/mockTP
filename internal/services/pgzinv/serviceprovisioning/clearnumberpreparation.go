package serviceprovisioning

import (
	pgzinvmodel "github.com/DilanHera/mockTP/internal/services/pgzinv/model"
)

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
	HttpStatusCode   int                                `json:"-"`
}

func (s *serviceProvisioning) ClearNumberPreparation(input *ClearNumberPreparationRequestResourceItem, requestHeader pgzinvmodel.HeaderServiceProvisioning) (*ClearNumberPreparationResponse, error) {
	res := ClearNumberPreparationResponse{}
	result, err := s.app.Service.GetApiInfo(input.ResourceName, &res)
	if result.State == "C" {
		if err != nil {
			return nil, err
		}
		res.HttpStatusCode = result.HttpCode
		return &res, nil
	}

	if result.State == "E" {
		return &ClearNumberPreparationResponse{
			ResponseHeader: pgzinvmodel.ResponseHeader{
				ResourceGroupId:  requestHeader.ResourceGroupId,
				ResourceOrderId:  "DBSIPGSA001G-PGZINV-202303171437060271",
				ReTransmit:       "0",
				UserSys:          requestHeader.UserSys,
				DeveloperMessage: "",
				ResultCode:       "50000",
				ResultDesc:       "Failed: " + input.ResourceName + " (1)",
			},
			ResourceItemList: []pgzinvmodel.ResourceItemListBase{
				{
					ResourceName:           input.ResourceName,
					ResourceItemStatus:     "Failed",
					ErrorFlag:              "0",
					ResourceItemErrMessage: "Failed: " + input.ResourceName + " (1)",
					SpecialErrHandling: pgzinvmodel.SpecialErrHandling{
						SuppCode:             []string{},
						TaskKeyCondition:     []string{},
						TaskDeveloperMessage: []string{},
					},
				},
			},
			HttpStatusCode: 500,
		}, nil
	}

	return &ClearNumberPreparationResponse{
		ResponseHeader: pgzinvmodel.ResponseHeader{
			ResourceGroupId:  requestHeader.ResourceGroupId,
			ResourceOrderId:  "DBSIPGSA001G-PGZINV-202303171437060271",
			ReTransmit:       "0",
			UserSys:          requestHeader.UserSys,
			DeveloperMessage: "",
			ResultCode:       "20000",
			ResultDesc:       "Success",
		},
		ResourceItemList: []pgzinvmodel.ResourceItemListBase{
			{
				ResourceName:           input.ResourceName,
				ResourceItemStatus:     "Success",
				ErrorFlag:              "1",
				ResourceItemErrMessage: "Success",
				SpecialErrHandling: pgzinvmodel.SpecialErrHandling{
					SuppCode:             []string{},
					TaskKeyCondition:     []string{},
					TaskDeveloperMessage: []string{},
				},
			},
		},
		HttpStatusCode: 200,
	}, nil
}
