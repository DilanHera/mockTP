package pgzinvmodel

/* Request Model */
type HeaderServiceProvisioning struct {
	ResourceGroupId   string `json:"resourceGroupId"`
	CustomerOrderType string `json:"customerOrderType"`
	UserSys           string `json:"userSys"`
	ReTransmit        string `json:"reTransmit"`
}

type ResourceItemServiceProvisioning struct {
	ResourceName     string `json:"resourceName"`
	UserId           string `json:"userId"`
	OriginalOperator string `json:"originalOperator"`
	ChargeNode       string `json:"chargeNode"`
	CategoryCode     string `json:"categoryCode"`
	ProductType      string `json:"productType"`
	RegionCode       string `json:"regionCode"`
	NumberPattern    string `json:"numberPattern"`
	ClassifyCode     string `json:"classifyCode"`
	NumberStatusFrom string `json:"numberStatusFrom"`
	NumberStatusTo   string `json:"numberStatusTo"`
	LuckyName        string `json:"luckyName"`
	LuckyType        string `json:"luckyType"`
	Quantity         string `json:"quantity"`
	MobileNo         string `json:"mobileNo"`
	MachineNo        string `json:"machineNo"`
	JobId            string `json:"jobId"`
	Package          string `json:"package"`
	OperationType    string `json:"operationType"`
	SimSerialNo      string `json:"simSerialNo"`
	PrepNo           string `json:"prepNo"`
	ProjectName      string `json:"projectName"`
	SimType          string `json:"simType"`
	PackageRowId     string `json:"packageRowId"`
	OfferingName     string `json:"offeringName"`
	OfferingCode     string `json:"offeringCode"`
	EanCode          string `json:"eanCode"`
	MaterialCode     string `json:"materialCode"`
	ExpiryDate       string `json:"expiryDate"`
	Aging            string `json:"aging"`
	ChargeType       string `json:"chargeType"`
	SubPackageType   string `json:"subPackageType"`
	PackageType      string `json:"packageType"`
	NetworkType      string `json:"networkType"`
	LocationCode     string `json:"locationCode"`
	MatCode          string `json:"matCode"`
	SourceSystem     string `json:"sourceSystem"`
	Key              string `json:"key"`
	Date             string `json:"date"`
	SimProject       string `json:"simProject"`
	MatDesc          string `json:"matDesc"`
}

type ServiceProvisioningRequest struct {
	RequestHeader    HeaderServiceProvisioning         `json:"requestHeader"`
	ResourceItemList []ResourceItemServiceProvisioning `json:"resourceItemList"`
}

/* Transform Model */
type ServiceProvisioningPayload struct {
	ResourceName string
	Payload      []byte
}

/* Response Model */
type ResponseHeader struct {
	ResourceGroupId  string `json:"resourceGroupId" validate:"required"`
	ResourceOrderId  string `json:"resourceOrderId" validate:"required"`
	ResultCode       string `json:"resultCode" validate:"required"`
	ResultDesc       string `json:"resultDesc" validate:"required"`
	DeveloperMessage string `json:"developerMessage"`
	UserSys          string `json:"userSys" validate:"required"`
	ReTransmit       string `json:"reTransmit" validate:"required"`
}

type ResourceItemListBase struct {
	ResourceName           string             `json:"resourceName" validate:"required"`
	ResourceItemStatus     string             `json:"resourceItemStatus" validate:"required"`
	ErrorFlag              string             `json:"errorFlag" validate:"required"`
	ResourceItemErrMessage string             `json:"resourceItemErrMessage" validate:"required"`
	SpecialErrHandling     SpecialErrHandling `json:"specialErrHandling" validate:"required"`
}

type SpecialErrHandling struct {
	SuppCode             []string `json:"suppCode" validate:"required"`
	TaskKeyCondition     []string `json:"taskKeyCondition" validate:"required"`
	TaskDeveloperMessage []string `json:"taskDeveloperMessage" validate:"required"`
}

type ServiceProvisioningResponse struct {
	ResponseHeader   ResponseHeader `json:"responseHeader" validate:"required"`
	ResourceItemList []any          `json:"resourceItemList" validate:"required"`
}
