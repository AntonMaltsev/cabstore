package main

import (
	"errors"
	"github.com/codegangsta/cli"
	"gopkg.in/yaml.v1"
	"io/ioutil"
	log "gopkg.in/inconshreveable/log15.v2"
	"os"
	"github.com/antonmaltsev/cabstore/service"
	config "github.com/antonmaltsev/cabstore/cfg"
)

// Getting configuration file data
func getConfig(c *cli.Context) (config.Config, error) {
	yamlPath := c.GlobalString("config")
	config := config.Config{}

	if _, err := os.Stat(yamlPath); err != nil {
		return config, errors.New("config path not valid")
	}

	ymlData, err := ioutil.ReadFile(yamlPath)
	if err != nil {
		return config, err
	}

	err = yaml.Unmarshal([]byte(ymlData), &config)
	return config, err
}


func main() {

	app := cli.NewApp()
	app.Name = "Cabify Store"
	app.Usage = "Cabify"
	app.Version = "0.0.1"

	app.Flags = []cli.Flag{
		cli.StringFlag{"config, c", "config.yaml", "config file to use", "APP_CONFIG"},
	}

	app.Commands = []cli.Command{
		{
			Name:  "server",
			Usage: "Run Cabify Store http server",
			Action: func(c *cli.Context) {
				cfg, err := getConfig(c)
				if err != nil {
					log.Error(err.Error())
					return
				}

				svc := service.CabifyService{}

				if err = svc.Run(cfg); err != nil {
					log.Error(err.Error())
				}

				log.Info("Cabify Store server started")
			},
		},
	}
	app.Run(os.Args)
}
