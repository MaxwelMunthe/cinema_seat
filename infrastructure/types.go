package infrastructure

type app struct {
	Appname string `yaml:"name"`
	Debug   bool   `yaml:"debug"`
	Port    string `yaml:"port"`
	Service string `yaml:"service"`
	Host    string `yaml:"host"`
}

type CinemaConfigParam struct {
	Status string
	CreatedDate string
	BookedDate string
}

type Environment struct {
	App     app     `yaml:"app"`
	path    string
}