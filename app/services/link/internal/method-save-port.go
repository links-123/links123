package internal

import (
	"context"

	"github.com/pkg/errors"

	"github.com/ic2hrmk/links123/app/services/link/pb/link"
	"github.com/ic2hrmk/links123/app/services/link/persistence/model"
)

//
// Saves link to persistence
//
func (rcv *linkDomainService) SaveLink(
	ctx context.Context, in *link.SaveLinkRequest,
) (*link.SaveLinkResponse, error) {
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
	// Request handing
	//
	snapshot := &model.Link{
		LinkID:  in.GetLink().GetLinkID(),
		Name:    in.GetLink().GetName(),
		Address: in.GetLink().GetAddress(),
	}

	if in.Link.LinkID == "" {
		snapshot.LinkID = rcv.generateLinkID()
	}

	if _, err = rcv.linkRepository.Save(snapshot); err != nil {
		return nil, rcv.wrapInternalError(errors.Wrap(err, "unable to persist link"))
	}

	//
	// Assemble response
	//
	return &link.SaveLinkResponse{
		Link: &link.LinkEntity{
			LinkID:  snapshot.LinkID,
			Name:    snapshot.Name,
			Address: snapshot.Address,
		},
	}, nil
}
