package service

import (
	"net/http"

	linkPb "github.com/ic2hrmk/links123/app/services/link/pb/link"

	"github.com/emicklei/go-restful"

	"github.com/ic2hrmk/links123/app/gateways/link/api/rest/internal/api-version"
	"github.com/ic2hrmk/links123/app/gateways/link/api/rest/internal/api-version/v1/representation"
)

type linkRESTv2Service struct {
	webService        *restful.WebService
	linkServiceClient linkPb.LinkDomainServiceClient
}

func NewLinkRESTService(
	linkServiceClient linkPb.LinkDomainServiceClient,
) api_version.APIVersionService {
	rcv := &linkRESTv2Service{
		webService:        &restful.WebService{},
		linkServiceClient: linkServiceClient,
	}

	rcv.initWebService()

	return rcv
}

func (rcv *linkRESTv2Service) ProvideWebService() *restful.WebService {
	return rcv.webService
}

func (rcv *linkRESTv2Service) initWebService() {
	ws := rcv.webService

	ws.Path("/api/v2").
		ApiVersion("v2").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)

	ws.Route(ws.GET("").
		To(rcv.notImplemented).
		Returns(http.StatusBadRequest, http.StatusText(http.StatusInternalServerError), representation.ErrorResponse{}))
}
