package serviceprovisioning

import pgzinvmodel "github.com/DilanHera/mockTP/internal/services/pgzinv/model"

type QuerySimInfoRequestResourceItem struct {
	ResourceName string `json:"resourceName" validate:"required"`
	UserId       string `json:"userId" validate:"required"`
	SimSerialNo  string `json:"simSerialNo" validate:"required"`
	ChargeType   string `json:"chargeType" validate:"required"`
}

type QuerySimInfoResponse struct {
	ResponseHeader   pgzinvmodel.ResponseHeader `json:"responseHeader" validate:"required"`
	ResourceItemList []QuerySimInfoResponseItem `json:"resourceItemList" validate:"required,dive"`
}

type QuerySimInfoResponseItem struct {
	pgzinvmodel.ResourceItemListBase
	SimSerialNoList []SimSerialNoListItem `json:"simSerialNoList,omitempty" validate:"required,dive"`
}

type SimSerialNoListItem struct {
	SimSerialNo       string `json:"simSerialNo" validate:"required"`
	PreparationDate   string `json:"preparationDate" validate:"optional"`
	SimSerialNoStatus string `json:"simSerialNoStatus" validate:"required"`
	StatusDate        string `json:"statusDate" validate:"optional"`
	ExpiryDate        string `json:"expiryDate" validate:"required"`
	PackageNo         string `json:"packageNo" validate:"optional"`
	SubRegion         string `json:"subRegion" validate:"optional"`
	PackType          string `json:"packType" validate:"optional"`
	SubPackType       string `json:"subPackType" validate:"optional"`
	MobileNo          string `json:"mobileNo" validate:"optional"`
	MobileNoStatus    string `json:"mobileNoStatus" validate:"optional"`
	NumberClass       string `json:"numberClass" validate:"optional"`
	NumberPattern     string `json:"numberPattern" validate:"optional"`
	LuckyName         string `json:"luckyName" validate:"required"`
	LuckyType         string `json:"luckyType" validate:"required"`
	QRCodeInfo        string `json:"qrCodeInfo" validate:"required"`
	Material          string `json:"material" validate:"optional"`
}

func (s *serviceProvisioning) QuerySimInfo(input *QuerySimInfoRequestResourceItem, requestHeader pgzinvmodel.HeaderServiceProvisioning) (*QuerySimInfoResponse, error) {
	if UserQuerySimInfo != nil {
		return UserQuerySimInfo, nil
	}
	var response *QuerySimInfoResponse
	if s.app.ResponseState == "ERROR" {
		response = &QuerySimInfoResponse{
			ResponseHeader: pgzinvmodel.ResponseHeader{
				ResourceGroupId:  requestHeader.ResourceGroupId,
				ResourceOrderId:  "DBSIPGSA001G-PGZINV-202303171437060271",
				ReTransmit:       "0",
				UserSys:          requestHeader.UserSys,
				DeveloperMessage: "",
				ResultCode:       "50000",
				ResultDesc:       "Failed: querySimInfo (1)",
			},
			ResourceItemList: []QuerySimInfoResponseItem{
				{
					ResourceItemListBase: pgzinvmodel.ResourceItemListBase{
						ResourceName:           input.ResourceName,
						ResourceItemStatus:     "Failed",
						ErrorFlag:              "1",
						ResourceItemErrMessage: "Failed: querySimInfo (1)",
						SpecialErrHandling: pgzinvmodel.SpecialErrHandling{
							SuppCode:             []string{},
							TaskKeyCondition:     []string{},
							TaskDeveloperMessage: []string{},
						},
					},
				},
			},
		}
	} else {
		response = &QuerySimInfoResponse{
			ResponseHeader: pgzinvmodel.ResponseHeader{
				ResourceGroupId:  requestHeader.ResourceGroupId,
				ResourceOrderId:  "DBSIPGSA001G-PGZINV-202303171437060271",
				ReTransmit:       "0",
				UserSys:          requestHeader.UserSys,
				DeveloperMessage: "",
				ResultCode:       "20000",
				ResultDesc:       "Success",
			},
			ResourceItemList: []QuerySimInfoResponseItem{
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
					SimSerialNoList: []SimSerialNoListItem{
						{
							SimSerialNo:       input.SimSerialNo,
							PreparationDate:   "28/02/202410:07:06",
							SimSerialNoStatus: "Reserved",
							StatusDate:        "28/02/202410:07:06",
							ExpiryDate:        "28/02/202623:59:59",
							PackageNo:         "9991425266",
							SubRegion:         "C301",
							PackType:          "X",
							SubPackType:       "K1",
							MobileNo:          "0983044861",
							MobileNoStatus:    "Reserved",
							NumberClass:       "Normal",
							NumberPattern:     "77",
							LuckyName:         "Mor_MAN",
							LuckyType:         "GoodLove",
							QRCodeInfo:        "LPA:1$secsmsminiapp.eastcompeace.com$80D88923FADA3C76656D344AF",
							Material:          "1000022401",
						},
					},
				},
			},
		}
	}
	return response, nil
}
