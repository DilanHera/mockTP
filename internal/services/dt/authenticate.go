package dt

type AuthenticateRequest struct {
}

type AuthenticateResponse struct {
	StatusCode string `json:"statusCode"`
	Message    string `json:"message"`
	Token      string `json:"token"`
}
