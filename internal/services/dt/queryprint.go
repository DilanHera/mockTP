package dt

import "fmt"

type QueryPrintRequest struct {
	ReceiptNum string `json:"receiptNum" validate:"required"`
	Company    string `json:"company" validate:"required"`
	UserId     string `json:"userId" validate:"required"`
	ReprintFlg string `json:"reprintFlg" validate:"required"`
	ReportType string `json:"reportType" validate:"required"`
}

type QueryPrintResponse struct {
	ResultCode        string     `json:"resultCode"`
	ResultDescription string     `json:"resultDescription"`
	DeveloperMessage  string     `json:"developerMessage"`
	DataList          []DataList `json:"dataList"`
}

type DataList struct {
	ReceiptCompany      string `json:"RECEIPT_COMPANY"`
	CompanyName         string `json:"COMPANY_NAME"`
	CompanyAddr1        string `json:"COMPANY_ADDR1"`
	CompanyAddr2        string `json:"COMPANY_ADDR2"`
	CompanyProvince     string `json:"COMPANY_PROVINCE"`
	CompanyPostcode     string `json:"COMPANY_POSTCODE"`
	CompanyPhoneNo      string `json:"COMPANY_PHONE_NO"`
	CompanyTaxNo        string `json:"COMPANY_TAX_NO"`
	LocationCode        string `json:"LOCATION_CODE"`
	LocationName        string `json:"LOCATION_NAME"`
	LocationAddress1    string `json:"LOCATION_ADDRESS1"`
	LocationAddress2    string `json:"LOCATION_ADDRESS2"`
	Brand               string `json:"BRAND"`
	Qty                 string `json:"QTY"`
	PriceAmt            string `json:"PRICE_AMT"`
	CustomerName        string `json:"CUSTOMER_NAME"`
	CustomerAddr1       string `json:"CUSTOMER_ADDR1"`
	CustomerAddr2       string `json:"CUSTOMER_ADDR2"`
	MobilePhone         string `json:"MOBILE_PHONE"`
	MobileName          string `json:"MOBILE_NAME"`
	MobileAddr1         string `json:"MOBILE_ADDR1"`
	MobileAddr2         string `json:"MOBILE_ADDR2"`
	ShipCustName        string `json:"SHIP_CUST_NAME"`
	ShipCustAddr1       string `json:"SHIP_CUST_ADDR1"`
	ShipCustAddr2       string `json:"SHIP_CUST_ADDR2"`
	TradeName           string `json:"TRADE_NAME"`
	BrandNull           string `json:"BRAND_NULL"`
	PriceExcAmt         string `json:"PRICE_EXC_AMT"`
	ExcDiscAmt          string `json:"EXC_DISC_AMT"`
	Amount              string `json:"AMOUNT"`
	ReprintRemark       string `json:"REPRINT_REMARK"`
	ReprintSeq          string `json:"REPRINT_SEQ"`
	VatRate             string `json:"VAT_RATE"`
	VatAmt              string `json:"VAT_AMT"`
	RefNo               string `json:"REF_NO"`
	ReceiptId           string `json:"RECEIPT_ID"`
	ReceiptDt           string `json:"RECEIPT_DT"`
	ReceiptNum          string `json:"RECEIPT_NUM"`
	SpDiscIncAmt        string `json:"SP_DISC_INC_AMT"`
	CaDiscIncAmt        string `json:"CA_DISC_INC_AMT"`
	SaleDiscExcAmt      string `json:"SALE_DISC_EXC_AMT"`
	HeadStatus          string `json:"HEAD_STATUS"`
	DepositFlg          string `json:"DEPOSIT_FLG"`
	UserId              string `json:"USER_ID"`
	ReprintFlg          string `json:"REPRINT_FLG"`
	CompanyTax          string `json:"COMPANY_TAX"`
	InvId               string `json:"INV_ID"`
	RefMainSeq          string `json:"REF_MAIN_SEQ"`
	InvSeq              string `json:"INV_SEQ"`
	VatFlg              string `json:"VAT_FLG"`
	MobileNo            string `json:"MOBILE_NO"`
	WtAmt               string `json:"WT_AMT"`
	PaidExWt            string `json:"PAID_EX_WT"`
	ReceiptType         string `json:"RECEIPT_TYPE"`
	ShipToFlg           string `json:"SHIP_TO_FLG"`
	ShipLocationName    string `json:"SHIP_LOCATION_NAME"`
	CreateTransaction   string `json:"CREATE_TRANSACTION"`
	PayChannel          string `json:"PAY_CHANNEL"`
	MatCode             string `json:"MAT_CODE"`
	TaxFlag             string `json:"TAX_FLAG"`
	CitizenId           string `json:"CITIZEN_ID"`
	FocFlg              string `json:"FOC_FLG"`
	FocVat              string `json:"FOC_VAT"`
	FocAmt              string `json:"FOC_AMT"`
	FocSpDiscIncAmt     string `json:"FOC_SP_DISC_INC_AMT"`
	AddressThaiColumn   string `json:"ADDRESS_THAI_COLUMN"`
	FocDesc             string `json:"FOC_DESC"`
	SaleType            string `json:"SALE_TYPE"`
	FreeGoodsFlg        string `json:"FREE_GOODS_FLG"`
	PayType1            string `json:"PAY_TYPE_1"`
	BankName1           string `json:"BANK_NAME_1"`
	Installment1        string `json:"INSTALLMENT_1"`
	PaymentAmount1      string `json:"PAYMENT_AMOUNT_1"`
	PaidBy1             string `json:"PAID_BY_1"`
	BankName2           string `json:"BANK_NAME_2"`
	PaymentAmount2      string `json:"PAYMENT_AMOUNT_2"`
	BankName3           string `json:"BANK_NAME_3"`
	PaymentAmount3      string `json:"PAYMENT_AMOUNT_3"`
	BankName4           string `json:"BANK_NAME_4"`
	PaymentAmount4      string `json:"PAYMENT_AMOUNT_4"`
	BankName5           string `json:"BANK_NAME_5"`
	PaymentAmount5      string `json:"PAYMENT_AMOUNT_5"`
	BankName6           string `json:"BANK_NAME_6"`
	PaymentAmount6      string `json:"PAYMENT_AMOUNT_6"`
	BankName7           string `json:"BANK_NAME_7"`
	PaymentAmount7      string `json:"PAYMENT_AMOUNT_7"`
	TaxInfoSellerBranch string `json:"TAX_INFO_SELLER_BRANCH"`
	TaxInfoSellerAddr   string `json:"TAX_INFO_SELLER_ADDR"`
	TaxInfoBuyerBranch  string `json:"TAX_INFO_BUYER_BRANCH"`
	TaxInfoBuyerTaxNo   string `json:"TAX_INFO_BUYER_TAX_NO"`
	TaxInfoBuyerPhoneNo string `json:"TAX_INFO_BUYER_PHONE_NO"`
	PosNo               string `json:"POS_NO"`
	ACommerceFlg        string `json:"A_COMMERCE_FLG"`
	ServiceFlg          string `json:"SERVICE_FLG"`
	TradeIn             string `json:"TRADE_IN"`
	DescriptTi1         string `json:"DESCRIPT_TI1"`
	DescriptTi2         string `json:"DESCRIPT_TI2"`
	AcomGenDeliveryFlg  string `json:"ACOM_GEN_DELIVERY_FLG"`
	ThaiAmount          string `json:"THAI_AMOUNT"`
	FgRdStatus          string `json:"FG_RD_STATUS"`
	ShipLocationCode    string `json:"SHIP_LOCATION_CODE"`
	Total               string `json:"TOTAL"`
	GrandTotal          string `json:"GRAND_TOTAL"`
	EReceiptFlg         string `json:"E_RECEIPT_FLG"`
	CnRmaAmount         string `json:"CN_RMA_AMOUNT"`
	SaleOrderNo         string `json:"SALE_ORDER_NO"`
}

