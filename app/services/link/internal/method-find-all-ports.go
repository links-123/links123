package internal

import (
	"context"

	"github.com/pkg/errors"

	"github.com/ic2hrmk/links123/app/services/link/pb/link"
)

//
// Finds all existing links
//
func (rcv *linkDomainService) FindAllLinks(
	ctx context.Context, in *link.FindAllLinksRequest,
) (*link.FindAllLinksResponse, error) {
	//
	// Validation
	//
	if err := in.Validate(); err != nil {
		return nil, rcv.wrapInvalidRequest(err)
	}

	//
	// Request handing
	//
	records, err := rcv.linkRepository.FindAll(in.GetLimit(), in.GetOffset())
	if err != nil {
		return nil, rcv.wrapInternalError(errors.Wrap(err, "failed to perform search"))
	}

	//
	// Assemble response
	//
	out := &link.FindAllLinksResponse{
		Items: make([]*link.LinkEntity, len(records)),
	}

	for i := range records {
		out.Items[i] = &link.LinkEntity{
			LinkID:  records[i].LinkID,
			Name:    records[i].Name,
			Address: records[i].Address,
		}
	}

	return out, nil
}
