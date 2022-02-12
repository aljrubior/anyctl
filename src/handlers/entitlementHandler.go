package handlers

type EntitlementHandler interface {
	GetCurrentOrganizationEntitlement() error
}
