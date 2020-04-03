package service

import (
	linkPb "github.com/links-123/links123/app/services/link/pb/link"
	"net/http"

	"github.com/emicklei/go-restful"

	"github.com/links-123/links123/app/gateways/link/link-rest-gtw/api"
	"github.com/links-123/links123/app/gateways/link/link-rest-gtw/api/v1/representation"
)

type linkRESTService struct {
	webService        *restful.WebService
	linkServiceClient linkPb.LinkDomainServiceClient
}

func NewLinkRESTService(
	linkServiceClient linkPb.LinkDomainServiceClient,
) api.RESTAPIVersionedService {
	rcv := &linkRESTService{
		webService:        &restful.WebService{},
		linkServiceClient: linkServiceClient,
	}

	rcv.initWebService()

	return rcv
}

func (rcv *linkRESTService) ProvideWebService() *restful.WebService {
	return rcv.webService
}

func (rcv *linkRESTService) initWebService() {
	ws := rcv.webService

	ws.Path("/api/v2").
		ApiVersion("v2").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)

	ws.Route(ws.GET("/links").
		To(rcv.notImplemented).
		Returns(http.StatusBadRequest, http.StatusText(http.StatusInternalServerError), representation.ErrorResponse{}))
}