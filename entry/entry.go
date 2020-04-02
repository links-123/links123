package main

import (
	"github.com/links-123/links123/shared/version"
	"log"

	linkRESTGateway "github.com/links-123/links123/app/gateways/link/link-rest-gtw"
	linkService "github.com/links-123/links123/app/services/link"

	"github.com/pkg/errors"

	"github.com/links-123/links123/registry"
	"github.com/links-123/links123/shared/cmd"
	"github.com/links-123/links123/shared/env"
)

func main() {
	//
	// Load startup flags
	//
	flags := cmd.LoadFlags()

	//
	// Check for version request
	//
	if flags.ShowVersionOnly {
		log.Print(version.GetVersion())
		return
	}

	//
	// Load env.
	//
	if flags.EnvFile != "" {
		err := env.LoadEnvFile(flags.EnvFile)
		if err != nil {
			log.Fatal(errors.Wrap(err, "failed to load environment configurations"))
		}
	}

	//
	// Select service
	//
	registryContainer := registry.NewRegistryContainer()

	registryContainer.Add(linkRESTGateway.ServiceName, linkRESTGateway.FactoryMethod)
	registryContainer.Add(linkService.ServiceName, linkService.FactoryMethod)

	serviceFactory, err := registryContainer.Get(flags.Kind)
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to initialize service"))
	}

	//
	// Create service
	//
	service, err := serviceFactory()
	if err != nil {
		log.Fatal(errors.Wrapf(err, "failed to initialize application [%s]", flags.Kind))
	}

	//
	// Run till the death comes
	//
	log.Printf("%s, service [%s] is started", version.GetVersion(), flags.Kind)
	log.Fatal(service.Serve())
}
