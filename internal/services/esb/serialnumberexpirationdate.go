package esb

import (
	"encoding/json"
	"fmt"
)

func (e *esb) SerialNumberExpirationDate(input *json.RawMessage) (json.RawMessage, error) {
	result := e.GetApiInfo("serialNumberExpirationDate")
	if result.State == "C" {
		if UserSerialNumberExpirationDate != nil {
			return *UserSerialNumberExpirationDate, nil
		}
		return nil, fmt.Errorf("no custom response set for serialNumberExpirationDate")
	}
	if result.State == "E" {
		return json.RawMessage(`{"resultCode":"500","resultDesc":"Failed: serialNumberExpirationDate (1)"}`), nil
	}

	return json.RawMessage(`{"resultCode":"200","resultDesc":"Success"}`), nil
}

