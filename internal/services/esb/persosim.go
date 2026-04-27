package esb

import "fmt"

type PersosimRequest struct {
	TransactionID string `json:"transactionId" validate:"required"`
	OrderNo       string `json:"orderNo" validate:"required"`
	MobileNo      string `json:"mobileNo" validate:"required"`
	TrackingNo    string `json:"trackingNo" validate:"required"`
	SerialNo      string `json:"serialNo" validate:"required"`
	Imsi          string `json:"imsi" validate:"required"`
	Channel       string `json:"channel" validate:"required"`
	ResultCode    string `json:"resultCode" validate:"required"`
}

type PersosimResponse struct {
	StatusCode        string `json:"statusCode"`
	StatusDescription string `json:"statusDescription"`
}

func (e *esb) Persosim(input *PersosimRequest) (*PersosimResponse, error) {
	result := e.GetApiInfo("persoSim")
	if result.State == "C" {
		if UserPersosim != nil {
			return UserPersosim, nil
		}
		return nil, fmt.Errorf("no custom response set for persoSim")
	}
	if result.State == "E" {
		return &PersosimResponse{
			StatusCode:        "500",
			StatusDescription: "Failed: persoSim (1)",
		}, nil
	}

	return &PersosimResponse{
		StatusCode:        "200",
		StatusDescription: "Success",
	}, nil
}
