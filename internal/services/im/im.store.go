package im

import (
	"encoding/json"
	"fmt"

	"github.com/DilanHera/mockTP/internal/app"
	"github.com/DilanHera/mockTP/internal/store"
)

var (
	apiNames = []string{"sendSimSerialNo"}

	UserSendSimSerialNo *SendSimSerialNoResponse
)

func (i *im) GetApiInfo(apiName string) store.ApiInfo {
	res, err := i.app.ApiInfoStore.Get(apiName)
	if err != nil {
		return store.ApiInfo{}
	}
	if res.Resp != "" {
		CreateResponse([]byte(res.Resp), apiName)
	}
	return *res
}

func CreateResponse(resp []byte, name string) {
	switch name {
	case "sendSimSerialNo":
		var r SendSimSerialNoResponse
		err := json.Unmarshal(resp, &r)
		if err != nil {
			break
		}
		UserSendSimSerialNo = &r
	}
}

func InitIMStore(app *app.App) {
	for _, apiName := range apiNames {
		app.ApiInfoStore.Create(apiName, "", "S", 200)
	}
}

func (i *im) SetUserSendSimSerialNo(jsonData json.RawMessage) error {
	if jsonData == nil || string(jsonData) == "" {
		UserSendSimSerialNo = nil
		return nil
	}
	response := SendSimSerialNoResponse{}
	err := json.Unmarshal(jsonData, &response)
	if err != nil {
		return fmt.Errorf("failed to unmarshal: %w", err)
	}
	UserSendSimSerialNo = &response
	return nil
}
