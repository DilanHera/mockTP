package ids

type AuthenRequest struct {
}

type AuthenResponse struct {
}

func (i *ids) Authen(req *AuthenRequest) (AuthenResponse, error) {
	res := AuthenResponse{}
	return res, nil
}
