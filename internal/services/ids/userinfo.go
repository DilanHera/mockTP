package ids

type UserInfoRequest struct {
}

type UserInfoResponse struct {
	Sub               string `json:"sub" validate:"required"`
	Pincode           string `json:"pincode" validate:"required"`
	Firstname         string `json:"firstname" validate:"required"`
	Mobile            string `json:"mobile" validate:"required"`
	Groups            string `json:"groups" validate:"required"`
	Section           string `json:"section" validate:"required"`
	PreferredUsername string `json:"preferred_username" validate:"required"`
	Title             string `json:"title" validate:"required"`
	Consent           string `json:"consent" validate:"required"`
	Lastname          string `json:"lastname" validate:"required"`
	LocationCode      string `json:"location_code" validate:"required"`
	Name              string `json:"name" validate:"required"`
	Company           string `json:"company" validate:"required"`
	PhoneNumber       string `json:"phone_number" validate:"required"`
	Department        string `json:"department" validate:"required"`
	FamilyName        string `json:"family_name" validate:"required"`
	Email             string `json:"email" validate:"required"`
	Username          string `json:"username" validate:"required"`
	HttpStatusCode    int    `json:"-"`
}

func (i *ids) UserInfo(req *UserInfoRequest) (UserInfoResponse, error) {
	res := UserInfoResponse{}
	result, err := i.app.Service.GetApiInfo("simSerialNo", &res)
	if result.State == "C" {
		if err != nil {
			return res, err
		}
		res.HttpStatusCode = result.HttpCode
		return res, nil
	}

	if result.State == "E" {
		return UserInfoResponse{
			// ResultCode:        "50000",
			// ResultDescription: "ไม่สามารถทำรายการได้ในขณะนี้",
			// DeveloperMessage:  "server intenal timeout",
			HttpStatusCode: 500,
		}, nil
	}
	return UserInfoResponse{
		Sub:               "EMPLOYEELDAP",
		Pincode:           "00066026",
		Firstname:         "Chomnipha",
		Mobile:            "0992457344",
		Groups:            "Internal/everyone",
		Section:           "Supply Chain - E-Commerce Fulfillment",
		PreferredUsername: "Chomnipha Wetchiyo",
		Title:             "Administrative Support staff",
		Consent:           "Y",
		Lastname:          "Wetchiyo",
		LocationCode:      "Location_WDS",
		Name:              "Chomnipha Wetchiyo",
		Company:           "AWN",
		PhoneNumber:       "0992457344",
		Department:        "Supply Chain - WholeSale Fulfillment",
		FamilyName:        "Wetchiyo",
		Email:             "chomnipw@ais.co.th",
		Username:          "chomnipw",
		HttpStatusCode:    200,
	}, nil
}
