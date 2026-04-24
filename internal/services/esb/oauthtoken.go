package esb

import (
	"encoding/json"
	"fmt"
)

func (e *esb) OauthToken(input *json.RawMessage) (json.RawMessage, error) {
	result := e.GetApiInfo("oauthToken")
	if result.State == "C" {
		if UserOauthToken != nil {
			return *UserOauthToken, nil
		}
		return nil, fmt.Errorf("no custom response set for oauthToken")
	}
	if result.State == "E" {
		return json.RawMessage(`{"error":"server_error","error_description":"Failed: oauthToken (1)"}`), nil
	}

	return json.RawMessage(`{"access_token":"mock_access_token","token_type":"Bearer","expires_in":3600}`), nil
}

