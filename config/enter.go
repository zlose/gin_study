package config

type config struct {
	Mysql  mysql  `yaml:"mysql"`
	Logger logger `yaml:"logger"`
	System system `yaml:"system"`
}
