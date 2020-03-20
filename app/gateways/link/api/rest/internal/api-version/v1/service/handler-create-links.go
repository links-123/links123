package service

import (
	"context"

	linkPb "github.com/ic2hrmk/links123/app/services/link/pb/link"

	"github.com/emicklei/go-restful"

	"github.com/ic2hrmk/links123/app/gateways/link/api/rest/internal/api-version/v1/representation"
	"github.com/ic2hrmk/links123/shared/gateway/helpers"
)

func (rcv *linkRESTService) createLink(request *restful.Request, response *restful.Response) {
	var (
		createRequest = new(representation.LinkCreateRequest)
		err           error
	)

	if err = request.ReadEntity(createRequest); err != nil {
		helpers.ResponseWithBadRequest(response, err, representation.ErrIncorrectRequest)
		return
	}

	if err = createRequest.Validate(); err != nil {
		helpers.ResponseWithBadRequest(response, err, representation.ErrInvalidCreateRequest)
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
		helpers.ResponseWithInternalError(response, err)
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

	helpers.ResponseWithOK(response, out)
}
