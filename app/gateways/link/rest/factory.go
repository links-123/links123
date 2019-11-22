package rest

import (
	"fmt"

	linkPb "github.com/ic2hrmk/links123/app/services/link/pb/link"

	"google.golang.org/grpc"

	"github.com/ic2hrmk/links123/app"
	"github.com/ic2hrmk/links123/app/gateways/link/rest/config"
	"github.com/ic2hrmk/links123/app/gateways/link/rest/internal"
)

const ServiceName = "link-rest-gtw"

//
// Service constructor
//
func FactoryMethod() (app.MicroService, error) {
	//
	// Resolve configurations
	//	- clients's configurations
	//  - gateway's configurations
	//
	configurations, err := resolveConfigurations()
	if err != nil {
		return nil, err
	}

	//
	// Init. gateway configurations
	//
	gatewayConfigurationBuilder := internal.NewLinkRESTGatewayConfigBuilder()

	// ... set any other configurations here

	gatewayConfiguration, err := gatewayConfigurationBuilder.Build()
	if err != nil {
		return nil, err
	}

	//
	// Init. clients
	//
	linkServiceClient, err := initLinkServiceClient(configurations.LinkDomainServiceAddress)
	if err != nil {
		return nil, fmt.Errorf("failed to init. link service client, %s", err)
	}

	return internal.NewLinkDomainService(
		gatewayConfiguration,
		linkServiceClient,
	), nil
}

func resolveConfigurations() (*config.ConfigurationContainer, error) {
	return config.ResolveConfigurations()
}

func initLinkServiceClient(address string) (linkPb.LinkDomainServiceClient, error) {
	var conn *grpc.ClientConn

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("did not connect: %s", err)
	}

	return linkPb.NewLinkDomainServiceClient(conn), nil
}
