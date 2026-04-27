package esb

import "fmt"

type CreateFreightOrderRequest struct {
	MessageID            string                   `json:"MessageID" validate:"required"`
	PartnerName          string                   `json:"PartnerName" validate:"required"`
	PartnerMessageID     string                   `json:"PartnerMessageID" validate:"required"`
	DocumentCategory     string                   `json:"DocumentCategory" validate:"required"`
	FreightOrderType     string                   `json:"FreightOrderType" validate:"required"`
	PurchasingOrgID      string                   `json:"PurchasingOrgID" validate:"required"`
	PlanExecOrg          string                   `json:"PlanExecOrg" validate:"required"`
	SourceStopIdentifier string                   `json:"SourceStopIdentifier" validate:"required"`
	SourceLocation       string                   `json:"SourceLocation" validate:"required"`
	DestStopIdentifier   string                   `json:"DestStopIdentifier" validate:"required"`
	DestName1            string                   `json:"DestName1" validate:"required"`
	DestName2            string                   `json:"DestName2" validate:"required"`
	DestName3            string                   `json:"DestName3" validate:"required"`
	DestName4            string                   `json:"DestName4" validate:"required"`
	DestStreet4          string                   `json:"DestStreet4" validate:"required"`
	DestStreet2          string                   `json:"DestStreet2" validate:"required"`
	DestStreet3          string                   `json:"DestStreet3" validate:"required"`
	DestStreet           string                   `json:"DestStreet" validate:"required"`
	DestOtherCity        string                   `json:"DestOtherCity" validate:"required"`
	DestDistrict         string                   `json:"DestDistrict" validate:"required"`
	DestCity             string                   `json:"DestCity" validate:"required"`
	DestPostalCode       string                   `json:"DestPostalCode" validate:"required"`
	DestRegion           string                   `json:"DestRegion" validate:"required"`
	DestCountry          string                   `json:"DestCountry" validate:"required"`
	DestTelephone        string                   `json:"DestTelephone" validate:"required"`
	SourceSystem         string                   `json:"SourceSystem" validate:"required"`
	ReferenceNumber      string                   `json:"ReferenceNumber" validate:"required"`
	Items                []CreateFreightOrderItem `json:"Items" validate:"required"`
}

type CreateFreightOrderItem struct {
	ItemID        string `json:"ItemID" validate:"required"`
	ItemType      string `json:"ItemType" validate:"required"`
	Quantity      string `json:"Quantity" validate:"required"`
	UnitofMeasure string `json:"UnitofMeasure" validate:"required"`
}

type CreateFreightOrderResponse struct {
	FreightOrderNumber string `json:"FreightOrderNumber" validate:"required"`
	MessageType        string `json:"MessageType" validate:"required"`
	MessageDesc        string `json:"MessageDesc" validate:"required"`
	MessageID          string `json:"MessageID" validate:"required"`
	PartnerName        string `json:"PartnerName" validate:"required"`
	PartnerMessageID   string `json:"PartnerMessageID" validate:"required"`
}

func (e *esb) CreateFreightOrder(input *CreateFreightOrderRequest) (*CreateFreightOrderResponse, error) {
	result := e.GetApiInfo("createFreightOrder")
	if result.State == "C" {
		if UserCreateFreightOrder != nil {
			return UserCreateFreightOrder, nil
		}
		return nil, fmt.Errorf("no custom response set for createFreightOrder")
	}
	if result.State == "E" {
		return &CreateFreightOrderResponse{
			FreightOrderNumber: "",
			MessageType:        "E",
			MessageDesc:        "Reference Number already exist in Freight Order 6200087606",
			MessageID:          "58969D1B6FAF410DB04DCC65D0689918",
			PartnerName:        "OPTIMUS",
			PartnerMessageID:   "2026042115083192543",
		}, nil
	}

	return &CreateFreightOrderResponse{
		FreightOrderNumber: "6200088900",
		MessageType:        "S",
		MessageDesc:        "Business document with temporary number $1 saved as business doc. 6200088900",
		MessageID:          "ECF92F6BF2EA4976B423247278CEB458",
		PartnerName:        "OPTIMUS",
		PartnerMessageID:   "20260427080253557",
	}, nil
}
