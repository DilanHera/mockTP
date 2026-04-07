package phx

type NewRegistrationRequest struct {
	KeyCorrelate       string           `json:"keyCorrelate" validate:"required"`
	PublicIdType       string           `json:"publicIdType" validate:"required"`
	PublicIdValue      string           `json:"publicIdValue" validate:"required"`
	SourceSystem       string           `json:"sourceSystem" validate:"required"`
	CustomerOrderType  string           `json:"customerOrderType" validate:"required"`
	UserName           string           `json:"userName" validate:"required"`
	ChargeType         string           `json:"chargeType" validate:"required"`
	Channel            string           `json:"channel" validate:"required"`
	BillingSystem      string           `json:"billingSystem" validate:"required"`
	RequestType        string           `json:"requestType" validate:"required"`
	LocationCode       string           `json:"locationCode" validate:"required"`
	SyncFlag           string           `json:"syncFlag" validate:"required"`
	CustomerItem       CustomerItem     `json:"customerItem" validate:"required"`
	PhxRequestItemList []PhxRequestItem `json:"phxRequestItemList" validate:"required,dive"`
}

type CustomerItem struct {
	IdCardNo            string `json:"idCardNo" validate:"required"`
	IdCardTypeId        string `json:"idCardTypeId" validate:"required"`
	CustomerFirstName   string `json:"customerFirstName" validate:"required"`
	CustomerLastName    string `json:"customerLastName" validate:"required"`
	CustomerFirstNameEN string `json:"customerFirstNameEN" validate:"required"`
	CustomerLastNameEN  string `json:"customerLastNameEN" validate:"required"`
	IdenDate            string `json:"idenDate" validate:"required"`
	IdentLocationCode   string `json:"identLocationCode" validate:"required"`
}

type PhxRequestItem struct {
	RequestId           string `json:"requestId" validate:"required"`
	ServiceOrderSubType string `json:"serviceOrderSubType" validate:"required"`
	ServiceOrderType    string `json:"serviceOrderType" validate:"required"`
}

type NewRegistrationResponse struct {
	ResultCode string `json:"resultCode" validate:"required"`
	ResultDesc string `json:"resultDesc" validate:"required"`
}

func (p *phx) NewRegistration(input *NewRegistrationRequest) (*NewRegistrationResponse, error) {
	if UserNewRegistration != nil {
		return UserNewRegistration, nil
	}
	response := &NewRegistrationResponse{
		ResultCode: "20000",
		ResultDesc: "Success",
	}
	UserNewRegistration = response
	return response, nil
}
