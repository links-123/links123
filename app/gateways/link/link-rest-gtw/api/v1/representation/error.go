package representation

import "github.com/links-123/links123/shared/gateway/representation"

type ErrorResponse representation.ErrorResponse

//
// Verbose error messages
//
const (
	ErrIncorrectRequest = "IncorrectRequest"

	ErrInvalidLimitParameter  = "InvalidLimitParameter"
	ErrInvalidOffsetParameter = "InvalidOffsetParameter"

	ErrInvalidCreateRequest = "InvalidCreateRequest"
	ErrInvalidDeleteRequest = "InvalidDeleteRequest"
)
