package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/iver-wharf/wharf-core/pkg/config"
)

type Config struct {
	BindAddress string
	flag        string
	FlagFile    string
	Equations   int
	RngSeed     *int64
	PublicPort  int
	PublicHost  string
}

var defaultConfig = Config{
	BindAddress: "0.0.0.0:7070",
	FlagFile:    "flag.txt",
	Equations:   127,
	RngSeed:     nil,
	PublicPort:  7070,
	PublicHost:  "localhost",
}

func loadConfig() (Config, error) {
	cfgBuilder := config.NewBuilder(defaultConfig)
	cfgBuilder.AddConfigYAMLFile("./conf.yaml")
	cfgBuilder.AddEnvironmentVariables("CTF")
	var cfg Config
	if err := cfgBuilder.Unmarshal(&cfg); err != nil {
		return Config{}, err
	}

	flagFile, err := os.Open("flag.txt")
	if err != nil {
		return Config{}, fmt.Errorf("open flag file: %w", err)
	}
	defer flagFile.Close()
	flagBytes, err := io.ReadAll(flagFile)
	if err != nil {
		return Config{}, fmt.Errorf("read flag file: %w", err)
	}
	cfg.flag = strings.TrimSpace(string(flagBytes))

	return cfg, nil
}
