package main

import (
	"log"

	"github.com/pkg/errors"

	linkRESTGateway "github.com/links-123/links123/app/gateways/link/link-rest-gtw"
	linkService "github.com/links-123/links123/app/services/link"

	"github.com/links-123/links123/shared/configuration"
	"github.com/links-123/links123/shared/registry"
	"github.com/links-123/links123/shared/version"
)

func main() {
	behavior, config, err := configuration.LoadSharedConfigurations()
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to startup parse configurations"))
	}

	//
	// Check for version request
	//
	if behavior.ShowVersionOnly {
		log.Print(version.GetVersion())
		return
	}

	//
	// Select service
	//
	registryContainer := registry.NewRegistryContainer()

	registryContainer.Add(linkRESTGateway.ServiceName, linkRESTGateway.FactoryMethod)
	registryContainer.Add(linkService.ServiceName, linkService.FactoryMethod)

	serviceFactory, err := registryContainer.Get(behavior.ServiceKind)
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to initialize service"))
	}

	//
	// Create service
	//
	service, err := serviceFactory(config)
	if err != nil {
		log.Fatal(errors.Wrapf(err, "failed to initialize application [%s]", behavior.ServiceKind))
	}

	//
	// Run till the death comes
	//
	log.Printf("%s, service [%s] is started", version.GetVersion(), behavior.ServiceKind)
	if err = service.Run(); err != nil {
		log.Fatal(errors.Wrapf(err, "service [%s] stopped unexpectedly due to error", behavior.ServiceKind))
	}
	log.Printf("service [%s] gracefully shutted down, bye bye!", behavior.ServiceKind)
}
