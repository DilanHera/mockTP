package esb

import "fmt"

type CreateFreightOrderRequest struct {
}

type CreateFreightOrderResponse struct {
	ResultCode string `json:"resultCode"`
	ResultDesc string `json:"resultDesc"`
}

func (e *esb) CreateFreightOrder(input *CreateFreightOrderRequest) (*CreateFreightOrderResponse, error) {
	result := e.GetApiInfo("createFreightOrder")
	if result.State == "C" {
		if UserCreateFreightOrder != nil {
			return UserCreateFreightOrder, nil
		}
		return nil, fmt.Errorf("no custom response set for createFreightOrder")
	}
	if result.State == "E" {
		return &CreateFreightOrderResponse{
			ResultCode: "500",
			ResultDesc: "Failed: createFreightOrder (1)",
		}, nil
	}

	return &CreateFreightOrderResponse{
		ResultCode: "200",
		ResultDesc: "Success",
	}, nil
}
