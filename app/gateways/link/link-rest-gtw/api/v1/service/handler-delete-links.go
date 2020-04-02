package service

import (
	"context"

	linkPb "github.com/links-123/links123/app/services/link/pb/link"

	"github.com/emicklei/go-restful"

	"github.com/links-123/links123/app/gateways/link/link-rest-gtw/api/v1/representation"
	"github.com/links-123/links123/shared/gateway/helpers"
)

func (rcv *linkRESTService) deleteLink(request *restful.Request, response *restful.Response) {
	var (
		deleteRequest = new(representation.LinkDeleteRequest)
		err           error
	)

	if err = request.ReadEntity(deleteRequest); err != nil {
		helpers.ResponseWithBadRequest(response, err, representation.ErrIncorrectRequest)
		return
	}

	if err = deleteRequest.Validate(); err != nil {
		helpers.ResponseWithBadRequest(response, err, representation.ErrInvalidDeleteRequest)
		return
	}

	//
	// Request information
	//
	_, err = rcv.linkServiceClient.DeleteLink(context.Background(),
		&linkPb.DeleteLinkRequest{
			UserID: "demo",
			LinkID: deleteRequest.LinkID,
		})
	if err != nil {
		helpers.ResponseWithInternalError(response, err)
		return
	}

	//
	// Assemble response
	//
	helpers.ResponseWithNoContent(response)
}
