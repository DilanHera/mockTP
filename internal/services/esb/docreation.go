package esb

type DOCreationRequest struct {
	PartnerName        string              `json:"PartnerName" validate:"required"`
	PartnerMessageID   string              `json:"PartnerMessageID" validate:"required"`
	ExternalDeliveryID string              `json:"ExternalDeliveryID" validate:"required"`
	ShippingPoint      string              `json:"ShippingPoint" validate:"required"`
	Partner            []DOCreationPartner `json:"Partner" validate:"required"`
	Text               []DOCreationText    `json:"Text" validate:"required"`
	Item               []DOCreationItem    `json:"Item" validate:"required"`
}

type DOCreationPartner struct {
	PartnerFunction    string `json:"PartnerFunction" validate:"required"`
	PartnerNumber      string `json:"PartnerNumber" validate:"required"`
	Name               string `json:"Name" validate:"required"`
	Name2              string `json:"Name2" validate:"required"`
	Name3              string `json:"Name3" validate:"required"`
	Name4              string `json:"Name4" validate:"required"`
	HouseNumber        string `json:"HouseNumber" validate:"required"`
	Street4            string `json:"Street4" validate:"required"`
	Street2            string `json:"Street2" validate:"required"`
	Street3            string `json:"Street3" validate:"required"`
	Street             string `json:"Street" validate:"required"`
	HomeCity           string `json:"HomeCity" validate:"required"`
	District           string `json:"District" validate:"required"`
	City               string `json:"City" validate:"required"`
	PostalCode         string `json:"PostalCode" validate:"required"`
	Region             string `json:"Region" validate:"required"`
	CountryKey         string `json:"CountryKey" validate:"required"`
	TransportationZone string `json:"TransportationZone" validate:"required"`
}

type DOCreationText struct {
	ChangeMode  string `json:"ChangeMode" validate:"required"`
	TextID      string `json:"TextID" validate:"required"`
	LanguageKey string `json:"LanguageKey" validate:"required"`
	Tagcolumn   string `json:"Tagcolumn" validate:"required"`
	TextLine    string `json:"TextLine" validate:"required"`
}

type DOCreationItem struct {
	DeliveryItem     int                `json:"DeliveryItem" validate:"required"`
	SDDocument       string             `json:"SDDocument" validate:"required"`
	SDItem           string             `json:"SDItem" validate:"required"`
	Plant            string             `json:"Plant" validate:"required"`
	QuantitySalesUom string             `json:"QuantitySalesUom" validate:"required"`
	SalesUnit        string             `json:"SalesUnit" validate:"required"`
	DeliveryDate     string             `json:"DeliveryDate" validate:"required"`
	DocumentCategory string             `json:"DocumentCategory" validate:"required"`
	OrderCombination string             `json:"OrderCombination" validate:"required"`
	Serial           []DOCreationSerial `json:"Serial" validate:"required"`
}

type DOCreationSerial struct {
	DeliveryItem string `json:"DeliveryItem" validate:"required"`
	SerialNumber string `json:"SerialNumber" validate:"required"`
}

type DOCreationResponse struct {
	MessageType      string                   `json:"MessageType" validate:"required"`
	MessageDesc      string                   `json:"MessageDesc" validate:"required"`
	MessageID        string                   `json:"MessageID" validate:"required"`
	PartnerName      string                   `json:"PartnerName" validate:"required"`
	PartnerMessageID string                   `json:"PartnerMessageID" validate:"required"`
	Item             []DOCreationResponseItem `json:"Item" validate:"required"`
	HttpStatusCode   int                      `json:"-"`
}

type DOCreationResponseItem struct {
	MessageType        string `json:"MessageType" validate:"required"`
	MessageDesc        string `json:"MessageDesc" validate:"required"`
	InterfacerecordID  string `json:"InterfacerecordID" validate:"required"`
	SDdocumentcateg    string `json:"SDdocumentcateg" validate:"required"`
	SDDocument         string `json:"SDDocument" validate:"required"`
	SDItem             string `json:"SDItem" validate:"required"`
	QuantitySalesUom   string `json:"QuantitySalesUom" validate:"required"`
	SalesUnit          string `json:"SalesUnit" validate:"required"`
	SDDocumentCategory string `json:"SDDocumentCategory" validate:"required"`
}

func (e *esb) DOCreation(input *DOCreationRequest) (*DOCreationResponse, error) {
	res := DOCreationResponse{}
	result, err := e.app.Service.GetApiInfo("doCreation", &res)
	if result.State == "C" {
		if err != nil {
			return nil, err
		}
		res.HttpStatusCode = result.HttpCode
		return &res, nil
	}
	if result.State == "E" {
		return &DOCreationResponse{
			MessageType:      "E",
			MessageDesc:      "Failed: doCreation (1)",
			MessageID:        "",
			PartnerName:      "",
			PartnerMessageID: "",
			Item: []DOCreationResponseItem{
				{
					MessageType:        "E",
					MessageDesc:        "Delivery is duplicated with DO: 2001085195",
					InterfacerecordID:  "1",
					SDdocumentcateg:    "J",
					SDDocument:         "2001085195",
					SDItem:             "10",
					QuantitySalesUom:   "1.000",
					SalesUnit:          "PC",
					SDDocumentCategory: "J",
				},
			},
			HttpStatusCode: 500,
		}, nil
	}

	return &DOCreationResponse{
		MessageType:      "S",
		MessageDesc:      "Outbound Delivery created",
		MessageID:        "A4DA6775FDAC4415BF683366C0E3401A",
		PartnerName:      "SAP",
		PartnerMessageID: "8e864f32-6beb-4e90-bb76-0355d7e31d86_1",
		Item: []DOCreationResponseItem{
			{
				MessageType:        "S",
				MessageDesc:        "Outbound Delivery created",
				InterfacerecordID:  "1",
				SDdocumentcateg:    "J",
				SDDocument:         "2001085811",
				SDItem:             "10",
				QuantitySalesUom:   "1.000",
				SalesUnit:          "PC",
				SDDocumentCategory: "J",
			},
			{
				MessageType:        "S",
				MessageDesc:        "Outbound Delivery created",
				InterfacerecordID:  "2",
				SDdocumentcateg:    "J",
				SDDocument:         "2001085811",
				SDItem:             "20",
				QuantitySalesUom:   "1.000",
				SalesUnit:          "PC",
				SDDocumentCategory: "J",
			},
		},
		HttpStatusCode: 200,
	}, nil
}
