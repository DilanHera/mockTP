package esb

import "fmt"

type SerialNumberExpirationDateRequest struct {
}

type SerialNumberExpirationDateResponse struct {
	ResultCode string `json:"resultCode"`
	ResultDesc string `json:"resultDesc"`
}

func (e *esb) SerialNumberExpirationDate(input *SerialNumberExpirationDateRequest) (*SerialNumberExpirationDateResponse, error) {
	result := e.GetApiInfo("serialNumberExpirationDate")
	if result.State == "C" {
		if UserSerialNumberExpirationDate != nil {
			return UserSerialNumberExpirationDate, nil
		}
		return nil, fmt.Errorf("no custom response set for serialNumberExpirationDate")
	}
	if result.State == "E" {
		return &SerialNumberExpirationDateResponse{
			ResultCode: "500",
			ResultDesc: "Failed: serialNumberExpirationDate (1)",
		}, nil
	}

	return &SerialNumberExpirationDateResponse{
		ResultCode: "200",
		ResultDesc: "Success",
	}, nil
}
