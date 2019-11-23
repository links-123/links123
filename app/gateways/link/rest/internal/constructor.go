package internal

import (
	linkPb "github.com/ic2hrmk/links123/app/services/link/pb/link"

	"github.com/ic2hrmk/links123/app"

	"github.com/emicklei/go-restful"
	"github.com/go-ozzo/ozzo-validation"
	"github.com/pkg/errors"
)

//
// Link Gateway micro-service
//
type linksGateway struct {
	webContainer   *restful.Container
	configurations *linkDomainGatewayConfigurations

	//
	// Clients to back-services
	//
	linkServiceClient linkPb.LinkDomainServiceClient
}

func NewLinkDomainService(
	config *linkDomainGatewayConfigurations,
	linkServiceClient linkPb.LinkDomainServiceClient,
) app.MicroService {
	service := &linksGateway{
		configurations:    config,
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
type linkDomainGatewayConfigurations struct {
	linkServiceAddress string
	serveAddress       string
}

//
// Configuration container builder
//
type linkDomainGatewayConfigBuilder struct {
	linkServiceAddress string
	serveAddress       string
}

func NewLinkRESTGatewayConfigBuilder() *linkDomainGatewayConfigBuilder {
	return &linkDomainGatewayConfigBuilder{}
}

func (rcv *linkDomainGatewayConfigBuilder) SetServeAddress(address string) *linkDomainGatewayConfigBuilder {
	rcv.serveAddress = address
	return rcv
}

func (rcv *linkDomainGatewayConfigBuilder) SetLinkServiceAddress(address string) *linkDomainGatewayConfigBuilder {
	rcv.linkServiceAddress = address
	return rcv
}

//
// Builds gateway configuration object. It's the only way to initialize it's settings
//
func (rcv *linkDomainGatewayConfigBuilder) Build() (*linkDomainGatewayConfigurations, error) {
	if err := rcv.Validate(); err != nil {
		return nil, errors.Wrap(err, "configuration is invalid")
	}

	return &linkDomainGatewayConfigurations{
		linkServiceAddress: rcv.linkServiceAddress,
		serveAddress:       rcv.serveAddress,
	}, nil
}

//
// Validates acquired settings
//
func (rcv *linkDomainGatewayConfigBuilder) Validate() error {
	return validation.ValidateStruct(rcv,
		validation.Field(&rcv.linkServiceAddress, validation.Required),
		validation.Field(&rcv.serveAddress, validation.Required),
	)
}
