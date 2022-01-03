package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"strings"

	"github.com/iver-wharf/wharf-core/pkg/config"
)

type Config struct {
	BindAddress string
	flag        string
	FlagFile    string
	CIDRs       []string
	cidrs       []*net.IPNet
}

var defaultConfig = Config{
	BindAddress: "0.0.0.0:8080",
	FlagFile:    "flag.txt",
	CIDRs: []string{
		"192.168.0.0/16",
		"127.0.0.0/8",
	},
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

	for _, cidr := range cfg.CIDRs {
		_, ipNet, err := net.ParseCIDR(cidr)
		if err != nil {
			return Config{}, err
		}
		cfg.cidrs = append(cfg.cidrs, ipNet)
	}
	log.Debug().WithStringf("cird", "%v", cfg.cidrs).Message("Blocking IPs based on CIRD.")

	return cfg, nil
}
