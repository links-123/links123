package service

import (
	"context"

	linkPb "github.com/links-123/links123/app/services/link/pb/link"

	"github.com/emicklei/go-restful"

	"github.com/links-123/links123/app/gateways/link/link-rest-gtw/api/v1/representation"
	"github.com/links-123/links123/shared/gateway/helpers"
)

func (rcv *linkRESTService) getLinks(request *restful.Request, response *restful.Response) {
	var limit, offset uint64
	var err error

	//
	// Get limit and offset
	//
	if limit, err = getLimitParameter(request); err != nil {
		helpers.RespondWithBadRequest(response, rcv.log, err, representation.ErrInvalidLimitParameter)
		return
	}

	if offset, err = getOffsetParameter(request); err != nil {
		helpers.RespondWithBadRequest(response, rcv.log, err, representation.ErrInvalidOffsetParameter)
		return
	}

	//
	// Request information
	//
	linksDetails, err := rcv.linkServiceClient.FindLinks(context.Background(),
		&linkPb.FindLinksRequest{
			UserID: "demo",
			Limit:  limit,
			Offset: offset,
		})

	if err != nil {
		helpers.RespondWithInternalError(response, rcv.log, err)
		return
	}

	//
	// Assemble response
	//
	out := &representation.LinkListResponse{
		Items:            make([]*representation.LinkEntityResponse, len(linksDetails.GetItems())),
		TotalLinksNumber: linksDetails.GetTotalLinksNumber(),
	}

	for i, link := range linksDetails.GetItems() {
		out.Items[i] = &representation.LinkEntityResponse{
			LinkID:  link.GetLinkID(),
			Name:    link.GetName(),
			Address: link.GetAddress(),
		}
	}

	helpers.RespondWithOK(response, rcv.log, out)
}
