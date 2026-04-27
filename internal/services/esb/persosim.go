package esb

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
	HttpStatusCode    int    `json:"-"`
}

func (e *esb) Persosim(input *PersosimRequest) (*PersosimResponse, error) {
	res := PersosimResponse{}
	result, err := e.app.Service.GetApiInfo("persoSim", &res)
	if result.State == "C" {
		if err != nil {
			return nil, err
		}
		res.HttpStatusCode = result.HttpCode
		return &res, nil
	}
	if result.State == "E" {
		return &PersosimResponse{
			StatusCode:        "500",
			StatusDescription: "Failed: persoSim (1)",
			HttpStatusCode:    500,
		}, nil
	}

	return &PersosimResponse{
		StatusCode:        "200",
		StatusDescription: "Success",
		HttpStatusCode:    200,
	}, nil
}
