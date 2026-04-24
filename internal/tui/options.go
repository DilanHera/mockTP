package tui

var Services = []string{"PGZINV", "PHX", "DT", "IM", "ESB"}

var PgzinvApis = []string{"serviceProvisioning"}

var ServiceProvisioningResources = []string{
	"lockNumberByCriteriaPrepaid",
	"lockNumberByCriteriaPostpaid",
	"lockNumberByMobilePrepaid",
	"lockNumberByMobilePostpaid",
	"clearNumberPreparationPrepaid",
	"clearNumberPreparationPostpaid",
	"querySimInfo",
	"requestPrepNoPrepaid",
	"requestPrepNoPostpaid",
	"confirmPreparationPrepaid",
	"confirmPreparationPostpaid",
}

var PHXApis = []string{"requestESIM", "newRegistration", "encryptLib", "checkPerso", "productProvisioning"}

var DTApis = []string{"listOrderNoByDono", "pickingDocument", "queryPrint", "queryStockImeiMyStore", "reprintReceiptForm", "updateSimSerialPerso", "authenticate"}

var IMApis = []string{"sendSimSerialNo"}

var ESBApis = []string{
	"oauthToken",
	"createFreightOrder",
	"doCreation",
	"legoUpdateOrderStatus",
	"persoSim",
	"serialNumberExpirationDate",
}
