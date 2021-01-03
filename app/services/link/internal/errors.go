package internal

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/links-123/links123/app/services/link/pb/link"
)

func (rcv *linkDomainService) wrapInternalError(err error) error {
	rcv.log.Errorf("internal error: %s", err)
	return status.Error(codes.Code(link.LinkDomainServiceErrorCode_Internal), err.Error())
}

func (rcv *linkDomainService) wrapInvalidRequest(err error) error {
	rcv.log.Errorf("invalid request: %s", err)
	return status.Error(codes.Code(link.LinkDomainServiceErrorCode_InvalidRequest), err.Error())
}
