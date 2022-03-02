package handlers

type OrganizationHandler interface {
	GetCurrentOrganizationUsage() error
	GetAllOrganizationsUsage() error
	GetSingleOrganizationUsage(orgId string) error
	GetCurrentOrganizationQuotas() error
}
