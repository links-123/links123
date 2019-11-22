package main

import (
	"fmt"
	"log"

	linkRESTGateway "github.com/ic2hrmk/links123/app/gateways/link/rest"
	linkService "github.com/ic2hrmk/links123/app/services/link"

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
			log.Fatal(err)
		}
	}

	//
	// Select service
	//
	reg := registry.NewRegistryContainer()

	reg.Add(linkRESTGateway.ServiceName, linkRESTGateway.FactoryMethod)
	reg.Add(linkService.ServiceName, linkService.FactoryMethod)

	serviceFactory, err := reg.Get(flags.Kind)
	if err != nil {
		log.Fatal(err)
	}

	//
	// Create service
	//
	service, err := serviceFactory()
	if err != nil {
		log.Fatal(err)
	}

	//
	// Run till the death comes
	//
	log.Printf("[%s] started serving on '%s'", flags.Kind, flags.Address)
	log.Fatal(service.Serve(flags.Address))
}
