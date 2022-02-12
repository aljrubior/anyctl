package response

type Profile struct {
	Id                      string                             `json:"id"`
	OrganizationId          string                             `json:"organizationId"`
	OrganizationPreferences map[string]OrganizationPreferences `json:"organizationPreferences"`
	Organization            organization                       `json:"organization"`
}

func (this Profile) GetDefaultEnvironment() Environment {
	defaultEnvId := this.OrganizationPreferences[this.OrganizationId].DefaultEnvironment

	for _, v := range this.Organization.Environments {
		if v.Id == defaultEnvId {
			return v
		}
	}

	return this.Organization.Environments[0]
}
