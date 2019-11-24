package service

import (
	"net/http"

	linkPb "github.com/ic2hrmk/links123/app/services/link/pb/link"

	"github.com/emicklei/go-restful"

	"github.com/ic2hrmk/links123/app/gateways/link/api/rest/internal/api-version"
	"github.com/ic2hrmk/links123/app/gateways/link/api/rest/internal/api-version/v1/representation"
)

type linkRESTService struct {
	webService        *restful.WebService
	linkServiceClient linkPb.LinkDomainServiceClient
}

func NewLinkRESTService(
	linkServiceClient linkPb.LinkDomainServiceClient,
) api_version.APIVersionService {
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

	ws.Path("/api/v1").
		ApiVersion("v1").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)

	ws.Route(ws.GET("/links").
		To(rcv.getLinks).
		Operation("getLinks").
		Param(ws.QueryParameter(representation.LimitParameterName, "Limit").DataType("integer")).
		Param(ws.QueryParameter(representation.OffsetParameterName, "Offset").DataType("integer")).
		Writes(representation.LinkListResponse{}).
		Returns(http.StatusOK, http.StatusText(http.StatusOK), representation.LinkListResponse{}).
		Returns(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), representation.ErrorResponse{}))
}
