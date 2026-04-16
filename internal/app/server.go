package server

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"traffic-monitor/internal/config"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

func Run(conf *config.Config) error {
	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.RequestLogger())

	e.Validator = &CustomValidator{validator: validator.New(validator.WithRequiredStructEnabled())}
	defineRoutes(e)

	// Use mTLS to ensure clients are valid
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(conf.CACert)

	tlsConfig := &tls.Config{
		ClientCAs:  caCertPool,
		ClientAuth: tls.RequireAndVerifyClientCert,
		MinVersion: tls.VersionTLS13,
		CipherSuites: []uint16{
			// Restrict to secure TLS 1.3 cipher suite
			tls.TLS_AES_256_GCM_SHA384,
		},
	}

	sc := echo.StartConfig{
		Address:   ":8443",
		TLSConfig: tlsConfig,
	}

	// Serve the app at port 8443
	err := sc.StartTLS(context.Background(), e, conf.ServerCert, conf.ServerKey)
	if err != nil {
		return fmt.Errorf("failed to start TLS server: %v", err)
	}

	return nil
}
