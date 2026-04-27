package phx

type RequestESIMRequest struct {
	Msisdn        string `json:"msisdn" validate:"required"`
	ESimProject   string `json:"eSimProject" validate:"required"`
	SourceSystem  string `json:"sourceSystem" validate:"required"`
	UserName      string `json:"userName" validate:"required"`
	LocationCode  string `json:"locationCode" validate:"required"`
	SecureKey     string `json:"secureKey" validate:"required"`
	ChargeType    string `json:"chargeType" validate:"required"`
	BillingSystem string `json:"billingSystem" validate:"required"`
	SimService    string `json:"simService" validate:"required"`
	SimType       string `json:"simType" validate:"required"`
	PersoFlag     string `json:"persoFlag" validate:"required"`
	ReferenceId   string `json:"referenceId" validate:"omitempty"`
	Channel       string `json:"channel" validate:"required"`
}

type RequestESIMResponse struct {
	ResultCode     string     `json:"resultCode" validate:"required"`
	ResultDesc     string     `json:"resultDesc" validate:"required"`
	ResultData     ResultData `json:"resultData" validate:"omitempty"`
	HttpStatusCode int        `json:"-"`
}

type ResultData struct {
	NewSimItem NewSimItem `json:"newSimItem" validate:"required"`
}

type NewSimItem struct {
	Imsi       string `json:"imsi" validate:"required"`
	QRCodeInfo string `json:"qrCodeInfo" validate:"required"`
	RegionCode string `json:"regionCode" validate:"required"`
	SerialNo   string `json:"serialNo" validate:"required"`
}

func (p *phx) RequestESIM(input *RequestESIMRequest) (*RequestESIMResponse, error) {
	res := RequestESIMResponse{}
	result, err := p.app.Service.GetApiInfo("requestESIM", &res)
	if result.State == "C" {
		if err != nil {
			return nil, err
		}
		res.HttpStatusCode = result.HttpCode
		return &res, nil
	}

	if result.State == "E" {
		return &RequestESIMResponse{
			ResultCode:     "50000",
			ResultDesc:     "Failed: requestESIM (1)",
			HttpStatusCode: 500,
		}, nil
	}

	return &RequestESIMResponse{
		ResultCode:     "20000",
		ResultDesc:     "Success",
		HttpStatusCode: 200,
		ResultData: ResultData{
			NewSimItem: NewSimItem{
				Imsi:       "1234567890",
				QRCodeInfo: "1234567890",
				RegionCode: "1234567890",
				SerialNo:   "1234567890",
			},
		},
	}, nil
}
