package esb

import "fmt"

type SerialNumberExpirationDateRequest struct {
	MessageID        string                           `json:"messageID"`
	PartnerMessageID string                           `json:"partnerMessageID"`
	PartnerName      string                           `json:"partnerName"`
	Item             []SerialNumberExpirationDateItem `json:"item"`
}

type SerialNumberExpirationDateItem struct {
	ConfirmStatus  string `json:"confirmStatus"`
	SerialNumber   string `json:"serialNumber"`
	ExpirationDate string `json:"expirationDate"`
	Material       string `json:"material"`
}

type SerialNumberExpirationDateResponse struct {
	MessageID        string                                   `json:"messageID"`
	PartnerName      string                                   `json:"partnerName"`
	PartnerMessageID string                                   `json:"partnerMessageID"`
	Item             []SerialNumberExpirationDateResponseItem `json:"item"`
}

type SerialNumberExpirationDateResponseItem struct {
	MessageType    string `json:"messageType"`
	MessageClass   string `json:"messageClass"`
	MessageNumber  string `json:"messageNumber"`
	MessageDesc    string `json:"messageDesc"`
	Material       string `json:"material"`
	SerialNumber   string `json:"serialNumber"`
	ExpirationDate string `json:"expirationDate"`
	ConfirmStatus  string `json:"confirmStatus"`
}

func (e *esb) SerialNumberExpirationDate(input *SerialNumberExpirationDateRequest) (*SerialNumberExpirationDateResponse, error) {
	result := e.GetApiInfo("serialNumberExpirationDate")
	if result.State == "C" {
		if UserSerialNumberExpirationDate != nil {
			return UserSerialNumberExpirationDate, nil
		}
		return nil, fmt.Errorf("no custom response set for serialNumberExpirationDate")
	}
	if result.State == "E" {
		return &SerialNumberExpirationDateResponse{
			MessageID:        "AGnu3gnT_T6pmJOnuzxDk0BX4EEg",
			PartnerName:      "SAP",
			PartnerMessageID: "0953A48139EC4CBCA97AF7B402305C6E",
			Item: []SerialNumberExpirationDateResponseItem{
				{
					MessageType:   "E",
					MessageClass:  "ZSCM00",
					MessageNumber: "000",
					MessageDesc:   "Equipment 37127803 has not been updated",
				},
			},
		}, nil
	}

	return &SerialNumberExpirationDateResponse{
		MessageID:        "AGnu3gnT_T6pmJOnuzxDk0BX4EEg",
		PartnerName:      "SAP",
		PartnerMessageID: "0953A48139EC4CBCA97AF7B402305C6E",
		Item: []SerialNumberExpirationDateResponseItem{
			{
				MessageType:    "S",
				MessageClass:   "ZSCM00",
				MessageNumber:  "000",
				MessageDesc:    "Equipment 37127803 has been updated",
				Material:       "1000000178",
				SerialNumber:   "2610506305004",
				ExpirationDate: "20270531",
				ConfirmStatus:  "X",
			},
		},
	}, nil
}
