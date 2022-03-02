package conf

type PrivateSpaceResource struct {
	PrivateSpacePath  string `yaml:"privateSpacePath"`
	PrivateSpacesPath string `yaml:"privateSpacesPath"`
	FabricPath        string `yaml:"fabricPath"`
	FabricsPath       string `yaml:"fabricsPath"`
}
