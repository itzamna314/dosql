package main

import (
	"fmt"
	"github.com/BurntSushi/toml"
)

const (
	connStrFmt = `server=%s;port=%s;database=%s;user id=%s;password=%s;encrypt=%s;TrustServerCertificate=%s;connection timeout=%s;`
)

type config struct {
	Server                 string `toml:"server"`
	Port                   string `toml:"port"`
	Database               string `toml:"database"`
	UserId                 string `toml:"user_id"`
	Password               string `toml:"password"`
	Encrypt                string `toml:"encrypt"`
	TrustServerCertificate string `toml:"trust_server_certificate"`
	ConnectionTimeout      string `toml:"connection_timeout"`
	Driver                 string `toml:"driver"`
}

var environments map[string]config

func LoadConnectionString(configFile, environment string) (string, string, error) {
	if _, err := toml.DecodeFile(configFile, &environments); err != nil {
		return "", "", err
	}

	c, ok := environments[environment]
	if !ok {
		return "", "", fmt.Errorf("Environment not found: %s", environment)
	}

	connStr := fmt.Sprintf(connStrFmt, c.Server, c.Port, c.Database, c.UserId, c.Password, c.Encrypt, c.TrustServerCertificate, c.ConnectionTimeout)

	return connStr, c.Driver, nil
}
