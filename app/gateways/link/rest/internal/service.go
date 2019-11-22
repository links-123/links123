package internal

import (
	"net/http"

	"github.com/emicklei/go-restful"
	"github.com/ic2hrmk/links123/app/gateways/link/rest/representation"
	"github.com/ic2hrmk/links123/shared/gateway/filters"
)

func (rcv *linksGateway) init() {
	ws := &restful.WebService{}

	ws.Path("/api").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)

	ws.Route(ws.GET("/links").
		To(rcv.getLinks).
		Operation("getLinks").
		Param(ws.QueryParameter(limitParameterName, "Limit").DataType("integer")).
		Param(ws.QueryParameter(offsetParameterName, "Offset").DataType("integer")).
		Writes(representation.LinkListResponse{}).
		Returns(http.StatusOK, http.StatusText(http.StatusOK), representation.LinkListResponse{}).
		Returns(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), representation.ErrorResponse{}))

	ws.Filter(filters.LogRequest)

	rcv.webContainer.Add(ws)
}

func (rcv *linksGateway) Serve(address string) error {
	return rcv.serve(address)
}

func (rcv *linksGateway) serve(address string) error {
	return http.ListenAndServe(address, rcv.webContainer)
}
