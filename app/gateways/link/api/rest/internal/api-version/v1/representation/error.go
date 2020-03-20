package representation

import "github.com/ic2hrmk/links123/shared/gateway/representation"

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
