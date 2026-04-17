package config

import (
	"fmt"
	"os"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	DBHost     string `envconfig:"DB_HOST" required:"true"`
	DBPort     string `envconfig:"DB_PORT" required:"true"`
	DBUser     string `envconfig:"DB_USER" required:"true"`
	DBPassword string `envconfig:"DB_PASSWORD" required:"true"`
	DBName     string `envconfig:"DB_NAME" required:"true"`
	CACert     []byte
	ServerCert []byte
	ServerKey  []byte
}

func New() (*Config, error) {
	var c Config
	err := envconfig.Process("", &c)
	if err != nil {
		return nil, fmt.Errorf("error processing environment variables: %v", err)
	}

	c.CACert, err = os.ReadFile("/certs/rootCA.crt")
	if err != nil {
		return nil, fmt.Errorf("error reading CA certificate: %v", err)
	}

	c.ServerCert, err = os.ReadFile("/certs/server.crt")
	if err != nil {
		return nil, fmt.Errorf("error reading server certificate: %v", err)
	}
	c.ServerKey, err = os.ReadFile("/certs/server.key")
	if err != nil {
		return nil, fmt.Errorf("error reading server key: %v", err)
	}

	return &c, nil
}
