package dt

import "fmt"

type QueryStockImeiMyStoreRequest struct {
	LocationCode string `json:"locationCode" validate:"required"`
	SubStock     string `json:"subStock" validate:"required"`
	ListSerialNo string `json:"listSerialNo" validate:"required"`
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
	result := d.GetApiInfo("queryStockImeiMyStore")
	if result.State == "C" {
		if UserQueryStockImeiMyStore != nil {
			return UserQueryStockImeiMyStore, nil
		}
		return nil, fmt.Errorf("no custom response set for queryStockImeiMyStore")
	}

	if result.State == "E" {
		return &QueryStockImeiMyStoreResponse{
			ResultCode:        "50000",
			ResultDescription: "Product not found",
			DeveloperMessage:  "Product not found",
			ListProduct:       []ListProduct{},
		}, nil
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
