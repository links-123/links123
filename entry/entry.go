package main

import (
	"fmt"
	"log"

	linkRESTGateway "github.com/ic2hrmk/links123/app/gateways/link/api/rest"
	linkService "github.com/ic2hrmk/links123/app/services/link"

	"github.com/pkg/errors"

	"github.com/ic2hrmk/links123/registry"
	"github.com/ic2hrmk/links123/shared/cmd"
	"github.com/ic2hrmk/links123/shared/env"
	"github.com/ic2hrmk/links123/shared/version"
)

const ApplicationName = "links123"

func printVersion() {
	fmt.Printf("%s %s\n", ApplicationName, version.GetVersion())
}

func main() {
	printVersion()

	//
	// Load startup flags
	//
	flags := cmd.LoadFlags()

	//
	// Check for version request
	//
	if flags.ShowVersionOnly {
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
	log.Printf("service [%s] is started", flags.Kind)
	log.Fatal(service.Serve())
}
