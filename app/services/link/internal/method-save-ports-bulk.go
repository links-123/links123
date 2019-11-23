package internal

import (
	"context"

	"github.com/pkg/errors"

	"github.com/ic2hrmk/links123/app/services/link/pb/link"
	"github.com/ic2hrmk/links123/app/services/link/persistence/model"
)

//
// Saves links to persistence via bulk
//
func (rcv *linkDomainService) SaveLinksBulk(
	ctx context.Context, in *link.SaveLinksBulkRequest,
) (*link.SaveLinksBulkResponse, error) {
	var err error

	//
	// Validation
	//
	if in == nil {
		return nil, rcv.wrapInvalidRequest(errors.New("request is empty"))
	}

	if err = in.Validate(); err != nil {
		return nil, rcv.wrapInvalidRequest(err)
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

		if linkDetails.GetLinkID() == "" {
			records[i].LinkID = rcv.generateLinkID()
		}
	}

	if err = rcv.linkRepository.SaveBulk(records); err != nil {
		return nil, rcv.wrapInternalError(errors.Wrap(err, "unable to persist link"))
	}

	//
	// Assemble response
	//
	responseEntities := make([]*link.LinkEntity, 0, len(records))
	for i := range records {
		responseEntities = append(responseEntities, &link.LinkEntity{
			LinkID:  records[i].LinkID,
			Name:    records[i].Name,
			Address: records[i].Address,
		})
	}

	return &link.SaveLinksBulkResponse{
		Items: responseEntities,
	}, nil
}
