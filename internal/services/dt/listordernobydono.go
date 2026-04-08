package dt

type ListOrderNoByDonoRequest struct {
	RequestBody []string `json:"requestBody"`
}

type ListOrderNoByDonoResponse struct {
	ResultObj []ListOrderNo `json:"resultObj"`
}

type ListOrderNo struct {
	ListOrderNo []string `json:"listOrderNo"`
}

func (d *dt) ListOrderNoByDono(input *ListOrderNoByDonoRequest) (*ListOrderNoByDonoResponse, error) {
	if UserListOrderNoByDono != nil {
		return UserListOrderNoByDono, nil
	}
	response := &ListOrderNoByDonoResponse{
		ResultObj: []ListOrderNo{
			{
				ListOrderNo: []string{"1234567890"},
			},
		},
	}
	return response, nil
}
