package serviceprovisioning

import (
	"math/rand"
	"strconv"

	pgzinvmodel "github.com/DilanHera/mockTP/internal/services/pgzinv/model"
)

type PreMatchingRequestResourceItem struct {
	ResourceName    string `json:"resourceName" validate:"required"`
	RequestType     string `json:"requestType" validate:"required"`
	ActionType      string `json:"actionType" validate:"required"`
	PreMatchingType string `json:"preMatchingType" validate:"required"`
	UserId          string `json:"userId" validate:"required"`
	SimSerialNo     string `json:"simSerialNo" validate:"required"`
	MobileNo        string `json:"mobileNo" validate:"required"`
	SubRegion       string `json:"subRegion" validate:"required"`
	MobileStatus    string `json:"mobileStatus" validate:"required"`
	MatCode         string `json:"matCode" validate:"required"`
	EanCode         string `json:"eanCode" validate:"required"`
	Product         string `json:"product" validate:"required"`
	ProjectName     string `json:"projectName" validate:"required"`
}

type PreMatchingResponse struct {
	ResponseHeader   pgzinvmodel.ResponseHeader `json:"responseHeader" validate:"required"`
	ResourceItemList []PreMatchingResourceItem  `json:"resourceItemList" validate:"required,dive"`
	HttpStatusCode   int                        `json:"-"`
}

type PreMatchingResourceItem struct {
	pgzinvmodel.ResourceItemListBase
	PreMatchingResult PreMatchingResult `json:"preMatchingResult" validate:"omitempty"`
}

type PreMatchingResult struct {
	SimSerialNo string `json:"simSerialNo" validate:"required"`
	QRCodeInfo  string `json:"qrCodeInfo" validate:"omitempty"`
	MobileNo    string `json:"mobileNo" validate:"required"`
	PrepNo      string `json:"prepNo" validate:"required"`
	ExpiryDate  string `json:"expiryDate" validate:"required"`
	EanCode     string `json:"eanCode" validate:"required"`
	MatCode     string `json:"matCode" validate:"required"`
}

func (s *serviceProvisioning) PreMatching(input *PreMatchingRequestResourceItem, requestHeader pgzinvmodel.HeaderServiceProvisioning) (PreMatchingResponse, error) {
	res := PreMatchingResponse{}
	result, err := s.app.Service.GetApiInfo(input.ResourceName, &res)
	if result.State == "C" {
		if err != nil {
			return PreMatchingResponse{}, err
		}
		res.HttpStatusCode = result.HttpCode
		return res, nil
	}
	if result.State == "E" {
		return PreMatchingResponse{
			ResponseHeader: pgzinvmodel.ResponseHeader{
				CustomerOrderType: requestHeader.CustomerOrderType,
				ReTransmit:        "0",
				UserSys:           requestHeader.UserSys,
				ResourceGroupId:   requestHeader.ResourceGroupId,
				ResourceOrderId:   "DBSIPGSA001G-PGZINV-202303171437060271",
				ResultCode:        "50000",
				ResultDesc:        "Failed: preMatching failed.",
				DeveloperMessage:  "",
			},
			ResourceItemList: []PreMatchingResourceItem{
				{
					ResourceItemListBase: pgzinvmodel.ResourceItemListBase{
						ResourceName:           input.ResourceName,
						ResourceItemStatus:     "Failed",
						ResourceItemErrMessage: "failed.",
						ErrorFlag:              "10",
						SpecialErrHandling: pgzinvmodel.SpecialErrHandling{
							SuppCode:             []string{},
							TaskKeyCondition:     []string{},
							TaskDeveloperMessage: []string{},
						},
					},
				},
			},
			HttpStatusCode: 500,
		}, nil
	}
	if result.State == "T" {
		s.app.Helper.Delay(30)
	}
	return PreMatchingResponse{
		ResponseHeader: pgzinvmodel.ResponseHeader{
			ResourceGroupId:  requestHeader.ResourceGroupId,
			ResourceOrderId:  "DBSIPGSA001G-PGZINV-202303171437060271",
			ReTransmit:       "0",
			UserSys:          requestHeader.UserSys,
			DeveloperMessage: "",
			ResultCode:       "20000",
			ResultDesc:       "Success",
		},
		ResourceItemList: []PreMatchingResourceItem{
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
				PreMatchingResult: PreMatchingResult{
					SimSerialNo: input.SimSerialNo,
					MobileNo:    input.MobileNo,
					PrepNo:      strconv.Itoa(rand.Intn(9000000000) + 1000000000),
					ExpiryDate:  "31/06/2027",
					EanCode:     "9999999999999",
					MatCode:     "1000000101",
					QRCodeInfo:  "LPA:1$secsmsminiapp.eastcompeace.com$80D88923FADA3C76656D344AF",
				},
			},
		},
		HttpStatusCode: 200,
	}, nil
}
