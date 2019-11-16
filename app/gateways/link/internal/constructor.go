package internal

import (
	linkPb "github.com/ic2hrmk/links123/app/services/link/pb/link"

	"github.com/emicklei/go-restful"
	"github.com/go-ozzo/ozzo-validation"
	"github.com/ic2hrmk/links123/app"
)

//
// Link Gateway micro-service
//
type linksGateway struct {
	webContainer *restful.Container
	config       *linkDomainGatewayConfig

	//
	// Clients to back-services
	//
	linkServiceClient linkPb.LinkDomainServiceClient
}

func NewLinkDomainService(
	config *linkDomainGatewayConfig,
	linkServiceClient linkPb.LinkDomainServiceClient,
) app.MicroService {
	service := &linksGateway{
		config:            config,
		webContainer:      restful.NewContainer(),
		linkServiceClient: linkServiceClient,
	}

	service.init()

	return service
}

//
// =============== Additional configuration section ================
//

//
// Configuration container
//
type linkDomainGatewayConfig struct {
}

//
// Configuration container builder
//
type linkDomainGatewayConfigBuilder struct {
}

func NewLinkGatewayConfigBuilder() *linkDomainGatewayConfigBuilder {
	return &linkDomainGatewayConfigBuilder{}
}

//
// Builds gateway configuration object. It's the only way to initialize it's settings
//
func (rcv *linkDomainGatewayConfigBuilder) Build() (*linkDomainGatewayConfig, error) {
	if err := rcv.Validate(); err != nil {
		return nil, err
	}

	return &linkDomainGatewayConfig{}, nil
}

//
// Validates acquired settings
//
func (rcv *linkDomainGatewayConfigBuilder) Validate() error {
	return validation.ValidateStruct(rcv,
	)
}
