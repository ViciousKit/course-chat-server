package config

import (
	"flag"
	"os"
)

const defaultConfigPath = "local.env"

func (p *Path) Load() error {
	var path string

	flag.StringVar(&path, "config", "", "path to cfg file")
	flag.Parse()

	if path == "" {
		path = os.Getenv("CONFIG_FILE")
	}

	if path != "" {
		p.path = path
	} else {
		p.path = defaultConfigPath
	}

	return nil
}
