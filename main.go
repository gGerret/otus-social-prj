package main

import (
	"flag"
	"fmt"
	"github.com/gGerret/otus-social-prj/config"
	"github.com/gGerret/otus-social-prj/repository"
	"github.com/gGerret/otus-social-prj/router"
	"github.com/gGerret/otus-social-prj/social"
	"os"
)

/*
 * metrics library "github.com/hashicorp/go-metrics"
 * OpenAPI swagger lib "github.com/swaggo/gin-swagger"
 */

const (
	PathConfigDefault = "./config.json"
	AppName           = "Gerret's Social Project (Otus HLA)"
	AppVersion        = "1.0.0"
	Copyright         = "© 2022 #gGerret"
)

type Arguments struct {
	Config        string
	VersionCalled bool
}

func main() {
	args := initArguments()

	if args.VersionCalled {
		printVersion()
		os.Exit(0)
	}

	cfg := initConfig(args.Config)
	logger := social.InitLogger(cfg.Logger)
	mainLogger := logger.Named("main")
	defer func() {
		err := logger.Sync()
		if err != nil {
			panic(fmt.Sprintf("Failed to sync logger with error: %s", err.Error()))
		}
	}()
	mainLogger.Info("Application started")

	socialWeb, err := router.NewSocialServer(cfg.Server, mainLogger)
	if err != nil {
		mainLogger.DPanicf("Failed to initialize server with error: %s", err.Error())
	}
	mainLogger.Info("Initialize database connection...")
	repository.InitDb(cfg.Db)
	mainLogger.Infof("Database '%s' initialized", repository.DbDriver)

	err = socialWeb.RunServer()
	if err != nil {
		mainLogger.DPanicf("Server error occured: %s", err.Error())
	}

	mainLogger.Info("Application shutdown")
}

func initArguments() *Arguments {
	args := &Arguments{}

	flag.StringVar(&args.Config, "c", PathConfigDefault, "Path to config file.")
	flag.BoolVar(&args.VersionCalled, "v", false, "Prints application version and exit")

	flag.Parse()

	return args
}

func initConfig(cfgPath string) *config.Config {
	cfg, err := config.NewConfig(cfgPath)
	if err != nil {
		if _, ok := err.(*os.PathError); ok {
			panic(fmt.Sprintf("ServerConfig not found at path: %s", cfgPath))
		} else {
			panic(fmt.Sprintf("ServerConfig is corrupted. Error: %s", err.Error()))
		}
	}

	return cfg
}

func printVersion() {
	fmt.Println(fmt.Sprintf("%s v%s", AppName, AppVersion))
	fmt.Println(Copyright)
}
