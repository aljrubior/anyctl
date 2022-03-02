package conf

type AssetResource struct {
	AssetsPath          string `yaml:"assetsPath"`
	LatestVersionPath   string `yaml:"latestVersionPath"`
	SpecificVersionPath string `yaml:"specificVersionPath"`
	UploadAssetPath     string `yaml:"uploadAssetPath"`
}
