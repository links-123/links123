package errors

import (
	"log"

	"github.com/ic2hrmk/links123/app/services/link/pb/link"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Internal(err error) error {
	log.Printf("[link-srv] Internal error: %s", err)
	return status.Error(codes.Code(link.LinkDomainServiceErrorCode_Internal), err.Error())
}

func InvalidRequest(err error) error {
	log.Printf("[link-srv] Invalid request: %s", err)
	return status.Error(codes.Code(link.LinkDomainServiceErrorCode_InvalidRequest), err.Error())
}
