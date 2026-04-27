package dt

type PickingDocumentRequest struct {
	DocNo          string        `json:"docNo" validate:"required"`
	DocType        string        `json:"docType" validate:"required"`
	Company        string        `json:"company" validate:"required"`
	UserId         string        `json:"userId" validate:"required"`
	UserLocCode    string        `json:"userLocCode" validate:"required"`
	UserMacAddress string        `json:"userMacAddress" validate:"required"`
	EmpCode        string        `json:"empCode" validate:"omitempty"`
	PickingItems   []PickingItem `json:"pickingItems" validate:"required,dive"`
}

type PickingItem struct {
	InvSeq          string `json:"invSeq" validate:"required"`
	MatCode         string `json:"matCode" validate:"required"`
	SerialNo        string `json:"serialNo" validate:"omitempty"`
	PickFlg         string `json:"pickFlg" validate:"required"`
	Qty             string `json:"qty" validate:"required"`
	BypassOptiusFlg string `json:"bypassOptiusFlg" validate:"omitempty"`
	SimSerialNo     string `json:"simSerialNo" validate:"omitempty"`
}

type PickingDocumentResponse struct {
	ResultCode string `json:"resultCode"`
	ResultDesc string `json:"resultDescription"`
	Status     string `json:"status"`
	HttpStatusCode int `json:"-"`
}

func (d *dt) PickingDocument(input *PickingDocumentRequest) (*PickingDocumentResponse, error) {
	res := PickingDocumentResponse{}
	result, err := d.app.Service.GetApiInfo("pickingDocument", &res)
	if result.State == "C" {
		if err != nil {
			return nil, err
		}
		res.HttpStatusCode = result.HttpCode
		return &res, nil
	}

	if result.State == "E" {
		return &PickingDocumentResponse{
			ResultCode: "50000",
			ResultDesc: "Not found picking items to picking document",
			Status:     "F",
			HttpStatusCode: 500,
		}, nil
	}

	response := &PickingDocumentResponse{
		ResultCode: "20000",
		ResultDesc: "Success picking document list",
		Status:     "S",
		HttpStatusCode: 200,
	}
	return response, nil
}
