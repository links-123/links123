package internal

import (
	"context"

	linkPb "github.com/ic2hrmk/links123/app/services/link/pb/link"

	"github.com/emicklei/go-restful"
	"github.com/ic2hrmk/links123/app/gateways/link/rest/errors"
	"github.com/ic2hrmk/links123/app/gateways/link/rest/representation"
	"github.com/ic2hrmk/links123/shared/gateway/helpers"
)

func (rcv *linksGateway) getLinks(
	request *restful.Request,
	response *restful.Response,
) {
	var limit, offset uint64
	var err error

	//
	// Get limit and offset
	//
	if limit, err = getLimitParameter(request); err != nil {
		helpers.ResponseWithBadRequest(response, err, errors.ErrInvalidLimitParameter)
		return
	}

	if offset, err = getOffsetParameter(request); err != nil {
		helpers.ResponseWithBadRequest(response, err, errors.ErrInvalidOffsetParameter)
		return
	}

	//
	// Request information
	//
	linksDetails, err := rcv.linkServiceClient.FindAllLinks(context.Background(),
		&linkPb.FindAllLinksRequest{
			Limit:  limit,
			Offset: offset,
		})

	if err != nil {
		helpers.ResponseWithInternalError(response, err)
		return
	}

	//
	// Assemble response
	//
	out := &representation.LinkListResponse{
		Items: make([]*representation.LinkEntityResponse, len(linksDetails.GetItems())),
		Found: len(linksDetails.GetItems()),
	}

	for i, link := range linksDetails.GetItems() {
		out.Items[i] = &representation.LinkEntityResponse{
			LinkID:  link.GetLinkID(),
			Name:    link.GetName(),
			Address: link.GetAddress(),
		}
	}

	helpers.ResponseWithOK(response, out)
}
