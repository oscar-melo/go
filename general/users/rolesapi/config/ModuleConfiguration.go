package config

import (
	"fmt"
	"os"
	"gopkg.in/yaml.v2"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
    Server struct {
        Port string `yaml:"port", envconfig:"SERVER_PORT"`
    } `yaml:"server"`
    Database struct {
        Port string `yaml:"port", envconfig:"SERVER_HOST"`
		Host string `yaml:"host", envconfig:"SERVER_HOST"`
		Name string `yaml:"name", envconfig:"SERVER_HOST"`
		User string `yaml:"user", envconfig:"SERVER_HOST"`
		Password string `yaml:"password", envconfig:"SERVER_HOST"`
    } `yaml:"database"`
} 

var Cfg Config

func LoadConfig() {
    
    readFile(&Cfg)
    readEnv(&Cfg)
/*
	fmt.Println(Cfg.Server.Port)
	fmt.Println(Cfg.Database.Name) */
}

func processError(err error) {
    fmt.Println(err)
    os.Exit(2)
}

func readFile(Cfg *Config) {
	//Carga por defecto un archivo yml llamado config.
    f, err := os.Open("config/config.yml")
    if err != nil {
        processError(err)
    }
    defer f.Close()

    decoder := yaml.NewDecoder(f)
    err = decoder.Decode(Cfg)
    if err != nil {
        processError(err)
    }
} 

func readEnv(Cfg *Config) { 
    err := envconfig.Process("", Cfg) 
    if err != nil { 
        processError(err)
    }
}