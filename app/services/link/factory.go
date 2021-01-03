package link

import (
	"github.com/globalsign/mgo"
	"github.com/micro/go-micro/v2/config"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	linkPb "github.com/links-123/links123/app/services/link/pb/link"
	sharedMongo "github.com/links-123/links123/shared/persistence/mongo"

	microLogger "github.com/micro/go-micro/v2/logger"

	"github.com/links-123/links123/app"
	"github.com/links-123/links123/app/services/link/internal"
	"github.com/links-123/links123/app/services/link/persistence/repository/mongo"
	"github.com/links-123/links123/shared/logging/local/micrologrus"
	"github.com/links-123/links123/shared/microservice"
)

const (
	ServiceName = "links123.link.srv"
	mongoDBName = "link-srv"
)

//
// ServiceName constructor
//
func FactoryMethod(configSource config.Config, logger *logrus.Logger) (app.MicroService, error) {
	//
	// Resolve configurations
	//	- service's configurations
	//  - DB's configurations
	//
	configurations, err := extractConfigurations(configSource)
	if err != nil {
		return nil, errors.Wrap(err, "failed to resolve configurations")
	}

	//
	// Init. persistence
	//
	mongoDB, err := initMongoDB(configurations, mongoDBName)
	if err != nil {
		return nil, err
	}

	linkRepository := mongo.NewLinkRepository(mongoDB)

	//
	// Init. the whole microservice
	//

	// Business logic processor
	linkDomainService := internal.NewLinkDomainService(
		logger,
		linkRepository,
	)

	//
	// Patch framework's logger to use Logrus instead
	// NOTE: Known issue when logger patch is not applied - https://github.com/micro/go-micro/issues/1514
	microLogger.DefaultLogger, err = micrologrus.NewMicroLogrus(logger)
	if err != nil {
		return nil, errors.Wrap(err, "failed to resolve configurations")
	}

	// Request serving service
	microService := microservice.NewBackendService(
		configurations.instanceID,
		ServiceName,
	)

	// ServiceName linking
	err = linkPb.RegisterLinkDomainServiceHandler(microService.Server(), linkDomainService)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to register service [name=%s]", ServiceName)
	}

	return microService, nil
}

func initMongoDB(config *serviceConfiguration, dbName string) (*mgo.Database, error) {
	return sharedMongo.InitConnection(config.mongoURL, dbName)
}
