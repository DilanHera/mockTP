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
}

func (d *dt) PickingDocument(input *PickingDocumentRequest) (*PickingDocumentResponse, error) {
	if UserPickingDocument != nil {
		return UserPickingDocument, nil
	}
	response := &PickingDocumentResponse{
		ResultCode: "20000",
		ResultDesc: "Success",
		Status:     "S",
	}
	return response, nil
}
