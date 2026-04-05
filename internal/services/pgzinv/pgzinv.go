package pgzinv

import (
	"errors"

	"github.com/DilanHera/mockTP/internal/app"
	pgzinvmodel "github.com/DilanHera/mockTP/internal/services/pgzinv/model"
	"github.com/DilanHera/mockTP/internal/services/pgzinv/serviceprovisioning"
)

type Pgzinv interface {
	ServiceProvisioning(input *pgzinvmodel.ServiceProvisioningPayload) (interface{}, error)
}

type pgzinv struct {
	app *app.App
}

func NewPgzinv(app *app.App) Pgzinv {
	return &pgzinv{
		app: app,
	}
}

func (p *pgzinv) ServiceProvisioning(input *pgzinvmodel.ServiceProvisioningPayload) (interface{}, error) {
	spResource := serviceprovisioning.NewServiceProvisioning()
	switch input.ResourceName {
	case "lockNumberByCriteriaPrepaid", "lockNumberByCriteriaPostpaid":
		request := &serviceprovisioning.LockNumberByCriteriaRequestResourceItem{}
		err := p.app.Helper.UnmarshalAndValidate(input.Payload, request)
		if err != nil {
			return nil, err
		}
		response, err := spResource.LockNumberByCriteria(request)
		if err != nil {
			return nil, err
		}
		// response, err := buildResponse(&resourceResponse)
		// if err != nil {
		// 	return nil, err
		// }
		return response, nil
	default:
		return nil, errors.New("invalid resource name")
	}
	// return serviceprovisioning.NewServiceProvisioningResources()
}
