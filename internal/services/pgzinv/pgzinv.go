package pgzinv

import (
	"errors"

	"github.com/DilanHera/mockTP/internal/app"
	pgzinvmodel "github.com/DilanHera/mockTP/internal/services/pgzinv/model"
	"github.com/DilanHera/mockTP/internal/services/pgzinv/serviceprovisioning"
)

type Pgzinv interface {
	ServiceProvisioning(input *pgzinvmodel.ServiceProvisioningPayload) (any, error)
}

type pgzinv struct {
	app *app.App
}

func NewPgzinv(app *app.App) Pgzinv {
	return &pgzinv{
		app: app,
	}
}

func (p *pgzinv) ServiceProvisioning(input *pgzinvmodel.ServiceProvisioningPayload) (any, error) {
	spResource := serviceprovisioning.NewServiceProvisioning(p.app)
	switch input.ResourceName {
	case "lockNumberByCriteriaPrepaid", "lockNumberByCriteriaPostpaid":
		request := &serviceprovisioning.LockNumberByCriteriaRequestResourceItem{}
		err := p.app.Helper.UnmarshalAndValidate(input.Payload, request)
		if err != nil {
			return nil, err
		}
		response, err := spResource.LockNumberByCriteria(request, input.RequestHeader)
		if err != nil {
			return nil, err
		}
		// response, err := buildResponse(&resourceResponse)
		// if err != nil {
		// 	return nil, err
		// }
		return response, nil
	case "lockNumberByMobilePrepaid", "lockNumberByMobilePostpaid":
		request := &serviceprovisioning.LockNumberByMobileRequestResourceItem{}
		err := p.app.Helper.UnmarshalAndValidate(input.Payload, request)
		if err != nil {
			return nil, err
		}
		response, err := spResource.LockNumberByMobile(request, input.RequestHeader)
		if err != nil {
			return nil, err
		}
		return response, nil
	case "clearNumberPreparationPrepaid", "clearNumberPreparationPostpaid":
		request := &serviceprovisioning.ClearNumberPreparationRequestResourceItem{}
		err := p.app.Helper.UnmarshalAndValidate(input.Payload, request)
		if err != nil {
			return nil, err
		}
		response, err := spResource.ClearNumberPreparation(request, input.RequestHeader)
		if err != nil {
			return nil, err
		}
		return response, nil
	case "querySimInfo":
		request := &serviceprovisioning.QuerySimInfoRequestResourceItem{}
		err := p.app.Helper.UnmarshalAndValidate(input.Payload, request)
		if err != nil {
			return nil, err
		}
		response, err := spResource.QuerySimInfo(request, input.RequestHeader)
		if err != nil {
			return nil, err
		}
		return response, nil
	case "requestPrepNoPrepaid", "requestPrepNoPostpaid":
		request := &serviceprovisioning.RequestPrepNoRequestResourceItem{}
		err := p.app.Helper.UnmarshalAndValidate(input.Payload, request)
		if err != nil {
			return nil, err
		}
		response, err := spResource.RequestPrepNo(request, input.RequestHeader)
		if err != nil {
			return nil, err
		}
		return response, nil
	case "confirmPreparationPrepaid", "confirmPreparationPostpaid":
		request := &serviceprovisioning.ConfirmPreparationRequestResourceItem{}
		err := p.app.Helper.UnmarshalAndValidate(input.Payload, request)
		if err != nil {
			return nil, err
		}
		response, err := spResource.ConfirmPreparation(request, input.RequestHeader)
		if err != nil {
			return nil, err
		}
		return response, nil
	default:
		return nil, errors.New("invalid resource name")
	}
	// return serviceprovisioning.NewServiceProvisioningResources()
}
