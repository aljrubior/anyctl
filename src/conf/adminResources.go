package conf

type AdminResources struct {
	SharedSpaces  SharedSpaceResource  `yaml:"sharedSpaces"`
	PrivateSpaces PrivateSpaceResource `yaml:"privateSpaces"`
	Fabrics       FabricResource       `yaml:"fabrics"`
	Accounts      AccountResource      `yaml:"accounts"`
}
