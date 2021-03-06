package service

import (
	"context"

	linkPb "github.com/links-123/links123/app/services/link/pb/link"

	"github.com/emicklei/go-restful"

	"github.com/links-123/links123/app/gateways/link/link-rest-gtw/api/v1/representation"
	"github.com/links-123/links123/shared/gateway/helpers"
)

func (rcv *linkRESTService) createLink(request *restful.Request, response *restful.Response) {
	var (
		createRequest = new(representation.LinkCreateRequest)
		err           error
	)

	if err = request.ReadEntity(createRequest); err != nil {
		helpers.RespondWithBadRequest(response, rcv.log, err, representation.ErrIncorrectRequest)
		return
	}

	if err = createRequest.Validate(); err != nil {
		helpers.RespondWithBadRequest(response, rcv.log, err, representation.ErrInvalidCreateRequest)
		return
	}

	//
	// Request information
	//
	linksDetails, err := rcv.linkServiceClient.CreateLink(context.Background(),
		&linkPb.CreateLinkRequest{
			Link: &linkPb.LinkEntity{
				UserID:  "demo",
				Name:    createRequest.Name,
				Address: createRequest.Address,
			},
		})

	if err != nil {
		helpers.RespondWithInternalError(response, rcv.log, err)
		return
	}

	//
	// Assemble response
	//
	out := &representation.LinkEntityResponse{
		LinkID:  linksDetails.GetLink().GetLinkID(),
		Name:    linksDetails.GetLink().GetName(),
		Address: linksDetails.GetLink().GetAddress(),
	}

	helpers.RespondWithOK(response, rcv.log, out)
}
