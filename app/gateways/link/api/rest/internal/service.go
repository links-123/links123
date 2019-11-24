package internal

import (
	"log"
	"net/http"

	serviceV1 "github.com/ic2hrmk/links123/app/gateways/link/api/rest/internal/api-version/v1/service"
	serviceV2 "github.com/ic2hrmk/links123/app/gateways/link/api/rest/internal/api-version/v2/service"
	linkPb "github.com/ic2hrmk/links123/app/services/link/pb/link"

	"github.com/emicklei/go-restful"
	"github.com/go-ozzo/ozzo-validation"
	"github.com/ic2hrmk/links123/app"
	"github.com/pkg/errors"
)

//
// Link Gateway micro-service
//
type linkGateway struct {
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
	service := &linkGateway{
		configurations:    config,
		webContainer:      restful.NewContainer(),
		linkServiceClient: linkServiceClient,
	}

	service.init()

	return service
}

func (rcv *linkGateway) Serve() error {
	log.Printf("service is available via [http] at address [%s]",
		rcv.configurations.serveAddress)
	return http.ListenAndServe(rcv.configurations.serveAddress, rcv.webContainer)
}

func (rcv *linkGateway) init() {
	var (
		linkRESTv1Service = serviceV1.NewLinkRESTService(rcv.linkServiceClient)
		linkRESTv2Service = serviceV2.NewLinkRESTService(rcv.linkServiceClient)
	)

	rcv.webContainer.Add(linkRESTv1Service.ProvideWebService())
	rcv.webContainer.Add(linkRESTv2Service.ProvideWebService())
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
