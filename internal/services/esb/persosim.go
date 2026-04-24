package esb

import (
	"encoding/json"
	"fmt"
)

func (e *esb) Persosim(input *json.RawMessage) (json.RawMessage, error) {
	result := e.GetApiInfo("persoSim")
	if result.State == "C" {
		if UserPersosim != nil {
			return *UserPersosim, nil
		}
		return nil, fmt.Errorf("no custom response set for persoSim")
	}
	if result.State == "E" {
		return json.RawMessage(`{"resultCode":"500","resultDesc":"Failed: persoSim (1)"}`), nil
	}

	return json.RawMessage(`{"resultCode":"200","resultDesc":"Success"}`), nil
}

