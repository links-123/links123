package internal

import (
	"context"

	"github.com/pkg/errors"

	"github.com/links-123/links123/app/services/link/pb/link"
)

//
// Creates link to persistence
//
func (rcv *linkDomainService) DeleteLink(
	ctx context.Context, in *link.DeleteLinkRequest,
) (*link.DeleteLinkResponse, error) {
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
	if err = rcv.linkRepository.Delete(in.GetLinkID(), in.GetUserID()); err != nil {
		return nil, rcv.wrapInternalError(errors.Wrap(err, "unable to remove link"))
	}

	//
	// Assemble response
	//
	return &link.DeleteLinkResponse{}, nil
}
