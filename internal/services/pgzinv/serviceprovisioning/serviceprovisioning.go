package serviceprovisioning

type ServiceProvisioning interface {
	LockNumberByCriteria(input *LockNumberByCriteriaRequestResourceItem) (LockNumberByCriteriaResponse, error)
}

type serviceProvisioning struct {
}

func NewServiceProvisioning() ServiceProvisioning {
	return &serviceProvisioning{}
}
