package config

import (
	"fmt"
	"os"
)

type Config struct {
	CACert     []byte
	ServerCert []byte
	ServerKey  []byte
}

func New() (*Config, error) {
	caCert, err := os.ReadFile("./certs/rootCA.crt")
	if err != nil {
		return nil, fmt.Errorf("error reading CA certificate: %v", err)
	}

	serverCert, err := os.ReadFile("./certs/server.crt")
	if err != nil {
		return nil, fmt.Errorf("error reading server certificate: %v", err)
	}
	serverKey, err := os.ReadFile("./certs/server.key")
	if err != nil {
		return nil, fmt.Errorf("error reading server key: %v", err)
	}

	return &Config{
		CACert:     caCert,
		ServerCert: serverCert,
		ServerKey:  serverKey,
	}, nil
}
