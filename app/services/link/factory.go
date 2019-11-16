package link

import (
	sharedMongo "github.com/ic2hrmk/ship_links/shared/persistence/mongo"

	"github.com/globalsign/mgo"
	"github.com/ic2hrmk/ship_links/app"
	"github.com/ic2hrmk/ship_links/app/services/link/config"
	"github.com/ic2hrmk/ship_links/app/services/link/internal"
	"github.com/ic2hrmk/ship_links/app/services/link/persistence/repository/mongo"
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
	// Init. persistence
	//
	mongoDB, err := initMongoDB(configurations.MongoURL, ServiceName)
	if err != nil {
		return nil, err
	}

	// Init. link repository
	linkRepository := mongo.NewLinkRepository(mongoDB)

	return internal.NewLinkDomainService(
		linkRepository,
	), nil
}

func resolveConfigurations() (*config.ConfigurationContainer, error) {
	return config.ResolveConfigurations()
}

func initMongoDB(mongoURL, dbName string) (*mgo.Database, error) {
	return sharedMongo.InitConnection(mongoURL, dbName)
}
