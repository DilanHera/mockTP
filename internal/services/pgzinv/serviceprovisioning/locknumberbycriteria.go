package serviceprovisioning

import (
	"fmt"
	"math/rand"
	"strconv"

	pgzinvmodel "github.com/DilanHera/mockTP/internal/services/pgzinv/model"
)

type LockNumberByCriteriaRequestResourceItem struct {
	ResourceName     string `json:"resourceName" validate:"required"`
	UserId           string `json:"userId" validate:"required"`
	OriginalOperator string `json:"originalOperator" validate:"required"`
	ChargeNode       string `json:"chargeNode" validate:"required"`
	CategoryCode     string `json:"categoryCode" validate:"required"`
	ProductType      string `json:"productType" validate:"required"`
	RegionCode       string `json:"regionCode" validate:"required"`
	NumberPattern    string `json:"numberPattern" validate:"required"`
	ClassifyCode     string `json:"classifyCode" validate:"required"`
	NumberStatusFrom string `json:"numberStatusFrom" validate:"required"`
	NumberStatusTo   string `json:"numberStatusTo" validate:"omitempty"`
	LocationCode     string `json:"locationCode" validate:"omitempty"`
	LuckyName        string `json:"luckyName" validate:"omitempty"`
	LuckyType        string `json:"luckyType" validate:"omitempty"`
	Quantity         string `json:"quantity" validate:"required"`
}

type LockNumberByCriteriaResponse struct {
	ResponseHeader   pgzinvmodel.ResponseHeader         `json:"responseHeader" validate:"required"`
	ResourceItemList []LockNumberByCriteriaResponseItem `json:"resourceItemList" validate:"required,dive"`
}

type LockNumberByCriteriaResponseItem struct {
	pgzinvmodel.ResourceItemListBase
	Key                 string                  `json:"key" validate:"omitempty"`
	RequestPrepResponse []RequestPrepResponseItem `json:"requestPrepResponse" validate:"required,min=1,dive"`
}

type RequestPrepResponseItem struct {
	MobileNo string `json:"mobileNo" validate:"required"`
}

func (s *serviceProvisioning) LockNumberByCriteria(input *LockNumberByCriteriaRequestResourceItem, requestHeader pgzinvmodel.HeaderServiceProvisioning) (LockNumberByCriteriaResponse, error) {
	if input.ResourceName == "lockNumberByCriteriaPrepaid" && UserLockNumberByCriteriaPrepaid != nil {
		return *UserLockNumberByCriteriaPrepaid, nil
	}
	if input.ResourceName == "lockNumberByCriteriaPostpaid" && UserLockNumberByCriteriaPostpaid != nil {
		return *UserLockNumberByCriteriaPostpaid, nil
	}
	response := LockNumberByCriteriaResponse{
		ResponseHeader: pgzinvmodel.ResponseHeader{
			ResourceGroupId:  requestHeader.ResourceGroupId,
			ResourceOrderId:  "DBSIPGSA001G-PGZINV-202303171437060271",
			ReTransmit:       "0",
			UserSys:          requestHeader.UserSys,
			DeveloperMessage: "",
			ResultCode:       "20000",
			ResultDesc:       "Success",
		},
		ResourceItemList: []LockNumberByCriteriaResponseItem{
			{
				ResourceItemListBase: pgzinvmodel.ResourceItemListBase{
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
				Key:                 "1234567",
				RequestPrepResponse: []RequestPrepResponseItem{},
			},
		},
	}
	quantity, err := strconv.Atoi(input.Quantity)
	if err != nil {
		return response, err
	}
	for range quantity {
		response.ResourceItemList[0].RequestPrepResponse = append(response.ResourceItemList[0].RequestPrepResponse, RequestPrepResponseItem{
			MobileNo: fmt.Sprintf("061%07d", rand.Intn(10000000)),
		})
	}
	return response, nil
}
