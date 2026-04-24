package dt

import "fmt"

type UpdateSimSerialPersoRequest struct {
	OrderNo  string `json:"orderNo" validate:"required"`
	MobileNo string `json:"mobileNo" validate:"required"`
	SerialNo string `json:"serialNo" validate:"required"`
}

type UpdateSimSerialPersoResponse struct {
	ResultCode        string `json:"resultCode"`
	ResultDescription string `json:"resultDescription"`
	Status            string `json:"status"`
}

func (d *dt) UpdateSimSerialPerso(input *UpdateSimSerialPersoRequest) (*UpdateSimSerialPersoResponse, error) {
	result := d.GetApiInfo("updateSimSerialPerso")
	if result.State == "C" {
		if UserUpdateSimSerialPerso != nil {
			return UserUpdateSimSerialPerso, nil
		}
		return nil, fmt.Errorf("no custom response set for updateSimSerialPerso")
	}

	if result.State == "E" {
		return &UpdateSimSerialPersoResponse{
			ResultCode:        "50000",
			ResultDescription: "Data Not Found",
			Status:            "F",
		}, nil
	}

	return &UpdateSimSerialPersoResponse{
		ResultCode:        "20000",
		ResultDescription: "Success",
		Status:            "S",
	}, nil
}
