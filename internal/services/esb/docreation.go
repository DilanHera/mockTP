package esb

import "fmt"

type DOCreationRequest struct {
}

type DOCreationResponse struct {
	ResultCode string `json:"resultCode"`
	ResultDesc string `json:"resultDesc"`
}

func (e *esb) DOCreation(input *DOCreationRequest) (*DOCreationResponse, error) {
	result := e.GetApiInfo("doCreation")
	if result.State == "C" {
		if UserDOCreation != nil {
			return UserDOCreation, nil
		}
		return nil, fmt.Errorf("no custom response set for doCreation")
	}
	if result.State == "E" {
		return &DOCreationResponse{
			ResultCode: "500",
			ResultDesc: "Failed: doCreation (1)",
		}, nil
	}

	return &DOCreationResponse{
		ResultCode: "200",
		ResultDesc: "Success",
	}, nil
}
