package service

import (
	"errors"

	"github.com/emicklei/go-restful"

	"github.com/links-123/links123/app/gateways/link/link-rest-gtw/api/v2/representation"
	"github.com/links-123/links123/shared/gateway/helpers"
)

func (rcv *linkRESTService) notImplemented(request *restful.Request, response *restful.Response) {
	helpers.RespondWithBadRequest(
		response, rcv.log, errors.New("attempt to access not implemented API"),
		representation.ErrNotImplemented,
	)
}
