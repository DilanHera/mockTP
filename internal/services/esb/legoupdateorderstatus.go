package esb

import "fmt"

type LegoupdateOrderStatusRequest struct {
}

type LegoupdateOrderStatusResponse struct {
	ResultCode string `json:"resultCode"`
	ResultDesc string `json:"resultDesc"`
}

func (e *esb) LegoupdateOrderStatus(input *LegoupdateOrderStatusRequest) (*LegoupdateOrderStatusResponse, error) {
	result := e.GetApiInfo("legoUpdateOrderStatus")
	if result.State == "C" {
		if UserLegoupdateOrderStatus != nil {
			return UserLegoupdateOrderStatus, nil
		}
		return nil, fmt.Errorf("no custom response set for legoUpdateOrderStatus")
	}
	if result.State == "E" {
		return &LegoupdateOrderStatusResponse{
			ResultCode: "500",
			ResultDesc: "Failed: legoUpdateOrderStatus (1)",
		}, nil
	}

	return &LegoupdateOrderStatusResponse{
		ResultCode: "200",
		ResultDesc: "Success",
	}, nil
}
