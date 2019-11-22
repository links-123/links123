package internal

import (
	"github.com/ic2hrmk/links123/app/services/link/errors"
	"github.com/ic2hrmk/links123/app/services/link/pb/link"
	"github.com/ic2hrmk/links123/app/services/link/persistence/model"
	"golang.org/x/net/context"
)

//
// Saves link to persistence
//
func (rcv *linkDomainService) SaveLink(
	ctx context.Context, in *link.SaveLinkRequest,
) (*link.SaveLinkResponse, error) {
	//
	// Validation
	//
	if err := in.Validate(); err != nil {
		return nil, errors.InvalidRequest(err)
	}

	//
	// Request handing
	//
	if _, err := rcv.linkRepository.Save(&model.Link{
		LinkID:      in.GetLink().GetLinkID(),
		Name:        in.GetLink().GetName(),
		Code:        in.GetLink().GetCode(),
		Alias:       in.GetLink().GetAlias(),
		Unlocs:      in.GetLink().GetUnlocs(),
		Country:     in.GetLink().GetCountry(),
		Regions:     in.GetLink().GetRegions(),
		Province:    in.GetLink().GetProvince(),
		City:        in.GetLink().GetCity(),
		Coordinates: in.GetLink().GetCoordinates(),
		Timezone:    in.GetLink().GetTimezone(),
	}); err != nil {
		return nil, errors.Internal(err)
	}

	//
	// Assemble response
	//
	return &link.SaveLinkResponse{}, nil
}
