package conf

type AssetClientConfig struct {
	Protocol            string
	Host                string
	Port                int
	AssetsPath          string
	LatestVersionPath   string
	SpecificVersionPath string
	UploadAssetPath     string
}
