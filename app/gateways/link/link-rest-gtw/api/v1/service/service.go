package service

import (
	"net/http"

	"github.com/sirupsen/logrus"

	"github.com/links-123/links123/app"
	linkPb "github.com/links-123/links123/app/services/link/pb/link"

	"github.com/emicklei/go-restful"

	"github.com/links-123/links123/app/gateways/link/link-rest-gtw/api/v1/representation"
)

type linkRESTService struct {
	log               *logrus.Logger
	webService        *restful.WebService
	linkServiceClient linkPb.LinkDomainService
}

func NewLinkRESTService(
	logger *logrus.Logger,
	linkServiceClient linkPb.LinkDomainService,
) app.RESTAPIVersionedService {
	rcv := &linkRESTService{
		log:               logger,
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

	ws.Route(ws.POST("/links").
		To(rcv.createLink).
		Operation("createLink").
		Reads(representation.LinkCreateRequest{}).
		Writes(representation.LinkEntityResponse{}).
		Returns(http.StatusOK, http.StatusText(http.StatusOK), representation.LinkListResponse{}).
		Returns(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), representation.ErrorResponse{}))

	ws.Route(ws.DELETE("/links").
		To(rcv.deleteLink).
		Operation("deleteLink").
		Reads(representation.LinkDeleteRequest{}).
		Returns(http.StatusNoContent, http.StatusText(http.StatusOK), nil).
		Returns(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), representation.ErrorResponse{}))
}
