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
	NumberStatusTo   string `json:"numberStatusTo" validate:"required"`
	LocationCode     string `json:"locationCode" validate:"omitempty"`
	LuckyName        string `json:"luckyName" validate:"required"`
	LuckyType        string `json:"luckyType" validate:"required"`
	Quantity         string `json:"quantity" validate:"required"`
}

type LockNumberByCriteriaResponse struct {
	ResponseHeader   pgzinvmodel.ResponseHeader         `json:"responseHeader" validate:"required"`
	ResourceItemList []LockNumberByCriteriaResponseItem `json:"resourceItemList" validate:"required"`
}

type LockNumberByCriteriaResponseItem struct {
	pgzinvmodel.ResourceItemListBase
	RequestPrepResponse []RequestPrepResponseItem `json:"requestPrepResponse" validate:"required"`
}

type RequestPrepResponseItem struct {
	MobileNo string `json:"mobileNo" validate:"required"`
}

func (s *serviceProvisioning) LockNumberByCriteria(input *LockNumberByCriteriaRequestResourceItem) (LockNumberByCriteriaResponse, error) {
	response := LockNumberByCriteriaResponse{
		ResponseHeader: pgzinvmodel.ResponseHeader{
			ResourceGroupId: "123",
			ResourceOrderId: "123",
			ResultCode:      "123",
			ResultDesc:      "123",
		},
		ResourceItemList: []LockNumberByCriteriaResponseItem{
			{
				ResourceItemListBase: pgzinvmodel.ResourceItemListBase{
					ResourceName:           input.ResourceName,
					ResourceItemStatus:     "success",
					ErrorFlag:              "0",
					ResourceItemErrMessage: "",
					SpecialErrHandling: pgzinvmodel.SpecialErrHandling{
						SuppCode:             []string{},
						TaskKeyCondition:     []string{},
						TaskDeveloperMessage: []string{""},
					},
				},
				RequestPrepResponse: []RequestPrepResponseItem{},
			},
		},
	}
	quantity, err := strconv.Atoi(input.Quantity)
	if err != nil {
		return response, err
	}
	for i := 0; i < quantity; i++ {
		response.ResourceItemList[0].RequestPrepResponse = append(response.ResourceItemList[0].RequestPrepResponse, RequestPrepResponseItem{
			MobileNo: fmt.Sprintf("061%07d", rand.Intn(10000000)),
		})
	}
	return response, nil
}
