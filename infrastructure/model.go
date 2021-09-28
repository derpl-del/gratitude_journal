package infrastructure

type app struct {
	Appname string `yaml:"name"`
	Port    string `yaml:"port"`
}

type database struct {
	Name       string `yaml:"name"`
	Connection string `yaml:"connection"`
}

type Environment struct {
	App       app                 `yaml:"app"`
	Databases map[string]database `yaml:"databases"`
}
