package main

import (
	"flag"
	"fmt"
	"go.uber.org/zap"
	"os"
)

const (
	PathConfigDefault = "./config.json"
	AppVersion        = "1.0.0"
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
	logger := initLogger(cfg)
	mainLogger := logger.Named("main")
	defer func() {
		err := logger.Sync()
		if err != nil {
			panic(fmt.Sprintf("Failed to sync logger with error: %s", err.Error()))
		}
	}()
	mainLogger.Info("Application started")
}

func initArguments() *Arguments {
	args := &Arguments{}

	flag.StringVar(&args.Config, "c", PathConfigDefault, "Path to config file.")
	flag.BoolVar(&args.VersionCalled, "v", false, "Prints application version and exit")

	flag.Parse()

	return args
}

func initConfig(cfgPath string) *Config {
	cfg, err := NewConfig(cfgPath)
	if err != nil {
		if _, ok := err.(*os.PathError); ok {
			panic(fmt.Sprintf("ServerConfig not found at path: %s", cfgPath))
		} else {
			panic(fmt.Sprintf("ServerConfig is corrupted. Error: %s", err.Error()))
		}
	}

	return cfg
}

func initLogger(cfg *Config) *zap.SugaredLogger {
	logger, err := cfg.Logger.Build()
	if err != nil {
		panic(fmt.Sprintf("Logger config is corrupted. Error: %s", err.Error()))
	}

	return logger.Sugar()
}

func printVersion() {
	fmt.Println(fmt.Sprintf("Gerret's Social Project (Otus HLA) v%s", AppVersion))
	fmt.Println("© 2022 #gGerret")
}
