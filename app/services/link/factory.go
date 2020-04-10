package link

import (
	"github.com/globalsign/mgo"
	"github.com/micro/go-micro/v2/config"
	"github.com/pkg/errors"

	linkPb "github.com/links-123/links123/app/services/link/pb/link"
	"github.com/links-123/links123/app/services/link/persistence/repository/mongo"
	sharedMongo "github.com/links-123/links123/shared/persistence/mongo"

	"github.com/links-123/links123/app"
	"github.com/links-123/links123/app/services/link/internal"
	"github.com/links-123/links123/shared/microservice"
)

const (
	ServiceName = "links123.link.srv"
	mongoDBName = "link-srv"
)

//
// ServiceName constructor
//
func FactoryMethod(configSource config.Config) (app.MicroService, error) {
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
		linkRepository,
	)

	// Request serving service
	microService := microservice.NewBackendService(ServiceName)

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
