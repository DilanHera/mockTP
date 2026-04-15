package dt

type UpdateSimSerialPersoRequest struct {
	OrderNo  string `json:"orderNo"`
	MobileNo string `json:"mobileNo"`
	SerialNo string `json:"serialNo"`
}

type UpdateSimSerialPersoResponse struct {
	ResultCode        string `json:"resultCode"`
	ResultDescription string `json:"resultDescription"`
	Status            string `json:"status"`
}

func (d *dt) UpdateSimSerialPerso(input *UpdateSimSerialPersoRequest) (*UpdateSimSerialPersoResponse, error) {
	if UserUpdateSimSerialPerso != nil {
		return UserUpdateSimSerialPerso, nil
	}
	return &UpdateSimSerialPersoResponse{
		ResultCode:        "20000",
		ResultDescription: "Success",
		Status:            "S",
	}, nil
}
