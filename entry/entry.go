package main

import (
	"fmt"
	"os"

	"github.com/pkg/errors"

	linkRESTGateway "github.com/links-123/links123/app/gateways/link/link-rest-gtw"
	linkService "github.com/links-123/links123/app/services/link"

	"github.com/links-123/links123/shared/configuration"
	"github.com/links-123/links123/shared/logging"
	"github.com/links-123/links123/shared/registry"
	"github.com/links-123/links123/shared/version"
)

func main() {
	behavior, config, err := configuration.LoadSharedConfigurations()
	if err != nil {
		fmt.Println(errors.Wrap(err, "failed to startup parse configurations"))
		os.Exit(1)
	}

	//
	// Check for version request
	//
	if behavior.ShowVersionOnly {
		fmt.Print(version.GetVersion())
		return
	}

	//
	// Initialize logger
	//
	logger, err := logging.NewLogger(config)
	if err != nil {
		fmt.Println(errors.Wrap(err, "failed to initialize logging"))
		os.Exit(1)
	}

	//
	// Select service
	//
	registryContainer := registry.NewRegistryContainer()

	registryContainer.Add(linkRESTGateway.ServiceName, linkRESTGateway.FactoryMethod)
	registryContainer.Add(linkService.ServiceName, linkService.FactoryMethod)

	serviceFactory, err := registryContainer.Get(behavior.ServiceKind)
	if err != nil {
		logger.Fatal(errors.Wrap(err, "failed to initialize service"))
	}

	//
	// Create service
	//
	service, err := serviceFactory(config, logger)
	if err != nil {
		logger.Fatal(errors.Wrapf(err, "failed to initialize application [%s]", behavior.ServiceKind))
	}

	//
	// Run till the death comes
	//
	logger.Infof("%s, service [%s] is started", version.GetVersion(), behavior.ServiceKind)
	if err = service.Run(); err != nil {
		logger.Fatal(errors.Wrapf(err, "service [%s] stopped unexpectedly due to error", behavior.ServiceKind))
	}
	logger.Infof("service [%s] gracefully shutted down, bye bye!", behavior.ServiceKind)
}
