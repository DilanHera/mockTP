package mychannel

type SimSerialNoRequest struct {
	TransactionId string `json:"transactionId" validate:"required"`
	OrderNo       string `json:"orderNo" validate:"required"`
	MobileNo      string `json:"mobileNo" validate:"required"`
	TrackingNo    string `json:"trackingNo" validate:"required"`
	SerialNo      string `json:"serialNo" validate:"required"`
	Imsi          string `json:"imsi" validate:"required"`
	ResultCode    string `json:"resultCode" validate:"required"`
}

type SimSerialNoResponse struct {
	ResultCode        string          `json:"resultCode" validate:"required"`
	ResultDescription string          `json:"resultDescription" validate:"required"`
	DeveloperMessage  string          `json:"developerMessage" validate:"required"`
	Data              SimSerialNoData `json:"data" validate:"omitempty"`
	HttpStatusCode    int             `json:"-"`
}

type SimSerialNoData struct {
	IsSuccess bool `json:"isSuccess" validate:"required"`
}

func (m *myChannel) SimSerialNo(input *SimSerialNoRequest) (*SimSerialNoResponse, error) {
	res := SimSerialNoResponse{}
	result, err := m.app.Service.GetApiInfo("simSerialNo", &res)
	if result.State == "C" {
		if err != nil {
			return nil, err
		}
		res.HttpStatusCode = result.HttpCode
		return &res, nil
	}

	if result.State == "E" {
		return &SimSerialNoResponse{
			ResultCode:        "50000",
			ResultDescription: "ไม่สามารถทำรายการได้ในขณะนี้",
			DeveloperMessage:  "server intenal timeout",
			HttpStatusCode:    500,
		}, nil
	}

	return &SimSerialNoResponse{
		ResultCode:        "20000",
		ResultDescription: "Success",
		DeveloperMessage:  "Success",
		Data: SimSerialNoData{
			IsSuccess: true,
		},
		HttpStatusCode: 200,
	}, nil
}
