package im

type SendSimSerialNoRequest struct {
	TransactionId string `json:"transactionId"`
	Channel       string `json:"channel"`
	OrderNo       string `json:"orderNo"`
	MobileNo      string `json:"mobileNo"`
	TrackingNo    string `json:"trackingNo"`
	SerialNo      string `json:"serialNo"`
	Imsi          string `json:"imsi"`
	ResultCode    string `json:"resultCode"`
}

type SendSimSerialNoResponse struct {
	StatusDescription string `json:"statusDescription"`
	OrderNo           string `json:"orderNo"`
	TrackingNo        string `json:"trackingNo"`
	MobileNo          string `json:"mobileNo"`
	Imsi              string `json:"imsi"`
	TransactionId     string `json:"transactionId"`
	StatusCode        string `json:"statusCode"`
	SerialNo          string `json:"serialNo"`
	HttpStatusCode    int    `json:"-"`
}

func (i *im) SendSimSerialNo(input *SendSimSerialNoRequest) (*SendSimSerialNoResponse, error) {
	res := SendSimSerialNoResponse{}
	result, err := i.app.Service.GetApiInfo("sendSimSerialNo", &res)
	if result.State == "C" {
		if err != nil {
			return nil, err
		}
		res.HttpStatusCode = result.HttpCode
		return &res, nil
	}

	if result.State == "E" {
		return &SendSimSerialNoResponse{
			StatusDescription: "error",
			OrderNo:           "",
			TrackingNo:        "",
			MobileNo:          "",
			Imsi:              "",
			TransactionId:     "",
			StatusCode:        "500",
			SerialNo:          "",
			HttpStatusCode:    500,
		}, nil
	}

	return &SendSimSerialNoResponse{
		StatusDescription: "success",
		OrderNo:           input.OrderNo,
		TrackingNo:        input.TrackingNo,
		MobileNo:          input.MobileNo,
		Imsi:              input.Imsi,
		TransactionId:     input.TransactionId,
		StatusCode:        "200",
		SerialNo:          input.SerialNo,
		HttpStatusCode:    200,
	}, nil
}
