package phx

import "fmt"

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
	result := p.GetApiInfo("requestESIM")
	if result.State == "C" {
		if UserRequestESIM != nil {
			return UserRequestESIM, nil
		}
		return nil, fmt.Errorf("no custom response set for requestESIM")
	}

	if result.State == "E" {
		return &RequestESIMResponse{
			ResultCode: "50000",
			ResultDesc: "Failed: requestESIM (1)",
		}, nil
	}

	return &RequestESIMResponse{
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
	}, nil
}
