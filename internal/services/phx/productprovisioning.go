package phx

type ProductProvisioningRequest struct {
	SimType            string             `json:"simType" validate:"omitempty"`
	KeyCorrelate       string             `json:"keyCorrelate" validate:"required"`
	PublicIdType       string             `json:"publicIdType" validate:"required"`
	PublicIdValue      string             `json:"publicIdValue" validate:"required"`
	ReferenceId        string             `json:"referenceId" validate:"omitempty"`
	SourceSystem       string             `json:"sourceSystem" validate:"required"`
	CustomerOrderType  string             `json:"customerOrderType" validate:"required"`
	ChargeType         string             `json:"chargeType" validate:"required"`
	Channel            string             `json:"channel" validate:"required"`
	BillingSystem      string             `json:"billingSystem" validate:"required"`
	RequestType        string             `json:"requestType" validate:"required"`
	SyncFlag           string             `json:"syncFlag" validate:"required"`
	ReasonCode         string             `json:"reasonCode" validate:"omitempty"`
	AscMobileNo        string             `json:"ascMobileNo" validate:"omitempty"`
	AscCode            string             `json:"ascCode" validate:"omitempty"`
	LocationCode       string             `json:"locationCode" validate:"required"`
	DealerId           string             `json:"dealerId" validate:"omitempty"`
	RegisterCode       string             `json:"registerCode" validate:"omitempty"`
	EmpId              string             `json:"empId" validate:"omitempty"`
	AccessNo           string             `json:"accessNo" validate:"omitempty"`
	SerialNo           string             `json:"serialNo" validate:"required"`
	Imsi               string             `json:"imsi" validate:"required"`
	SrId               string             `json:"srId" validate:"required"`
	BodyRequest        BodyRequest        `json:"bodyRequest" validate:"omitempty"`
	PhxRequestItemList []PhxRequestItemPP `json:"phxRequestItemList" validate:"required,min=1,dive"`
	UserName           string             `json:"userName" validate:"required"`
}

type BodyRequest struct {
	ESimProject   string `json:"eSimProject" validate:"omitempty"`
	CardVendor    string `json:"cardVendor" validate:"omitempty"`
	SecureKey     string `json:"secureKey" validate:"omitempty"`
	ChangeSimFlag string `json:"changeSimFlag" validate:"omitempty"`
}

type PhxRequestItemPP struct {
	ServiceOrderType    string `json:"serviceOrderType" validate:"required"`
	ServiceOrderSubType string `json:"serviceOrderSubType" validate:"required"`
	RequestId           string `json:"requestId" validate:"required"`
}

type ProductProvisioningResponse struct {
	ResultCode          string              `json:"resultCode"`
	DeveloperMessage    string              `json:"developerMessage"`
	BusinessCode        string              `json:"businessCode"`
	Status              string              `json:"status"`
	ResultDesc          string              `json:"resultDesc"`
	CustomerOrderId     string              `json:"customerOrderId"`
	PhxResponseItemList []PhxResponseItemPP `json:"phxResponseItemList"`
	HttpStatusCode      int                 `json:"-"`
}

type PhxResponseItemPP struct {
	Status              string `json:"status"`
	ServiceOrderId      string `json:"serviceOrderId"`
	RequestId           string `json:"requestId"`
	ServiceOrderType    string `json:"serviceOrderType"`
	ServiceOrderSubType string `json:"serviceOrderSubType"`
}

func (p *phx) ProductProvisioning(input *ProductProvisioningRequest) (*ProductProvisioningResponse, error) {
	res := ProductProvisioningResponse{}
	result, err := p.app.Service.GetApiInfo("productProvisioning", &res)
	if result.State == "C" {
		if err != nil {
			return nil, err
		}
		res.HttpStatusCode = result.HttpCode
		return &res, nil
	}

	if result.State == "E" {
		return &ProductProvisioningResponse{
			ResultCode: "50000",
			ResultDesc: "Failed: productProvisioning (1)",
			HttpStatusCode: 500,
		}, nil
	}

	return &ProductProvisioningResponse{
		ResultCode: "20000",
		ResultDesc: "Success",
		HttpStatusCode: 200,
	}, nil
}
