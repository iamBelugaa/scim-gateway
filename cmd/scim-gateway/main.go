package main

import (
	"fmt"
	"log"

	"github.com/iamBelugaa/scim-gateway/internal/config"
	"github.com/iamBelugaa/scim-gateway/pkg/logger"
)

func main() {
	if err := run(); err != nil {
		log.Fatalln("application error", err)
	}
}

func run() error {
	conf := config.Load()

	logger, err := logger.NewWithConfig(conf.Application, conf.Logging)
	if err != nil {
		return fmt.Errorf("failed to construct logger : %w", err)
	}

	logger.Infow(fmt.Sprintf("starting %s service", conf.Application.Service))
	return nil
}
