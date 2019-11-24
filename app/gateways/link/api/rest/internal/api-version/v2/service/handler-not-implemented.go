package service

import (
	"errors"
	"github.com/emicklei/go-restful"

	"github.com/ic2hrmk/links123/app/gateways/link/api/rest/internal/api-version/v2/representation"
	"github.com/ic2hrmk/links123/shared/gateway/helpers"
)

func (rcv *linkRESTService) notImplemented(request *restful.Request, response *restful.Response) {
	helpers.ResponseWithBadRequest(
		response, errors.New("attempt to access not implemented API"),
		representation.ErrNotImplemented,
	)
}
