package phx

type RequestESIMRequest struct {
	Msisdn        string `json:"msisdn"`
	ESimProject   string `json:"eSimProject"`
	SourceSystem  string `json:"sourceSystem"`
	UserName      string `json:"userName"`
	LocationCode  string `json:"locationCode"`
	SecureKey     string `json:"secureKey"`
	ChargeType    string `json:"chargeType"`
	BillingSystem string `json:"billingSystem"`
	SimService    string `json:"simService"`
	SimType       string `json:"simType"`
	PersoFlag     string `json:"persoFlag"`
	ReferenceId   string `json:"referenceId"`
	Channel       string `json:"channel"`
}

type RequestESIMResponse struct {
	ResultCode string     `json:"resultCode"`
	ResultDesc string     `json:"resultDesc"`
	ResultData ResultData `json:"resultData"`
}

type ResultData struct {
	NewSimItem NewSimItem `json:"newSimItem"`
}

type NewSimItem struct {
	Imsi       string `json:"imsi"`
	QRCodeInfo string `json:"qrCodeInfo"`
	RegionCode string `json:"regionCode"`
	SerialNo   string `json:"serialNo"`
}

func (p *phx) RequestESIM(input *RequestESIMRequest) (*RequestESIMResponse, error) {
	if UserRequestESIM != nil {
		return UserRequestESIM, nil
	}
	response := &RequestESIMResponse{
		ResultCode: "20000",
		ResultDesc: "Success",
		ResultData: ResultData{
			NewSimItem: NewSimItem{
				Imsi:       "1234567890",
				QRCodeInfo: "1234567890",
				RegionCode: "1234567890",
				SerialNo:   "1234567890",
			},
		},
	}
	UserRequestESIM = response
	return response, nil
}