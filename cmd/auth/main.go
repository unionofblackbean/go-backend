package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	_ "embed"

	"github.com/unionofblackbean/backend/pkg/database"
	"github.com/unionofblackbean/backend/services/auth"
	"github.com/unionofblackbean/backend/services/auth/rest"
	"github.com/urfave/cli/v2"
)

//go:embed schema.sql
var databaseSchema string

var config = auth.Config{
	Rest: auth.RestConfig{
		BindAddress: "127.0.0.1",
		BindPort:    8080,
	},
	Database: auth.DatabaseConfig{
		Address: "127.0.0.1",
		Port:    5432,

		Username: "auth",
		Password: "auth",

		Name: "auth",
	},
}

func main() {
	app := &cli.App{
		Usage: "Union of Black Bean backend auth service",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "config",
				Usage: "configuration file path",
			},
			&cli.BoolFlag{
				Name:  "init-db",
				Usage: "initialize database (execute schema)",
			},
		},
		Action: func(ctx *cli.Context) error {
			configPath := ctx.String("config")
			if configPath != "" {
				configBytes, err := ioutil.ReadFile(configPath)
				if err != nil {
					return fmt.Errorf("failed to open config file -> %v", err)
				}

				err = json.Unmarshal(configBytes, &config)
				if err != nil {
					return fmt.Errorf("failed to parse config file -> %v", err)
				}
			}

			pool, err := database.NewPool(
				config.Database.Username, config.Database.Password,
				config.Database.Address, config.Database.Port,
				config.Database.Name,
			)
			if err != nil {
				return fmt.Errorf("failed to establish connection with database -> %v", err)
			}

			if ctx.Bool("init-db") {
				err := pool.Exec(databaseSchema)
				if err != nil {
					return fmt.Errorf("failed to execute database schema -> %v", err)
				}
			}

			rest.Init(pool)
			rest.Run(config.Rest.BindAddress, config.Rest.BindPort)

			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
