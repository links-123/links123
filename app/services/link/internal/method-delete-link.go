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
	ctx context.Context,
	in *link.DeleteLinkRequest,
	out *link.DeleteLinkResponse,
) error {
	var err error

	//
	// Validation
	//
	if in == nil {
		return rcv.wrapInvalidRequest(errors.New("request is empty"))
	}

	if err = in.Validate(); err != nil {
		return rcv.wrapInvalidRequest(err)
	}

	//
	// Request handing
	//
	if err = rcv.linkRepository.Delete(in.GetLinkID(), in.GetUserID()); err != nil {
		return rcv.wrapInternalError(errors.Wrap(err, "unable to remove link"))
	}

	//
	// Assemble response
	//
	out = &link.DeleteLinkResponse{}

	return nil
}
