package config

type Config struct {
	Mysql  mysql  `yaml:"mysql"`
	Logger logger `yaml:"logger"`
	System system `yaml:"system"`
}
