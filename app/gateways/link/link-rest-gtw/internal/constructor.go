package internal

import (
	"net/http"

	"github.com/sirupsen/logrus"

	serviceV1 "github.com/links-123/links123/app/gateways/link/link-rest-gtw/api/v1/service"
	serviceV2 "github.com/links-123/links123/app/gateways/link/link-rest-gtw/api/v2/service"
	linkPb "github.com/links-123/links123/app/services/link/pb/link"

	"github.com/emicklei/go-restful"
)

//
// Link Gateway
//
type linkGateway struct {
	log               *logrus.Logger
	webContainer      *restful.Container
	linkServiceClient linkPb.LinkDomainService
}

func NewLinkGatewayService(
	logger *logrus.Logger,
	linkServiceClient linkPb.LinkDomainService,
) *linkGateway {
	service := &linkGateway{
		log:               logger,
		webContainer:      restful.NewContainer(),
		linkServiceClient: linkServiceClient,
	}

	var (
		linkRESTv1Service = serviceV1.NewLinkRESTService(logger, service.linkServiceClient)
		linkRESTv2Service = serviceV2.NewLinkRESTService(logger, service.linkServiceClient)
	)

	service.webContainer.Add(linkRESTv1Service.ProvideWebService())
	service.webContainer.Add(linkRESTv2Service.ProvideWebService())

	return service
}

func (rcv *linkGateway) GetHttpHandler() http.Handler {
	return rcv.webContainer
}
