package conf

type AnypointResources struct {
	RuntimeManager RuntimeManagerResources `yaml:"runtimeManager"`
	Admin          AdminResources          `yaml:"admin"`
}
