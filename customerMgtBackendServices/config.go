package main

import (
	"github.com/BurntSushi/toml"
	"os"
	"log"
)
// Info from config file
type Config struct {
	broker_frontend_url_port   string
	broker_backend_url_port   string
	proxy_frontend_url_port   string
	proxy_backend_url_port   string	
}

// Reads info from config file
func ReadConfig(configfile string) Config {
	//var configfile = flags.Configfile
	_, err := os.Stat(configfile)
	if err != nil {
		log.Fatal("Config file is missing: ", configfile)
	}

	var config Config
	if _, err := toml.DecodeFile(configfile, &config); err != nil {
		log.Fatal(err)
	}
	//log.Print(config.Index)
	return config
}
