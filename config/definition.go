package config

type Config struct {
	Mongo Mongo `yaml:"mongo"`
}

type Mongo struct {
	Database   string `yaml:"database"`
	User       string `yaml:"user"`
	Password   string `yaml:"password"`
	Collection string `yaml:"collection"`
	Conn       string `yaml:"conn"`
}
