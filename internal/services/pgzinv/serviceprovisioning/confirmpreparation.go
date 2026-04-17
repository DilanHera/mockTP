package serviceprovisioning

import pgzinvmodel "github.com/DilanHera/mockTP/internal/services/pgzinv/model"

type ConfirmPreparationRequestResourceItem struct {
	ResourceName   string `json:"resourceName" validate:"required"`
	OperationType  string `json:"operationType" validate:"required"`
	UserId         string `json:"userId" validate:"required"`
	SimSerialNo    string `json:"simSerialNo" validate:"required"`
	MobileNo       string `json:"mobileNo" validate:"required"`
	PrepNo         string `json:"prepNo" validate:"required"`
	JobId          string `json:"jobId" validate:"required"`
	ProjectName    string `json:"projectName" validate:"omitempty"`
	SimType        string `json:"simType" validate:"omitempty"`
	PackageRowId   string `json:"packageRowId" validate:"omitempty"`
	OfferingName   string `json:"offeringName" validate:"omitempty"`
	OfferingCode   string `json:"offeringCode" validate:"omitempty"`
	Package        string `json:"package" validate:"omitempty"`
	RegionCode     string `json:"regionCode" validate:"required"`
	EanCode        string `json:"eanCode" validate:"omitempty"`
	MaterialCode   string `json:"materialCode" validate:"omitempty"`
	ExpiryDate     string `json:"expiryDate" validate:"omitempty"`
	Aging          string `json:"aging" validate:"required"`
	SimProject     string `json:"simProject" validate:"required"`
	MatDesc        string `json:"matDesc" validate:"omitempty"`
	LocationCode   string `json:"locationCode" validate:"omitempty"`
	ProductType    string `json:"productType" validate:"omitempty"`
	PackageType    string `json:"packageType" validate:"omitempty"`
	SubPackageType string `json:"subPackageType" validate:"omitempty"`
	MatCode        string `json:"matCode" validate:"omitempty"`
	SourceSystem   string `json:"sourceSystem" validate:"omitempty"`
}

type ConfirmPreparationResponse struct {
	ResponseHeader   pgzinvmodel.ResponseHeader       `json:"responseHeader" validate:"required"`
	ResourceItemList []ConfirmPreparationResponseItem `json:"resourceItemList" validate:"required,dive"`
}

type ConfirmPreparationResponseItem struct {
	pgzinvmodel.ResourceItemListBase
	ConfirmPrepResponse []ConfirmPrepResponseItem `json:"confirmPrepResponse" validate:"omitempty"`
}

type ConfirmPrepResponseItem struct {
	SimSerialNo    string `json:"simSerialNo" validate:"required"`
	MobileNo       string `json:"mobileNo" validate:"required"`
	PrepNo         string `json:"prepNo" validate:"required"`
	ExpiryDate     string `json:"expiryDate" validate:"required"`
	RegionCode     string `json:"regionCode" validate:"required"`
	ClassifyCode   string `json:"classifyCode" validate:"required"`
	PatternNo      string `json:"patternNo" validate:"required"`
	NumberStatusTo string `json:"numberStatusTo" validate:"required"`
	SimType        string `json:"simType" validate:"required"`
	Package        string `json:"package" validate:"required"`
	PackageRowId   string `json:"packageRowId" validate:"required"`
	LuckyName      string `json:"luckyName" validate:"omitempty"`
	LuckyType      string `json:"luckyType" validate:"omitempty"`
	QRCodeInfo     string `json:"qrCodeInfo" validate:"omitempty"`
}

func (s *serviceProvisioning) ConfirmPreparation(input *ConfirmPreparationRequestResourceItem, requestHeader pgzinvmodel.HeaderServiceProvisioning) (ConfirmPreparationResponse, error) {
	if UserConfirmPreparationPrepaid != nil && input.ResourceName == "confirmPreparationPrepaid" {
		return *UserConfirmPreparationPrepaid, nil
	}
	if UserConfirmPreparationPostpaid != nil && input.ResourceName == "confirmPreparationPostpaid" {
		return *UserConfirmPreparationPostpaid, nil
	}
	response := &ConfirmPreparationResponse{}
	if s.app.ResponseState == "ERROR" {
		response = &ConfirmPreparationResponse{
			ResponseHeader: pgzinvmodel.ResponseHeader{
				ReTransmit:       "0",
				UserSys:          requestHeader.UserSys,
				ResourceGroupId:  requestHeader.ResourceGroupId,
				ResourceOrderId:  "DBSIPGSA001G-PGZINV-202303171437060271",
				ResultCode:       "50000",
				ResultDesc:       "Failed: confirmPreparationPostpaid(1) mobile and sim are status Registered.",
				DeveloperMessage: "",
			},
			ResourceItemList: []ConfirmPreparationResponseItem{
				{
					ResourceItemListBase: pgzinvmodel.ResourceItemListBase{
						ResourceName:           input.ResourceName,
						ResourceItemStatus:     "Failed",
						ResourceItemErrMessage: "mobile and sim are status Registered.",
						ErrorFlag:              "10",
					},
				},
			},
		}
	} else {
		response = &ConfirmPreparationResponse{
			ResponseHeader: pgzinvmodel.ResponseHeader{
				ResourceGroupId:  requestHeader.ResourceGroupId,
				ResourceOrderId:  "DBSIPGSA001G-PGZINV-202303171437060271",
				ReTransmit:       "0",
				UserSys:          requestHeader.UserSys,
				DeveloperMessage: "",
				ResultCode:       "20000",
				ResultDesc:       "Success",
			},
			ResourceItemList: []ConfirmPreparationResponseItem{
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
					ConfirmPrepResponse: []ConfirmPrepResponseItem{
						{
							SimSerialNo:    input.SimSerialNo,
							MobileNo:       input.MobileNo,
							PrepNo:         input.PrepNo,
							ExpiryDate:     "31/06/2026",
							RegionCode:     input.RegionCode,
							ClassifyCode:   "N",
							PatternNo:      "51",
							NumberStatusTo: "B",
							SimType:        input.SimType,
							Package:        input.Package,
							PackageRowId:   input.PackageRowId,
							LuckyName:      "Mor_AIS",
							LuckyType:      "Good Money & Love",
							QRCodeInfo:     "LPA:1$secsmsminiapp.eastcompeace.com$80D88923FADA3C76656D344AF",
						},
					},
				},
			},
		}
	}
	return *response, nil
}
