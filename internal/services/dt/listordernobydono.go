package dt

import (
	"fmt"
	"math/rand"
	"strconv"
)

var ListOrderNoByDonoRequest []string

type ListOrderNoByDonoResponse struct {
	ResultCode        string        `json:"resultCode"`
	ResultDescription string        `json:"resultDescription"`
	ResultStatus      string        `json:"resultStatus"`
	ResultObj         []ListOrderNo `json:"resultObj"`
}

type ListOrderNo struct {
	ListOrderNo []string `json:"listOrderNo"`
}

func (d *dt) ListOrderNoByDono(input []string) (*ListOrderNoByDonoResponse, error) {
	result := d.GetApiInfo("listOrderNoByDono")
	if result.State == "C" {
		if UserListOrderNoByDono != nil {
			return UserListOrderNoByDono, nil
		}
		return nil, fmt.Errorf("no custom response set for listOrderNoByDono")
	}

	if result.State == "E" {
		return &ListOrderNoByDonoResponse{
			ResultCode:        "50000",
			ResultDescription: "Data Not Found.",
			ResultStatus:      "F",
			ResultObj:         []ListOrderNo{},
		}, nil
	}

	response := &ListOrderNoByDonoResponse{
		ResultCode:        "20000",
		ResultDescription: "Success",
		ResultStatus:      "S",
		ResultObj: []ListOrderNo{
			{
				ListOrderNo: []string{},
			},
		},
	}
	for range input {
		randI := rand.Intn(9000000000) + 1000000000
		response.ResultObj[0].ListOrderNo = append(response.ResultObj[0].ListOrderNo, strconv.Itoa(randI))
	}
	return response, nil
}
