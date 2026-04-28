package eos

type UpdateSimSerialNoRequest struct {
	TransactionId string `json:"transactionId" validate:"required"`
	MobileNo      string `json:"mobileNo" validate:"required"`
	TrackingNo    string `json:"trackingNo" validate:"required"`
	SerialNo      string `json:"serialNo" validate:"required"`
	Imsi          string `json:"imsi" validate:"required"`
	OrderNo       string `json:"orderNo" validate:"required"`
	Channel       string `json:"channel" validate:"required"`
	ResultCode    string `json:"resultCode" validate:"required"`
}

type UpdateSimSerialNoResponse struct {
	StatusCode        string `json:"statusCode" validate:"required"`
	StatusDescription string `json:"statusDescription" validate:"required"`
	HttpStatusCode    int    `json:"-"`
}

func (e *eos) UpdateSimSerialNo(input *UpdateSimSerialNoRequest) (UpdateSimSerialNoResponse, error) {
	res := UpdateSimSerialNoResponse{}
	result, err := e.app.Service.GetApiInfo("updateSimSerialNo", &res)
	if result.State == "C" {
		if err != nil {
			return res, err
		}
		res.HttpStatusCode = result.HttpCode
		return res, nil
	}

	if result.State == "E" {
		return UpdateSimSerialNoResponse{
			StatusCode:        "500",
			StatusDescription: "Fail",
			HttpStatusCode:    500,
		}, nil
	}
	return UpdateSimSerialNoResponse{
		StatusCode:        "200",
		StatusDescription: "Success",
		HttpStatusCode:    200,
	}, nil
}
