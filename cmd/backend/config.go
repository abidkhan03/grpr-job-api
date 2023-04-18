package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/grpr-job-api/internal/errors"
)

type Config struct {
	HttpPort int `yaml:"httpPort"`
}

func readConfig() (Config, error) {
	var c Config

	flag.IntVar(&c.HttpPort, "port", 2020, "Port for http server")
	flag.Parse()

	remainingArgs := flag.Args()
	if len(remainingArgs) > 0 {
		fmt.Println("Unknown arguments:", strings.Join(remainingArgs, " "))
		os.Exit(-1)
	}

	if x, err := strconv.ParseInt(os.Getenv("HTTP_PORT"), 10, 64); err != nil {
		c.HttpPort = int(x)
	}

	if c.HttpPort == 0 {
		return Config{}, errors.New(errors.InvalidArgument, "empty server port. -port flag required")
	}

	return c, nil
}
