package link_rest_gtw

import (
	"github.com/micro/go-micro/v2/client"
	"github.com/micro/go-micro/v2/config"
	"github.com/pkg/errors"

	linkPb "github.com/links-123/links123/app/services/link/pb/link"

	"github.com/links-123/links123/app"
	"github.com/links-123/links123/app/gateways/link/link-rest-gtw/internal"
	"github.com/links-123/links123/app/services/link"
	"github.com/links-123/links123/shared/microservice"
)

const ServiceName = "links123.link.restgtw"

//
// ServiceName constructor
//
func FactoryMethod(configSource config.Config) (app.MicroService, error) {
	//
	// Resolve configurations
	//	- clients's configurations
	//  - gateway's configurations
	//
	configurations, err := ResolveConfigurations(configSource)
	if err != nil {
		return nil, errors.Wrap(err, "failed to resolve configurations")
	}

	//
	// Init. clients
	//
	linkServiceClient, err := initLinkServiceClient()
	if err != nil {
		return nil, errors.Errorf("failed to init. link service client, %s", err)
	}

	//
	// Init. the whole microservice
	//

	// Business logic processor
	linkRESTGateway := internal.NewLinkDomainService(
		linkServiceClient,
	)

	// Request serving service
	microService, err := microservice.NewWebService(
		ServiceName,
		configurations.serveAddress,
		linkRESTGateway.GetHttpHandler(),
	)

	if err != nil {
		return nil, errors.Wrap(err, "failed to inti. web service")
	}

	return microService, nil
}

func initLinkServiceClient() (linkPb.LinkDomainService, error) {
	return linkPb.NewLinkDomainService(
		link.ServiceName,
		client.DefaultClient,
	), nil
}
