package esb

import "fmt"

type OauthTokenRequest struct {
	ClientId     string `json:"client_id" validate:"required"`
	ClientSecret string `json:"client_secret" validate:"required"`
	GrantType    string `json:"grant_type" validate:"required"`
	Nonce        string `json:"nonce"`
}

type OauthTokenResponse struct {
	AccessToken string `json:"access_token,omitempty"`
	TokenType   string `json:"token_type,omitempty"`
	ExpiresIn   int    `json:"expires_in,omitempty"`
	Error       string `json:"error,omitempty"`
	StatusCode  int    `json:"-"`
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
			Error:      "invalid_client",
			StatusCode: 401,
		}, nil
	}

	return &OauthTokenResponse{
		AccessToken: "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCIsImtpZCI6IlIyNHNTYzNGNzQifQ.eyJpc3MiOiJzcmYuYWlzLmNvLnRoL3NjZiIsInN1YiI6InRva2VuX2NsaWVudF9jcmVkZW50aWFscyIsImF1ZCI6ImREMlhmNElzeXBWaTg2Y2FsL1ZSSmFENklyTERnN0loSmdqdUh5WlQ3ckhYQmRkLzRXSHlhdz09IiwiZXhwIjoxNTg3NzA3NTIwLCJpYXQiOjE1ODc3MDM5MjAsImp0aSI6ImpLUUg3QU1kOE9OYjNDeTBodld0cUMiLCJjbGllbnQiOiJNekF4TWpJc1ZHVnpkRk5wYm1kMWJHUmZpWFI1ZkVKaFkydGxibVI4TVM0d0xqQT0iLCJzc2lkIjoiSmw4MTVTWkFYM2FWRVRIMFN6enF0YSJ9.ljsVAvW9eCc8L_G3kaK9vB1TnkoB1A8nLZtZvTCxg9w66P_DhnScrbf2_6a7MYEwp5sRKPHRiPQvDrbjtGtmzWUhCZI_b2Z77zJ--jIUYzZmD1cTRDEXKItXqSxKd4aFp761BUMkOxw2KX_sWQaS4Z9OPy68p5XNIx9S0p9Mjc",
		TokenType:   "bearer",
		ExpiresIn:   3600,
		StatusCode:  200,
	}, nil
}
