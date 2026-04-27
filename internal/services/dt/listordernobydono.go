package dt

import (
	"math/rand"
	"strconv"
)

var ListOrderNoByDonoRequest []string

type ListOrderNoByDonoResponse struct {
	ResultCode        string        `json:"resultCode"`
	ResultDescription string        `json:"resultDescription"`
	ResultStatus      string        `json:"resultStatus"`
	ResultObj         []ListOrderNo `json:"resultObj"`
	HttpStatusCode    int           `json:"-"`
}

type ListOrderNo struct {
	ListOrderNo []string `json:"listOrderNo"`
}

func (d *dt) ListOrderNoByDono(input []string) (*ListOrderNoByDonoResponse, error) {
	res := ListOrderNoByDonoResponse{}
	result, err := d.app.Service.GetApiInfo("listOrderNoByDono", &res)
	if result.State == "C" {
		if err != nil {
			return nil, err
		}
		res.HttpStatusCode = result.HttpCode
		return &res, nil
	}

	if result.State == "E" {
		return &ListOrderNoByDonoResponse{
			ResultCode:        "50000",
			ResultDescription: "Data Not Found.",
			ResultStatus:      "F",
			ResultObj:         []ListOrderNo{},
			HttpStatusCode:    500,
		}, nil
	}

	response := &ListOrderNoByDonoResponse{
		ResultCode:        "20000",
		ResultDescription: "Success",
		ResultStatus:      "S",
		HttpStatusCode:    200,
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
