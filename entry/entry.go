package main

import (
	"log"

	linkGateway "github.com/ic2hrmk/links123/app/gateways/link"
	linkService "github.com/ic2hrmk/links123/app/services/link"

	"github.com/ic2hrmk/links123/registry"
	"github.com/ic2hrmk/links123/shared/cmd"
	"github.com/ic2hrmk/links123/shared/env"
)

//go:generate go run entry.go --kind=link-gtw --address=:8080 --env=docker-compose.env
//go:generate go run entry.go --kind=link-srv --address=:10001 --env=.env

func main() {
	//
	// Load startup flags
	//
	flags := cmd.LoadFlags()

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

	reg.Add(linkGateway.ServiceName, linkGateway.FactoryMethod)
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
