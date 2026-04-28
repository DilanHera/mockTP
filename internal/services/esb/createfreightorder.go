package esb

import (
	"context"
	"encoding/json"
	"fmt"
	"regexp"
	"time"

	"github.com/DilanHera/mockTP/internal/kafka"
)

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
	HttpStatusCode     int    `json:"-"`
}

// KafkaMessage is the top-level envelope published to Kafka after a freight order is created.
type KafkaMessage struct {
	Header KafkaHeader      `json:"header"`
	Body   KafkaFreightBody `json:"body"`
}

type KafkaHeader struct {
	Version        string        `json:"version"`
	Timestamp      string        `json:"timestamp"`
	OrgService     string        `json:"orgService"`
	Scope          string        `json:"scope"`
	From           string        `json:"from"`
	Channel        string        `json:"channel"`
	Agent          string        `json:"agent"`
	Broker         string        `json:"broker"`
	UseCase        string        `json:"useCase"`
	UseCaseStep    string        `json:"useCaseStep"`
	UseCaseAge     int           `json:"useCaseAge"`
	FunctionName   string        `json:"functionName"`
	MessageType    string        `json:"messageType"`
	Session        string        `json:"session"`
	Transaction    string        `json:"transaction"`
	Communication  string        `json:"communication"`
	GroupTags      []string      `json:"groupTags"`
	Identity       KafkaIdentity `json:"identity"`
	ReturnedError  string        `json:"returnedError"`
	InitUri        string        `json:"initUri"`
	InitMethod     string        `json:"initMethod"`
	TmfSpec        string        `json:"tmfSpec"`
	BaseApiVersion string        `json:"baseApiVersion"`
	SchemaVersion  string        `json:"schemaVersion"`
}

type KafkaIdentity struct {
	Public []string `json:"public"`
}

type KafkaFreightBody struct {
	HeaderText            string             `json:"HeaderText"`
	FreightOrderNumber    string             `json:"FreightOrderNumber"`
	CarrierTrackingNumber string             `json:"CarrierTrackingNumber"`
	CarrierName           string             `json:"CarrierName"`
	Carrier               string             `json:"Carrier"`
	EventCode             string             `json:"EventCode"`
	EventDescription      string             `json:"EventDescription"`
	ActualDate            string             `json:"ActualDate"`
	Item                  []KafkaFreightItem `json:"Item"`
	MessageID             string             `json:"MessageID"`
	PartnerName           string             `json:"PartnerName"`
	PartnerMessageID      string             `json:"PartnerMessageID"`
}

type KafkaFreightItem struct {
	ItemNo                    string `json:"ItemNo"`
	SourceSystemName          string `json:"SourceSystemName"`
	SourceSystem              string `json:"SourceSystem"`
	CustomerReference         string `json:"CustomerReference"`
	SalesOrderNumber          string `json:"SalesOrderNumber"`
	StockTransportOrderNumber string `json:"StockTransportOrderNumber"`
}

func (e *esb) callProducer(freightOrderNo, referenceNo string) error {
	messageId := "AGkAgK7iwQ0bBhmhtEU_1YcE9T5P"
	re := regexp.MustCompile(`^[a-zA-Z]*`)
	trackingNo := re.ReplaceAllString(referenceNo, "MOCK")
	msg := KafkaMessage{
		Header: KafkaHeader{
			Version:        "5.0",
			Timestamp:      time.Now().UTC().Format(time.RFC3339Nano),
			OrgService:     "SAP",
			Scope:          "global",
			From:           "SAP",
			FunctionName:   "FreightOrder",
			MessageType:    "event",
			Session:        messageId,
			Transaction:    messageId,
			Communication:  "unicast",
			GroupTags:      []string{},
			Identity:       KafkaIdentity{Public: []string{messageId}},
			TmfSpec:        "none",
			BaseApiVersion: "none",
			SchemaVersion:  "none",
		},
		Body: KafkaFreightBody{
			FreightOrderNumber:    freightOrderNo,
			CarrierTrackingNumber: trackingNo,
			CarrierName:           "บจ. โกลบอล เจท เอ็กซ์เพรส",
			Carrier:               "6100108050",
			EventCode:             "Z_PICKUP",
			ActualDate:            time.Now().UTC().Format("20060102150405"),
			Item: []KafkaFreightItem{
				{
					ItemNo:            "1",
					SourceSystemName:  "Redeem (Privilege)",
					SourceSystem:      "Z008",
					CustomerReference: referenceNo,
				},
			},
			MessageID:        messageId,
			PartnerName:      "SAP",
			PartnerMessageID: "13cbfb26-3283-1fd0-acfb-28b42d24b189",
		},
	}

	data, err := json.Marshal(msg)
	if err != nil {
		fmt.Println("failed to marshal message:", err)
		return err
	}

	return e.app.Kafka.Produce(kafka.KafkaProducerConfig{
		Context:  context.Background(),
		Topic:    "sap.proxy.trackingStatusUpdated",
		Messages: []string{string(data)},
	})
}

func (e *esb) CreateFreightOrder(input *CreateFreightOrderRequest) (*CreateFreightOrderResponse, error) {
	res := CreateFreightOrderResponse{}
	result, err := e.app.Service.GetApiInfo("createFreightOrder", &res)
	if result.State == "C" {
		if err != nil {
			return nil, err
		}
		_ = e.callProducer(res.FreightOrderNumber, input.ReferenceNumber)
		res.HttpStatusCode = result.HttpCode
		return &res, nil
	}
	if result.State == "E" {
		return &CreateFreightOrderResponse{
			FreightOrderNumber: "",
			MessageType:        "E",
			MessageDesc:        "Reference Number already exist in Freight Order 6200087606",
			MessageID:          "58969D1B6FAF410DB04DCC65D0689918",
			PartnerName:        "OPTIMUS",
			PartnerMessageID:   "2026042115083192543",
			HttpStatusCode:     500,
		}, nil
	}

	resp := &CreateFreightOrderResponse{
		FreightOrderNumber: "6200088900",
		MessageType:        "S",
		MessageDesc:        "Business document with temporary number $1 saved as business doc. 6200088900",
		MessageID:          "AGkAgK7iwQ0bBhmhtEU_1YcE9T5P",
		PartnerName:        "OPTIMUS",
		PartnerMessageID:   "13cbfb26-3283-1fd0-acfb-28b42d24b189",
		HttpStatusCode:     200,
	}

	_ = e.callProducer(resp.FreightOrderNumber, resp.FreightOrderNumber)

	return resp, nil
}
