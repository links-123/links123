package internal

import (
	"context"

	"github.com/ic2hrmk/links123/app/services/link/errors"
	"github.com/ic2hrmk/links123/app/services/link/pb/link"
	"github.com/ic2hrmk/links123/app/services/link/persistence/model"
)

//
// Saves links to persistence via bulk
//
func (rcv *linkDomainService) SaveLinksBulk(
	ctx context.Context, in *link.SaveLinksBulkRequest,
) (*link.SaveLinksBulkResponse, error) {
	//
	// Validation
	//
	if err := in.Validate(); err != nil {
		return nil, errors.InvalidRequest(err)
	}

	//
	// Discard empty bulk inserts
	//
	if len(in.GetItems()) == 0 {
		return &link.SaveLinksBulkResponse{}, nil
	}

	//
	// Request handing
	//
	records := make([]*model.Link, len(in.GetItems()))

	for i, linkDetails := range in.GetItems() {
		records[i] = &model.Link{
			LinkID:  linkDetails.GetLinkID(),
			Name:    linkDetails.GetName(),
			Address: linkDetails.GetAddress(),
		}
	}

	if err := rcv.linkRepository.SaveBulk(records); err != nil {
		return nil, errors.Internal(err)
	}

	//
	// Assemble response
	//
	return &link.SaveLinksBulkResponse{}, nil
}
