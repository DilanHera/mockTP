package esb

type LegoupdateOrderStatusRequest struct {
	ReferWebSessionID string                         `json:"referWebSessionID" validate:"required"`
	ReferChannel      string                         `json:"referChannel" validate:"required"`
	ReferChannelIP    string                         `json:"referChannelIP" validate:"required"`
	TransactionID     string                         `json:"transactionID" validate:"required"`
	OrderNo           string                         `json:"orderNo" validate:"required"`
	StatusOrder       string                         `json:"statusOrder" validate:"required"`
	SimCard           []LegoupdateOrderStatusSimCard `json:"simCard" validate:"required"`
}

type LegoupdateOrderStatusSimCard struct {
	SimMobileNo string `json:"simMobileNo" validate:"required"`
	SimSerial   string `json:"simSerial" validate:"required"`
}

type LegoupdateOrderStatusResponse struct {
	TransactionID string `json:"transactionID" validate:"required"`
	ResultCode    string `json:"resultCode" validate:"required"`
	ResultMessage string `json:"resultMessage" validate:"required"`
	Result        string `json:"result" validate:"required"`
	StatusOrder   string `json:"statusOrder" validate:"required"`
	OrderNo       string `json:"orderNo" validate:"required"`
	HttpStatusCode int   `json:"-"`
}

func (e *esb) LegoupdateOrderStatus(input *LegoupdateOrderStatusRequest) (*LegoupdateOrderStatusResponse, error) {
	res := LegoupdateOrderStatusResponse{}
	result, err := e.app.Service.GetApiInfo("legoUpdateOrderStatus", &res)
	if result.State == "C" {
		if err != nil {
			return nil, err
		}
		res.HttpStatusCode = result.HttpCode
		return &res, nil
	}
	if result.State == "E" {
		return &LegoupdateOrderStatusResponse{
			TransactionID: "2ad5135a-8cb6-482a-9f48-610ef68cc435",
			ResultCode:    "50000",
			ResultMessage: "Failed: legoUpdateOrderStatus (1)",
			HttpStatusCode: 500,
		}, nil
	}

	return &LegoupdateOrderStatusResponse{
		TransactionID: "2ad5135a-8cb6-482a-9f48-610ef68cc435",
		ResultCode:    "20000",
		ResultMessage: "Success.",
		Result:        "{}",
		StatusOrder:   "03",
		OrderNo:       "AS2604261179348",
		HttpStatusCode: 200,
	}, nil
}
