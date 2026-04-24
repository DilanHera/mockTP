package esb

import (
	"encoding/json"
	"fmt"
)

func (e *esb) LegoupdateOrderStatus(input *json.RawMessage) (json.RawMessage, error) {
	result := e.GetApiInfo("legoUpdateOrderStatus")
	if result.State == "C" {
		if UserLegoupdateOrderStatus != nil {
			return *UserLegoupdateOrderStatus, nil
		}
		return nil, fmt.Errorf("no custom response set for legoUpdateOrderStatus")
	}
	if result.State == "E" {
		return json.RawMessage(`{"resultCode":"500","resultDesc":"Failed: legoUpdateOrderStatus (1)"}`), nil
	}

	return json.RawMessage(`{"resultCode":"200","resultDesc":"Success"}`), nil
}

