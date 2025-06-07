package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/sitnikovik/lintcommit/internal/config"
)

const (
	defaultConfigFile = "lintcommit.yaml"
)

func main() {
	// Parse command-line flags.
	// Use a flag to specify the path to the configuration file
	configPath := flag.String(
		"config",
		defaultConfigFile,
		fmt.Sprintf(
			"Path to config file. Default is %s",
			defaultConfigFile,
		),
	)
	flag.Parse()

	// Load the configuration from the YAML file
	_, err := config.LoadFrom(*configPath)
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}
}
