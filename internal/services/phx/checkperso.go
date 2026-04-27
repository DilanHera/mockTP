package phx

type CheckPersoRequest struct {
	PublicIdType        string `json:"publicIdType" validate:"required"`
	PublicIdValue       string `json:"publicIdValue" validate:"omitempty"`
	SerialNo            string `json:"serialNo" validate:"required"`
	SourceSystem        string `json:"sourceSystem" validate:"required"`
	ChargeType          string `json:"chargeType" validate:"required"`
	Channel             string `json:"channel" validate:"required"`
	BillingSystem       string `json:"billingSystem" validate:"required"`
	LocationCode        string `json:"locationCode" validate:"required"`
	UsageType           string `json:"usageType" validate:"omitempty"`
	RegionCode          string `json:"regionCode" validate:"omitempty"`
	NetworkCode         string `json:"networkCode" validate:"omitempty"`
	DeviceOS            string `json:"deviceOS" validate:"omitempty"`
	SimService          string `json:"simService" validate:"required"`
	SecureKey           string `json:"secureKey" validate:"required"`
	SimCommandKey       string `json:"simCommandKey" validate:"required"`
	ServiceOrderSubType string `json:"serviceOrderSubType" validate:"required"`
	UserName            string `json:"userName" validate:"required"`
}

type CheckPersoResponse struct {
	ResultCode       string          `json:"resultCode"`
	DeveloperMessage string          `json:"developerMessage"`
	ResultDesc       string          `json:"resultDesc"`
	ResultData       ResultDataPerso `json:"resultData"`
	HttpStatusCode   int             `json:"-"`
}

type ResultDataPerso struct {
	SimItem    SimItem  `json:"simItem"`
	ImsiItem   ImsiItem `json:"imsiItem"`
	SimCommand string   `json:"simCommand"`
	SrId       string   `json:"srId"`
}

type SimItem struct {
	Imsi                    string `json:"imsi"`
	ImsiRegion              string `json:"imsiRegion"`
	SerialNo                string `json:"serialNo"`
	RegionCode              string `json:"regionCode"`
	ParentRegion            string `json:"parentRegion"`
	UsageType               string `json:"usageType"`
	SimState                string `json:"simState"`
	LocationId              string `json:"locationId"`
	MappingType             string `json:"mappingType"`
	ReserveExpiryDate       string `json:"reserveExpiryDate"`
	ExpiryDate              string `json:"expiryDate"`
	RemovalDate             string `json:"removalDate"`
	CapacityCode            string `json:"capacityCode"`
	Pin1                    string `json:"pin1"`
	Pin2                    string `json:"pin2"`
	Puk1                    string `json:"puk1"`
	Puk2                    string `json:"puk2"`
	SimType                 string `json:"simType"`
	CardProfile             string `json:"cardProfile"`
	Key                     string `json:"key"`
	Kic                     string `json:"kic"`
	SimVersion              string `json:"simVersion"`
	OfferingCode            string `json:"offeringCode"`
	PrepMobileNo            string `json:"prepMobileNo"`
	CreatedDate             string `json:"createdDate"`
	CreatedBy               string `json:"createdBy"`
	LastModifiedTime        string `json:"lastModifiedTime"`
	LastModifiedBy          string `json:"lastModifiedBy"`
	StateDate               string `json:"stateDate"`
	PreviousState           string `json:"previousState"`
	PrepFlag                string `json:"prepFlag"`
	PrepNo                  string `json:"prepNo"`
	PackageId               string `json:"packageId"`
	SimBatchNo              string `json:"simBatchNo"`
	LotNo                   string `json:"lotNo"`
	AccValue                string `json:"accValue"`
	Adm1                    string `json:"adm1"`
	Agent                   string `json:"agent"`
	Algorithm               string `json:"algorithm"`
	CleansingFlag           string `json:"cleansingFlag"`
	DeleteFlag              string `json:"deleteFlag"`
	Eki                     string `json:"eki"`
	Eopc                    string `json:"eopc"`
	K4id                    string `json:"k4id"`
	KioskId                 string `json:"kioskId"`
	OtaKey                  string `json:"otaKey"`
	PrepDealerCode          string `json:"prepDealerCode"`
	PrepIdCard              string `json:"prepIdCard"`
	PreviousUsageType       string `json:"previousUsageType"`
	QrCodeInfo              string `json:"qrCodeInfo"`
	Reason                  string `json:"reason"`
	Remark                  string `json:"remark"`
	ReusedDate              string `json:"reusedDate"`
	ReusedType              string `json:"reusedType"`
	ServiceType             string `json:"serviceType"`
	SimCategory             string `json:"simCategory"`
	SimProductType          string `json:"simProductType"`
	SimProfile              string `json:"simProfile"`
	SubLocationName         string `json:"subLocationName"`
	WriteCounter            string `json:"writeCounter"`
	NetworkCode             string `json:"networkCode"`
	PackageTypeMasterPrepId string `json:"packageTypeMasterPrepId"`
	SubUsageType            string `json:"subUsageType"`
	DummyImsi               string `json:"dummyImsi"`
	Iccid                   string `json:"iccid"`
	SMDPAddress             string `json:"SMDPAddress"`
	MatchingId              string `json:"matchingId"`
	AssignedBy              string `json:"assignedBy"`
	AssignedDate            string `json:"assignedDate"`
	ReceivedBy              string `json:"receivedBy"`
	ReceivedDate            string `json:"receivedDate"`
	RejectedBy              string `json:"rejectedBy"`
	RejectedDate            string `json:"rejectedDate"`
	FirstOutExpiredDate     string `json:"firstOutExpiredDate"`
	RowId                   string `json:"rowId"`
}

