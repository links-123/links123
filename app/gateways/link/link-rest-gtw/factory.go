package link_rest_gtw

import (
	microClient "github.com/micro/go-micro/v2/client"
	microConfig "github.com/micro/go-micro/v2/config"
	microLogger "github.com/micro/go-micro/v2/logger"

	linkPb "github.com/links-123/links123/app/services/link/pb/link"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	"github.com/links-123/links123/app"
	"github.com/links-123/links123/app/gateways/link/link-rest-gtw/internal"
	"github.com/links-123/links123/app/services/link"
	"github.com/links-123/links123/shared/logging/local/micrologrus"
	"github.com/links-123/links123/shared/microservice"
)

const ServiceName = "links123.link.restgtw"

//
// ServiceName constructor
//
func FactoryMethod(
	configSource microConfig.Config,
	logger *logrus.Logger,
) (app.MicroService, error) {
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
	// Patch framework's logger to use Logrus instead
	// NOTE: Known issue when logger patch is not applied - https://github.com/micro/go-micro/issues/1514
	microLogger.DefaultLogger, err = micrologrus.NewMicroLogrus(logger)
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
	linkRESTGateway := internal.NewLinkGatewayService(
		logger,
		linkServiceClient,
	)

	// Request serving service
	microService, err := microservice.NewWebService(
		ServiceName,
		configurations.instanceID,
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
		microClient.DefaultClient,
	), nil
}
