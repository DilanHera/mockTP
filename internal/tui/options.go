package tui

var Services = []string{"PGZINV", "PHX", "DT"}

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

var PHXApis = []string{"requestESIM", "newRegistration"}
