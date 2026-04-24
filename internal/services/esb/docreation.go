package esb

import (
	"encoding/json"
	"fmt"
)

func (e *esb) DOCreation(input *json.RawMessage) (json.RawMessage, error) {
	result := e.GetApiInfo("doCreation")
	if result.State == "C" {
		if UserDOCreation != nil {
			return *UserDOCreation, nil
		}
		return nil, fmt.Errorf("no custom response set for doCreation")
	}
	if result.State == "E" {
		return json.RawMessage(`{"resultCode":"500","resultDesc":"Failed: doCreation (1)"}`), nil
	}

	return json.RawMessage(`{"resultCode":"200","resultDesc":"Success"}`), nil
}

