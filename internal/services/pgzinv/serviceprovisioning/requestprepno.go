package serviceprovisioning

import (
	"fmt"
	"strconv"

	pgzinvmodel "github.com/DilanHera/mockTP/internal/services/pgzinv/model"
)

type RequestPrepNoRequestResourceItem struct {
	ResourceName   string `json:"resourceName" validate:"required"`
	UserId         string `json:"userId" validate:"required"`
	Quantity       string `json:"quantity" validate:"required"`
	MachineNo      string `json:"machineNo" validate:"required"`
	JobId          string `json:"jobId" validate:"required"`
	Package        string `json:"package" validate:"omitempty"`
	PackageType    string `json:"packageType" validate:"omitempty"`
	SubPackageType string `json:"subPackageType" validate:"omitempty"`
	ProductType    string `json:"productType" validate:"omitempty"`
}

type RequestPrepNoResponse struct {
	ResponseHeader   pgzinvmodel.ResponseHeader  `json:"responseHeader" validate:"required"`
	ResourceItemList []RequestPrepNoResponseItem `json:"resourceItemList" validate:"required,dive"`
}

type RequestPrepNoResponseItem struct {
	pgzinvmodel.ResourceItemListBase
	PackageRowId string `json:"packageRowId" validate:"omitempty"`
	OfferingName string `json:"offeringName" validate:"omitempty"`
	OfferingCode string `json:"offeringCode" validate:"omitempty"`
	PrepNoFrom   string `json:"prepNoFrom" validate:"omitempty"`
	PrepNoTo     string `json:"prepNoTo" validate:"omitempty"`
}

func (s *serviceProvisioning) RequestPrepNo(input *RequestPrepNoRequestResourceItem, requestHeader pgzinvmodel.HeaderServiceProvisioning) (*RequestPrepNoResponse, error) {
	if GetResourceState(input.ResourceName) == "C" {
		if UserRequestPrepNoPrepaid != nil && input.ResourceName == "requestPrepNoPrepaid" {
			return UserRequestPrepNoPrepaid, nil
		}
		if UserRequestPrepNoPostpaid != nil && input.ResourceName == "requestPrepNoPostpaid" {
			return UserRequestPrepNoPostpaid, nil
		}
		return nil, fmt.Errorf("no custom response set for %s", input.ResourceName)
	}

	if GetResourceState(input.ResourceName) == "E" {
		return &RequestPrepNoResponse{
			ResponseHeader: pgzinvmodel.ResponseHeader{
				ResourceGroupId:  requestHeader.ResourceGroupId,
				ResourceOrderId:  "DBSIPGSA001G-PGZINV-202303171437060271",
				ReTransmit:       "0",
				UserSys:          requestHeader.UserSys,
				DeveloperMessage: "",
				ResultCode:       "50000",
				ResultDesc:       "Failed: " + input.ResourceName + " (1) preparation number not available.",
			},
			ResourceItemList: []RequestPrepNoResponseItem{
				{
					ResourceItemListBase: pgzinvmodel.ResourceItemListBase{
						ResourceName:           input.ResourceName,
						ResourceItemStatus:     "Failed",
						ErrorFlag:              "0",
						ResourceItemErrMessage: "preparation number not available.",
						SpecialErrHandling: pgzinvmodel.SpecialErrHandling{
							SuppCode:             []string{},
							TaskKeyCondition:     []string{},
							TaskDeveloperMessage: []string{},
						},
					},
				},
			},
		}, nil
	}

	response := &RequestPrepNoResponse{
		ResponseHeader: pgzinvmodel.ResponseHeader{
			ResourceGroupId:  requestHeader.ResourceGroupId,
			ResourceOrderId:  "DBSIPGSA001G-PGZINV-202303171437060271",
			ReTransmit:       "0",
			UserSys:          requestHeader.UserSys,
			DeveloperMessage: "",
			ResultCode:       "20000",
			ResultDesc:       "Success",
		},
		ResourceItemList: []RequestPrepNoResponseItem{
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
				PackageRowId: "1234567890",
				OfferingName: "Offering Name",
				OfferingCode: "Offering Code",
				PrepNoFrom:   "9300015000",
			},
		},
	}
	prepNoFrom, err := strconv.Atoi(response.ResourceItemList[0].PrepNoFrom)
	if err != nil {
		return response, err
	}
	quantity, err := strconv.Atoi(input.Quantity)
	if err != nil {
		return response, err
	}
	response.ResourceItemList[0].PrepNoTo = strconv.Itoa(prepNoFrom + (quantity - 1))
	return response, nil
}
