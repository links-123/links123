package internal

import (
	"log"

	"github.com/links-123/links123/app/services/link/pb/link"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (rcv *linkDomainService) wrapInternalError(err error) error {
	log.Printf("[link-srv] wrapInternalError error: %s", err)
	return status.Error(codes.Code(link.LinkDomainServiceErrorCode_Internal), err.Error())
}

func (rcv *linkDomainService) wrapInvalidRequest(err error) error {
	log.Printf("[link-srv] Invalid request: %s", err)
	return status.Error(codes.Code(link.LinkDomainServiceErrorCode_InvalidRequest), err.Error())
}