type ImsiItem struct {
	Channel          string `json:"channel"`
	CreatedBy        string `json:"createdBy"`
	CreatedDate      string `json:"createdDate"`
	DealerCode       string `json:"dealerCode"`
	Imsi             string `json:"imsi"`
	ImsiCategory     string `json:"imsiCategory"`
	ImsiProductType  string `json:"imsiProductType"`
	ImsiState        string `json:"imsiState"`
	LastModifiedBy   string `json:"lastModifiedBy"`
	LastModifiedTime string `json:"lastModifiedTime"`
	LocationId       string `json:"locationId"`
	NetworkCode      string `json:"networkCode"`
	PreviousState    string `json:"previousState"`
	Reason           string `json:"reason"`
	RegionCode       string `json:"regionCode"`
	Remark           string `json:"remark"`
	SerialNo         string `json:"serialNo"`
	ServiceType      string `json:"serviceType"`
	UsageFlag        string `json:"usageFlag"`
	UsageType        string `json:"usageType"`
	AccValue         string `json:"accValue"`
}

func (p *phx) CheckPerso(input *CheckPersoRequest) (*CheckPersoResponse, error) {
	res := CheckPersoResponse{}
	result, err := p.app.Service.GetApiInfo("checkPerso", &res)
	if result.State == "C" {
		if err != nil {
			return nil, err
		}
		res.HttpStatusCode = result.HttpCode
		return &res, nil
	}

	if result.State == "E" {
		return &CheckPersoResponse{
			ResultCode:       "50000",
			DeveloperMessage: "Mock Error",
			ResultDesc:       "Mock Error",
			HttpStatusCode:   500,
		}, nil
	}

	return &CheckPersoResponse{
		ResultCode:       "20000",
		DeveloperMessage: "",
		ResultDesc:       "Success",
		HttpStatusCode:   200,
		ResultData: ResultDataPerso{
			SimItem: SimItem{
				Imsi:                    "",
				ImsiRegion:              "",
				SerialNo:                "1741691020750",
				RegionCode:              "",
				ParentRegion:            "",
				UsageType:               "Dynamic_DSA",
				SimState:                "PendingPerso",
				LocationId:              "1-19DZ6CB1",
				MappingType:             "",
				ReserveExpiryDate:       "",
				ExpiryDate:              "",
				RemovalDate:             "",
				CapacityCode:            "128k",
				Pin1:                    "1234",
				Pin2:                    "3162",
				Puk1:                    "10143174",
				Puk2:                    "29506931",
				SimType:                 "Service",
				CardProfile:             "ECAWNE1210",
				Key:                     "9F31EE9532241CE09911BF6C470BB890",
				Kic:                     "643252B5EA41E94E99BFDC2E36F958D2",
				SimVersion:              "3G",
				OfferingCode:            "",
				PrepMobileNo:            "",
				CreatedDate:             "20171219040722+0700",
				CreatedBy:               "SYSADMIN",
				LastModifiedTime:        "20200820132152+0700",
				LastModifiedBy:          "PGZINV",
				StateDate:               "",
				PreviousState:           "",
				PrepFlag:                "",
				PrepNo:                  "",
				PackageId:               "",
				SimBatchNo:              "",
				LotNo:                   "",
				AccValue:                "",
				Adm1:                    "3E5C05BAD037D1EE",
				Agent:                   "",
				Algorithm:               "128",
				CleansingFlag:           "N",
				DeleteFlag:              "",
				Eki:                     "102FC8DC6801B3983534CCE39C103116",
				Eopc:                    "3",
				K4id:                    "184",
				KioskId:                 "",
				OtaKey:                  "9F31EE9532241CE09911BF6C470BB890",
				PrepDealerCode:          "",
				PrepIdCard:              "",
				PreviousUsageType:       "",
				QrCodeInfo:              "",
				Reason:                  "8a7cc019740a755001740a87e54d004b",
				Remark:                  "SXPHX202008201321520388791500",
				ReusedDate:              "",
				ReusedType:              "",
				ServiceType:             "",
				SimCategory:             "",
				SimProductType:          "",
				SimProfile:              "1710",
				SubLocationName:         "",
				WriteCounter:            "0",
				NetworkCode:             "03",
				PackageTypeMasterPrepId: "",
				SubUsageType:            "",
				DummyImsi:               "530039000000231",
				Iccid:                   "",
				SMDPAddress:             "",
				MatchingId:              "",
				AssignedBy:              "",
				AssignedDate:            "",
				ReceivedBy:              "",
				ReceivedDate:            "",
				RejectedBy:              "",
				RejectedDate:            "",
				FirstOutExpiredDate:     "",
				RowId:                   "PHX202008201321530033361500",
			},
			ImsiItem: ImsiItem{
				Channel:          "",
				CreatedBy:        "PGZINV",
				CreatedDate:      "20200820132152+0700",
				DealerCode:       "WDS",
				Imsi:             "160033000001037",
				ImsiCategory:     "X30P30",
				ImsiProductType:  "3PP",
				ImsiState:        "PendingPerso",
				LastModifiedBy:   "PGZINV",
				LastModifiedTime: "20200820132152+0700",
				LocationId:       "",
				NetworkCode:      "03",
				PreviousState:    "",
				Reason:           "8a7cc019740a755001740a87e54d004b",
				RegionCode:       "X330",
				Remark:           "SXPHX202008201321520388791500",
				SerialNo:         "",
				ServiceType:      "Pre-paid",
				UsageFlag:        "",
				UsageType:        "Dynamic_STD",
				AccValue:         "0004",
			},
			SimCommand: "TRUE|||{{serialNo}}|||ChgMHR4UEAwXAxgRAh4TAREMDR4HFxkQCwkPGgcJChQzDsnQ8m+iSuKFV6lZ7hH381h6pHQhM8T5FbSYR3webZXFLitRwEZ5GLNwLXbGMPRV5Roeodh7wqOgOIMlnc6b/J3KDOdKpzA8PZn3kPqFXMqhiEKYs4oVCCw0Uzq3kD7vhJoFBajH1UbTaf0lPobnRDKYnDvlN8a0MkXv3ZqrCRMZ+07ae07kLX3YaYmwK4ANPQ0fcFgpv/iTH0Ks/5DfLH0ZcLB2H2jIL5D+YDkZpV/s1xcGcaXCBXclSEiVsGrgeWoHBlYRBYHPN/5oVjy95As8bM/gD//O8YDfPUp8ogthYfH3+jLfXVuNeCvYyaBouqhd6/AYCVfuSE5Nt9fUF1K1bRIChLnFrQ6wNlcPF8iQW5XWrfE65HWZkDJ7oA1b8vNDUtYQCdVDBpOq3Ks75PS+lKGWvUfOvCHREoUwl4Xs1WM0GNFRiFeADw==|||$2a$10$ZuxMbS9DjRpZOag0hP9Lju|||FNMgZuPC/rUUBei/HJpK1u/TMqSWjHgJ9Dvgv8qA21a0+FpQWUW2/HOb99b+dFM1T4If/WI33FggLQIawuRRhIlM0/3s6xR+6Jb3nWa0ZFoc9AN7e/0XOauB4bUo7e7creOK6kJ+PMCauuWo8XWh8TGRvIYJcoseAVn47f+SphPjw8LyPlXhlgiCKcFLts8ovEo7C351mcfXZPIvmgBKGZR49I6VL6MumOmAG7+FY/BeW8dUZfQPj/f0nZnR7tGed7dv7Erajce1CqjWRDbooeSteyIDD1Z2HPXClHhIirMwlWJ2YC7eaxXEQf+ZnIRroHPwOGzPvSFoWxdEbsBN6A==",
			SrId:       "8a7cc019740a755001740a87e54d004b",
		},
	}, nil
}
