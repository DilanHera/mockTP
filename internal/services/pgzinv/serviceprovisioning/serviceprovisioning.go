package serviceprovisioning

import (
	"encoding/json"

	"github.com/DilanHera/mockTP/internal/app"
	pgzinvmodel "github.com/DilanHera/mockTP/internal/services/pgzinv/model"
)

type ServiceProvisioning interface {
	LockNumberByCriteria(input *LockNumberByCriteriaRequestResourceItem, requestHeader pgzinvmodel.HeaderServiceProvisioning) (LockNumberByCriteriaResponse, error)
	LockNumberByMobile(input *LockNumberByMobileRequestResourceItem, requestHeader pgzinvmodel.HeaderServiceProvisioning) (*LockNumberByMobileResponse, error)
	ClearNumberPreparation(input *ClearNumberPreparationRequestResourceItem, requestHeader pgzinvmodel.HeaderServiceProvisioning) (*ClearNumberPreparationResponse, error)
	QuerySimInfo(input *QuerySimInfoRequestResourceItem, requestHeader pgzinvmodel.HeaderServiceProvisioning) (*QuerySimInfoResponse, error)
	RequestPrepNo(input *RequestPrepNoRequestResourceItem, requestHeader pgzinvmodel.HeaderServiceProvisioning) (*RequestPrepNoResponse, error)
	ConfirmPreparation(input *ConfirmPreparationRequestResourceItem, requestHeader pgzinvmodel.HeaderServiceProvisioning) (ConfirmPreparationResponse, error)
	// Set mock response from user
	SetUserLockNumberByCriteriaPrepaid(json.RawMessage) error
	SetUserLockNumberByCriteriaPostpaid(json.RawMessage) error
	SetUserLockNumberByMobilePrepaid(json.RawMessage) error
	SetUserLockNumberByMobilePostpaid(json.RawMessage) error
	SetUserClearNumberPreparationPrepaid(json.RawMessage) error
	SetUserClearNumberPreparationPostpaid(json.RawMessage) error
	SetUserQuerySimInfo(json.RawMessage) error
	SetUserRequestPrepNoPrepaid(json.RawMessage) error
	SetUserRequestPrepNoPostpaid(json.RawMessage) error
	SetUserConfirmPreparationPrepaid(json.RawMessage) error
	SetUserConfirmPreparationPostpaid(json.RawMessage) error
}

type serviceProvisioning struct {
	app *app.App
}

func NewServiceProvisioning(app *app.App) ServiceProvisioning {
	return &serviceProvisioning{
		app: app}
}