func (d *dt) QueryPrint(input *QueryPrintRequest) (*QueryPrintResponse, error) {
	result := d.GetApiInfo("queryPrint")
	if result.State == "C" {
		if UserQueryPrint != nil {
			return UserQueryPrint, nil
		}
		return nil, fmt.Errorf("no custom response set for queryPrint")
	}

	if result.State == "E" {
		return &QueryPrintResponse{
			ResultCode:        "50000",
			ResultDescription: "Data incorrect",
			DeveloperMessage:  "Data incorrect",
			DataList:          []DataList{},
		}, nil
	}

	response := &QueryPrintResponse{
		ResultCode:        "20000",
		ResultDescription: "Success",
		DeveloperMessage:  "Success",
		DataList: []DataList{
			{
				ReceiptCompany:      "AWN",
				CompanyName:         "บรษท แอดวานซ ไวรเลส เนทเวอรค จำกด",
				CompanyAddr1:        "414 ถนนพหลโยธน แขวงสามเสนใน",
				CompanyAddr2:        "เขตพญาไท",
				CompanyProvince:     "กรงเทพฯ",
				CompanyPostcode:     "10400",
				CompanyPhoneNo:      "02-6874014 ",
				CompanyTaxNo:        "0105548115897",
				LocationCode:        "1117",
				LocationName:        "สาขา เซนทรล พระราม 2",
				LocationAddress1:    "160  ถนนพระรามท2",
				LocationAddress2:    "แขวงแสมดำ เขตบางขนเทยน กรงเทพฯ 10150",
				Brand:               "AIS Group",
				Qty:                 "1",
				PriceAmt:            "179.00",
				CustomerName:        "นาย ทดสอบ จองจายเตม",
				CustomerAddr1:       "15 หม13 ถนน-",
				CustomerAddr2:       " ตำบลเชงเนน อำเภอเมองระยอง จงหวดระยอง 21000",
				MobilePhone:         "0987184155",
				MobileName:          "นาย ทดสอบ จองจายเตม",
				MobileAddr1:         "15 หม13 ถนน-",
				MobileAddr2:         " ตำบลเชงเนน อำเภอเมองระยอง จงหวดระยอง 21000",
				ShipCustName:        "ทดสอบ จองจายเตม",
				ShipCustAddr1:       "15 หม13 ถนน-",
				ShipCustAddr2:       " ตำบลเชงเนน อำเภอเมองระยอง จงหวดระยอง 21000",
				TradeName:           "null",
				BrandNull:           "คาขนสง online",
				PriceExcAmt:         "167.29",
				ExcDiscAmt:          "0.00",
				Amount:              "167.29",
				ReprintRemark:       "AO202601283249",
				ReprintSeq:          "0",
				VatRate:             "7",
				VatAmt:              "2824.79",
				RefNo:               "null",
				ReceiptId:           "5141209",
				ReceiptDt:           "2026-01-28 00:00:00",
				ReceiptNum:          "1117260100000006",
				SpDiscIncAmt:        "0.00",
				CaDiscIncAmt:        "0.00",
				SaleDiscExcAmt:      "0.00",
				HeadStatus:          "Y",
				DepositFlg:          "N",
				UserId:              "naroc669",
				ReprintFlg:          "N",
				CompanyTax:          "0105548115897",
				InvId:               "5141240",
				RefMainSeq:          "0",
				InvSeq:              "2",
				VatFlg:              "Y",
				MobileNo:            "0699933593",
				WtAmt:               "0",
				PaidExWt:            "43179.00",
				ReceiptType:         "FULL",
				ShipToFlg:           "H",
				ShipLocationName:    "AIS  Online Store สาขา อาคารคลงสนคา",
				CreateTransaction:   "28/01/2026",
				PayChannel:          "BRN",
				MatCode:             "3000002567",
				TaxFlag:             "N",
				CitizenId:           "1975145497083",
				FocFlg:              "N",
				FocVat:              "",
				FocAmt:              "0",
				FocSpDiscIncAmt:     "0.00",
				AddressThaiColumn:   "สาขาทออก",
				FocDesc:             "Special Discount",
				SaleType:            "PRE",
				FreeGoodsFlg:        "N",
				PayType1:            "เงนสด",
				BankName1:           "",
				Installment1:        "",
				PaymentAmount1:      "43179.00",
				PaidBy1:             "",
				BankName2:           "",
				PaymentAmount2:      "0",
				BankName3:           "",
				PaymentAmount3:      "0",
				BankName4:           "",
				PaymentAmount4:      "0",
				BankName5:           "",
				PaymentAmount5:      "0",
				BankName6:           "",
				PaymentAmount6:      "0",
				BankName7:           "",
				PaymentAmount7:      "0",
				TaxInfoSellerBranch: "(สาขาท 00031)",
				TaxInfoSellerAddr:   "160  ถนนพระรามท2แขวงแสมดำ เขตบางขนเทยน กรงเทพฯ 10150",
				TaxInfoBuyerBranch:  " ",
				TaxInfoBuyerTaxNo:   "",
				TaxInfoBuyerPhoneNo: "หมายเลขโทรศพททตดตอ: 0987184155",
				PosNo:               "",
				ACommerceFlg:        "Y",
				ServiceFlg:          "N",
				TradeIn:             "0",
				DescriptTi1:         "คาเครองเกา IMEI ",
				DescriptTi2:         "",
				AcomGenDeliveryFlg:  "Y",
				ThaiAmount:          "สหมนสามพนหนงรอยเจดสบแปดบาทเกาสบเกาสตางค",
				FgRdStatus:          "N",
				ShipLocationCode:    "4289",
				Total:               "2992.08",
				GrandTotal:          "2992.08",
				EReceiptFlg:         "Y",
				CnRmaAmount:         "179.00",
				SaleOrderNo:         "1000003077",
			},
		},
	}
	return response, nil
}
