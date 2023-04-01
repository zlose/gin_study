package core

import (
	"fmt"
	"gin_study/config"
	"gin_study/global"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

// init config
func Initconf() {
	const configfile = "settings.yaml"
	c := &config.Config{}
	yamlConf, err := os.ReadFile(configfile)
	if err != nil {
		panic(fmt.Errorf("failed get config:%s", err))
	}
	err = yaml.Unmarshal(yamlConf, c)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("config file success init... ")
	global.Config = c
}
