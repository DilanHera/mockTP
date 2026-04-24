package esb

import (
	"encoding/json"
	"fmt"
)

func (e *esb) CreateFreightOrder(input *json.RawMessage) (json.RawMessage, error) {
	result := e.GetApiInfo("createFreightOrder")
	if result.State == "C" {
		if UserCreateFreightOrder != nil {
			return *UserCreateFreightOrder, nil
		}
		return nil, fmt.Errorf("no custom response set for createFreightOrder")
	}
	if result.State == "E" {
		return json.RawMessage(`{"resultCode":"500","resultDesc":"Failed: createFreightOrder (1)"}`), nil
	}

	return json.RawMessage(`{"resultCode":"200","resultDesc":"Success"}`), nil
}

