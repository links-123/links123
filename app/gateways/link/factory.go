package link

import (
	"fmt"

	linkPb "github.com/ic2hrmk/links123/app/services/link/pb/link"

	"github.com/ic2hrmk/links123/app"
	"github.com/ic2hrmk/links123/app/gateways/link/config"
	"github.com/ic2hrmk/links123/app/gateways/link/internal"
	"google.golang.org/grpc"
)

const ServiceName = "link-gtw"

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
	gatewayConfigurationBuilder := internal.NewLinkGatewayConfigBuilder()

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
