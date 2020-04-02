package service

import (
	"strconv"

	"github.com/links-123/links123/app/gateways/link/link-rest-gtw/api/v1/representation"

	"github.com/emicklei/go-restful"
	"github.com/pkg/errors"
)

func getLimitParameter(request *restful.Request) (uint64, error) {
	var limit uint64

	limitParam := request.QueryParameter(representation.LimitParameterName)
	if limitParam == "" {
		return representation.DefaultLimitValue, nil
	}

	limit, err := strconv.ParseUint(limitParam, 10, 64)
	if err != nil {
		return limit, errors.Errorf("failed to parse parameter [%s], %s",
			representation.LimitParameterName, err)
	}

	return limit, nil
}

func getOffsetParameter(request *restful.Request) (uint64, error) {
	var offset uint64

	limitParam := request.QueryParameter(representation.OffsetParameterName)
	if limitParam == "" {
		return representation.DefaultOffsetValue, nil
	}

	offset, err := strconv.ParseUint(limitParam, 10, 64)
	if err != nil {
		return offset, errors.Errorf("failed to parse parameter [%s], %s", representation.OffsetParameterName, err)
	}

	return offset, nil
}
