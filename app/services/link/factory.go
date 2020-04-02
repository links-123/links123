package link

import (
	sharedMongo "github.com/links-123/links123/shared/persistence/mongo"
	"github.com/pkg/errors"

	"github.com/globalsign/mgo"
	"github.com/links-123/links123/app"
	"github.com/links-123/links123/app/services/link/config"
	"github.com/links-123/links123/app/services/link/internal"
	"github.com/links-123/links123/app/services/link/persistence/repository/mongo"
)

const ServiceName = "link-srv"

//
// Service constructor
//
func FactoryMethod() (app.MicroService, error) {
	//
	// Resolve configurations
	//	- service's configurations
	//
	configurations, err := resolveConfigurations()
	if err != nil {
		return nil, err
	}

	//
	// Init. service configurations
	//
	serviceConfigurationBuilder := internal.NewLinkDomainServiceConfigBuilder()

	serviceConfigurationBuilder.
		SetMongoURL(configurations.MongoURL).
		SetServeAddress(configurations.ServeAddress)

	serviceConfiguration, err := serviceConfigurationBuilder.Build()
	if err != nil {
		return nil, errors.Wrap(err, "failed to compose service configurations")
	}

	//
	// Init. persistence
	//
	mongoDB, err := initMongoDB(configurations.MongoURL, ServiceName)
	if err != nil {
		return nil, err
	}

	linkRepository := mongo.NewLinkRepository(mongoDB)

	return internal.NewLinkDomainService(
		serviceConfiguration,
		linkRepository,
	), nil
}

func resolveConfigurations() (*config.ConfigurationContainer, error) {
	return config.ResolveConfigurations()
}

func initMongoDB(mongoURL, dbName string) (*mgo.Database, error) {
	return sharedMongo.InitConnection(mongoURL, dbName)
}
