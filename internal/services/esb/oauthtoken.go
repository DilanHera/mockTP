package esb

import "fmt"

type OauthTokenRequest struct {
}

type OauthTokenResponse struct {
	AccessToken      string `json:"access_token"`
	TokenType        string `json:"token_type"`
	ExpiresIn        int    `json:"expires_in"`
	Error            string `json:"error,omitempty"`
	ErrorDescription string `json:"error_description,omitempty"`
}

func (e *esb) OauthToken(input *OauthTokenRequest) (*OauthTokenResponse, error) {
	result := e.GetApiInfo("oauthToken")
	if result.State == "C" {
		if UserOauthToken != nil {
			return UserOauthToken, nil
		}
		return nil, fmt.Errorf("no custom response set for oauthToken")
	}
	if result.State == "E" {
		return &OauthTokenResponse{
			Error:            "server_error",
			ErrorDescription: "Failed: oauthToken (1)",
		}, nil
	}

	return &OauthTokenResponse{
		AccessToken: "mock_access_token",
		TokenType:   "Bearer",
		ExpiresIn:   3600,
	}, nil
}
