package serviceprovisioning

import (
	pgzinvmodel "github.com/DilanHera/mockTP/internal/services/pgzinv/model"
)

type LockNumberByMobileRequestResourceItem struct {
	ResourceName     string `json:"resourceName"`
	UserId           string `json:"userId"`
	MobileNo         string `json:"mobileNo"`
	NumberStatusFrom string `json:"numberStatusFrom"`
	NumberStatusTo   string `json:"numberStatusTo"`
}

type LockNumberByMobileResponse struct {
	ResponseHeader   pgzinvmodel.ResponseHeader         `json:"responseHeader" validate:"required"`
	ResourceItemList []pgzinvmodel.ResourceItemListBase `json:"resourceItemList" validate:"required,dive"`
	HttpStatusCode   int                                `json:"-"`
}

func (s *serviceProvisioning) LockNumberByMobile(input *LockNumberByMobileRequestResourceItem, requestHeader pgzinvmodel.HeaderServiceProvisioning) (*LockNumberByMobileResponse, error) {
	res := LockNumberByMobileResponse{}
	result, err := s.app.Service.GetApiInfo(input.ResourceName, &res)
	if result.State == "C" {
		if err != nil {
			return nil, err
		}
		res.HttpStatusCode = result.HttpCode
		return &res, nil
	}

	if result.State == "E" {
		return &LockNumberByMobileResponse{
			ResponseHeader: pgzinvmodel.ResponseHeader{
				ResourceGroupId:  requestHeader.ResourceGroupId,
				ResourceOrderId:  "DBSIPGSA001G-PGZINV-202303171437060271",
				ReTransmit:       "0",
				UserSys:          requestHeader.UserSys,
				DeveloperMessage: "",
				ResultCode:       "50000",
				ResultDesc:       "Failed: " + input.ResourceName + " (1) mobile not found.",
			},
			ResourceItemList: []pgzinvmodel.ResourceItemListBase{
				{
					ResourceName:           input.ResourceName,
					ResourceItemStatus:     "Failed",
					ErrorFlag:              "0",
					ResourceItemErrMessage: "mobile not found.",
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

	response := &LockNumberByMobileResponse{
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
	}
	return response, nil
}
