package serviceprovisioning

import (
	"github.com/DilanHera/mockTP/internal/app"
	pgzinvmodel "github.com/DilanHera/mockTP/internal/services/pgzinv/model"
)

var resources = []string{"lockNumberByCriteriaPrepaid", "lockNumberByCriteriaPostpaid", "lockNumberByMobilePrepaid",
	"lockNumberByMobilePostpaid", "clearNumberPreparationPrepaid", "clearNumberPreparationPostpaid", "querySimInfo",
	"requestPrepNoPrepaid", "requestPrepNoPostpaid", "confirmPreparationPrepaid", "confirmPreparationPostpaid"}

type ServiceProvisioning interface {
	LockNumberByCriteria(input *LockNumberByCriteriaRequestResourceItem, requestHeader pgzinvmodel.HeaderServiceProvisioning) (LockNumberByCriteriaResponse, error)
	LockNumberByMobile(input *LockNumberByMobileRequestResourceItem, requestHeader pgzinvmodel.HeaderServiceProvisioning) (*LockNumberByMobileResponse, error)
	ClearNumberPreparation(input *ClearNumberPreparationRequestResourceItem, requestHeader pgzinvmodel.HeaderServiceProvisioning) (*ClearNumberPreparationResponse, error)
	QuerySimInfo(input *QuerySimInfoRequestResourceItem, requestHeader pgzinvmodel.HeaderServiceProvisioning) (*QuerySimInfoResponse, error)
	RequestPrepNo(input *RequestPrepNoRequestResourceItem, requestHeader pgzinvmodel.HeaderServiceProvisioning) (*RequestPrepNoResponse, error)
	ConfirmPreparation(input *ConfirmPreparationRequestResourceItem, requestHeader pgzinvmodel.HeaderServiceProvisioning) (ConfirmPreparationResponse, error)
}

type serviceProvisioning struct {
	app *app.App
}

func NewServiceProvisioning(app *app.App) ServiceProvisioning {
	app.Service.InitServiceStore(resources)
	return &serviceProvisioning{
		app: app}
}
