package link_rest_gtw

import (
	linkPb "github.com/links-123/links123/app/services/link/pb/link"

	"github.com/pkg/errors"
	"google.golang.org/grpc"

	"github.com/links-123/links123/app"
	"github.com/links-123/links123/app/gateways/link/link-rest-gtw/config"
	"github.com/links-123/links123/app/gateways/link/link-rest-gtw/internal"
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

	gatewayConfigurationBuilder.
		SetLinkServiceAddress(configurations.LinkDomainServiceAddress).
		SetServeAddress(configurations.ServeAddress)

	gatewayConfiguration, err := gatewayConfigurationBuilder.Build()
	if err != nil {
		return nil, err
	}

	//
	// Init. clients
	//
	linkServiceClient, err := initLinkServiceClient(configurations.LinkDomainServiceAddress)
	if err != nil {
		return nil, errors.Errorf("failed to init. link service client, %s", err)
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
		return nil, errors.Errorf("unable to connect: %s", err)
	}

	return linkPb.NewLinkDomainServiceClient(conn), nil
}
