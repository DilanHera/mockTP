package dt

type QueryStockImeiMyStoreRequest struct {
	LocationCode string `json:"locationCode"`
	SubStock     string `json:"subStock"`
	ListSerialNo string `json:"listSerialNo"`
}

type QueryStockImeiMyStoreResponse struct {
	ResultCode        string        `json:"resultCode"`
	ResultDescription string        `json:"resultDescription"`
	DeveloperMessage  string        `json:"developerMessage"`
	ListProduct       []ListProduct `json:"listProduct"`
}

type ListProduct struct {
	LocationCode   string `json:"locationCode"`
	Company        string `json:"company"`
	ProductType    string `json:"productType"`
	ProductSubtype string `json:"productSubtype"`
	Brand          string `json:"brand"`
	Model          string `json:"model"`
	UnitName       string `json:"unitName"`
	MatCode        string `json:"matCode"`
	SubStockCode   string `json:"subStockCode"`
	SerialNo       string `json:"serialNo"`
	SapDescription string `json:"sapDescription"`
	VatType        string `json:"vatType"`
	MatType        string `json:"matType"`
}

func (d *dt) QueryStockImeiMyStore(input *QueryStockImeiMyStoreRequest) (*QueryStockImeiMyStoreResponse, error) {
	if UserQueryStockImeiMyStore != nil {
		return UserQueryStockImeiMyStore, nil
	}
	response := &QueryStockImeiMyStoreResponse{
		ResultCode:        "20000",
		ResultDescription: "Success",
		DeveloperMessage:  "Success",
		ListProduct: []ListProduct{
			{
				LocationCode:   "4289",
				Company:        "WDS",
				ProductType:    "SIM",
				ProductSubtype: "PREPAID",
				Brand:          "OTC",
				Model:          "OTC110",
				UnitName:       "PC",
				MatCode:        "SIMSERVICEOPTIMUS",
				SubStockCode:   "OLS",
				SerialNo:       "000001920041336559",
				SapDescription: "SIM OTC 110B เบอร์สวยพนักงาน - EC",
				VatType:        "Y",
				MatType:        "Serial",
			},
		},
	}
	return response, nil
}
