package serviceprovisioning

import pgzinvmodel "github.com/DilanHera/mockTP/internal/services/pgzinv/model"

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
}

func (s *serviceProvisioning) LockNumberByMobile(input *LockNumberByMobileRequestResourceItem, requestHeader pgzinvmodel.HeaderServiceProvisioning) (*LockNumberByMobileResponse, error) {
	if UserLockNumberByMobilePrepaid != nil && input.ResourceName == "lockNumberByMobilePrepaid" {
		return UserLockNumberByMobilePrepaid, nil
	}
	if UserLockNumberByMobilePostpaid != nil && input.ResourceName == "lockNumberByMobilePostpaid" {
		return UserLockNumberByMobilePostpaid, nil
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
	}
	return response, nil
}
