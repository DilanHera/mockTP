package dt

import "fmt"

type AuthenticateRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type AuthenticateResponse struct {
	StatusCode string `json:"statusCode"`
	Message    string `json:"message"`
	Token      string `json:"token"`
}

func (d *dt) Authenticate(input *AuthenticateRequest) (*AuthenticateResponse, error) {
	result := d.GetApiInfo("authenticate")
	if result.State == "C" {
		if UserAuthenticate != nil {
			return UserAuthenticate, nil
		}
		return nil, fmt.Errorf("no custom response set for authenticate")
	}

	if result.State == "E" {
		return &AuthenticateResponse{
			StatusCode: "500",
			Message:    "authentication failed",
			Token:      "",
		}, nil
	}

	return &AuthenticateResponse{
		StatusCode: "200",
		Message:    "success authenticate",
		Token:      "eyJhbGciOiJSUzI1NiJ9.eyJpc3MiOiJkaWdpdGFsdHJhZGluZy1hdXRoIiwic3ViIjoib3B0aW11cyIsImlhdCI6MTc2OTU3NDM3NywiZXhwIjoxNzY5NjYwNzc3fQ.F_PM1Tzil-kDIvL0CcceFuVZVljofUSv0MRU8aHKXYKTorCUJvnUbuZ1R3Yi47Nfpq5okMyWmlvC2rAH-peXauKtu1CkCW9gZGSGSvPWjZJAjcwz5cVIXkpsBuI2PHobQn9AeCM8kRCGggKBFwX7ir4Rn6loBQkBuvIuJ13yBI2mJrvNAy-Hl_bfZAUpD0dGBD-izmbhYNds19I0zzDj-Qix_HA3CNlkPxf9FPDia5akqgLyKFz_sGR1Q2rbPfjymNz8vm49LUGoZ2YuDzC66_obw4wR6aHhv3VTV8uV5LEYz__ADqa2BWsi_Bb1jWYNQky09tUpNqgs_2khW4bwzw",
	}, nil
}
