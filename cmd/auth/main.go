package main

import (
	"fmt"
	"log"
	"os"

	_ "embed"

	"github.com/unionofblackbean/backend/pkg/config"
	"github.com/unionofblackbean/backend/pkg/database"
	"github.com/unionofblackbean/backend/services/auth"
	"github.com/unionofblackbean/backend/services/auth/rest"
	"github.com/urfave/cli/v2"
)

//go:embed schema.sql
var databaseSchema string

var conf = auth.Config{
	Rest: auth.RestConfig{
		BindAddress: "127.0.0.1",
		BindPort:    8000,
	},
	Database: database.Config{
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
				err := config.Load(configPath, &conf)
				if err != nil {
					return fmt.Errorf("failed to load config file -> %v", err)
				}
			}

			pool, err := database.NewPool(
				conf.Database.Username, conf.Database.Password,
				conf.Database.Address, conf.Database.Port,
				conf.Database.Name,
			)
			if err != nil {
				return fmt.Errorf("failed to create database connection pool -> %v", err)
			}

			if ctx.Bool("init-db") {
				err := pool.Exec(databaseSchema)
				if err != nil {
					return fmt.Errorf("failed to execute database schema -> %v", err)
				}
			}

			rest.Init(pool)
			rest.Run(conf.Rest.BindAddress, conf.Rest.BindPort)

			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
