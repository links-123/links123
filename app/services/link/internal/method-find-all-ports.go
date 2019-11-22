package internal

import (
	"github.com/ic2hrmk/links123/app/services/link/errors"
	"github.com/ic2hrmk/links123/app/services/link/pb/link"
	"golang.org/x/net/context"
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
		return nil, errors.InvalidRequest(err)
	}

	//
	// Request handing
	//
	records, err := rcv.linkRepository.FindAll(in.GetLimit(), in.GetOffset())
	if err != nil {
		return nil, errors.Internal(err)
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
