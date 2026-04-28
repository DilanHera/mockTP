package smis

type UpdateSerialRequest struct {
	Channel     string             `json:"channel" validate:"required"`
	TransID     string             `json:"transID" validate:"required"`
	ContractNo  string             `json:"contractNo" validate:"required"`
	SaleOrderNo string             `json:"saleOrderNo" validate:"required"`
	Company     string             `json:"company" validate:"required"`
	ActionType  string             `json:"actionType" validate:"required"`
	ItemList    []UpdateSerialItem `json:"itemList" validate:"required,dive"`
}

type UpdateSerialItem struct {
	ItemNo     int                    `json:"itemNo" validate:"required"`
	DoNo       string                 `json:"doNo" validate:"required"`
	DoDate     string                 `json:"doDate" validate:"required"`
	DoOption   string                 `json:"doOption" validate:"required"`
	DeviceInfo UpdateSerialDeviceInfo `json:"deviceInfo" validate:"required"`
	SimInfo    []UpdateSerialSimInfo  `json:"simInfo" validate:"required,dive"`
}

type UpdateSerialDeviceInfo struct {
	ImeiSerialNo string `json:"imeiSerialNo" validate:"required"`
	DeviceOption string `json:"deviceOption" validate:"required"`
}

type UpdateSerialSimInfo struct {
	MobileNo       string `json:"mobileNo" validate:"required"`
	SerialNo       string `json:"serialNo" validate:"required"`
	SimProductType string `json:"simProductType" validate:"required"`
}

type UpdateSerialResponse struct {
	ResponseCode    string `json:"ResponseCode" validate:"required"`
	ResponseMessage string `json:"ResponseMessage" validate:"required"`
	HttpStatusCode  int    `json:"-"`
}

func (s *smis) UpdateSerial(input *UpdateSerialRequest) (UpdateSerialResponse, error) {
	res := UpdateSerialResponse{}
	result, err := s.app.Service.GetApiInfo("updateSerial", &res)
	if result.State == "C" {
		if err != nil {
			return res, err
		}
		res.HttpStatusCode = result.HttpCode
		return res, nil
	}

	if result.State == "E" {
		return UpdateSerialResponse{
			ResponseCode:    "0005",
			ResponseMessage: "Data Not Found",
			HttpStatusCode:  500,
		}, nil
	}
	return UpdateSerialResponse{
		ResponseCode:    "0000",
		ResponseMessage: "Success",
		HttpStatusCode:  200,
	}, nil
}
