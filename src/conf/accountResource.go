package conf

type AccountResource struct {
	LoginPath        string `yaml:"loginPath"`
	ProfilePath      string `yaml:"profilePath"`
	OrganizationPath string `yaml:"organizationPath"`
}
