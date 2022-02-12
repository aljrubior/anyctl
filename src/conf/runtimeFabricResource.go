package conf

type RuntimeFabricResource struct {
	FabricPath             string `yaml:"fabricPath"`
	FabricsPath            string `yaml:"fabricsPath"`
	TargetPath             string `yaml:"targetPath"`
	TargetsPath            string `yaml:"targetsPath"`
	PrivateSpacePath       string `yaml:"privateSpacePath"`
	PrivateSpacesPath      string `yaml:"privateSpacesPath"`
	PrivateSpaceFabricPath string `yaml:"privateSpaceFabricsPath"`
}
