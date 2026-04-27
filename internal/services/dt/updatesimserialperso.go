package dt

type UpdateSimSerialPersoRequest struct {
	OrderNo  string `json:"orderNo" validate:"required"`
	MobileNo string `json:"mobileNo" validate:"required"`
	SerialNo string `json:"serialNo" validate:"required"`
}

type UpdateSimSerialPersoResponse struct {
	ResultCode        string `json:"resultCode"`
	ResultDescription string `json:"resultDescription"`
	Status            string `json:"status"`
	HttpStatusCode    int    `json:"-"`
}

func (d *dt) UpdateSimSerialPerso(input *UpdateSimSerialPersoRequest) (*UpdateSimSerialPersoResponse, error) {
	res := UpdateSimSerialPersoResponse{}
	result, err := d.app.Service.GetApiInfo("updateSimSerialPerso", &res)
	if result.State == "C" {
		if err != nil {
			return nil, err
		}
		res.HttpStatusCode = result.HttpCode
		return &res, nil
	}

	if result.State == "E" {
		return &UpdateSimSerialPersoResponse{
			ResultCode:        "50000",
			ResultDescription: "Data Not Found",
			Status:            "F",
			HttpStatusCode:    500,
		}, nil
	}

	return &UpdateSimSerialPersoResponse{
		ResultCode:        "20000",
		ResultDescription: "Success",
		Status:            "S",
		HttpStatusCode:    200,
	}, nil
}
