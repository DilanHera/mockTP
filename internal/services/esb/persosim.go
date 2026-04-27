package esb

import "fmt"

type PersosimRequest struct {
}

type PersosimResponse struct {
	ResultCode string `json:"resultCode"`
	ResultDesc string `json:"resultDesc"`
}

func (e *esb) Persosim(input *PersosimRequest) (*PersosimResponse, error) {
	result := e.GetApiInfo("persoSim")
	if result.State == "C" {
		if UserPersosim != nil {
			return UserPersosim, nil
		}
		return nil, fmt.Errorf("no custom response set for persoSim")
	}
	if result.State == "E" {
		return &PersosimResponse{
			ResultCode: "500",
			ResultDesc: "Failed: persoSim (1)",
		}, nil
	}

	return &PersosimResponse{
		ResultCode: "200",
		ResultDesc: "Success",
	}, nil
}
