// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: pb/link/pb.proto

package link

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/golang/protobuf/ptypes"
)

// ensure the imlinks are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = ptypes.DynamicAny{}
)

// Validate checks the field values on FindAllLinksRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *FindAllLinksRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Limit

	// no validation rules for Offset

	return nil
}

// FindAllLinksRequestValidationError is the validation error returned by
// FindAllLinksRequest.Validate if the designated constraints aren't met.
type FindAllLinksRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FindAllLinksRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FindAllLinksRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FindAllLinksRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FindAllLinksRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FindAllLinksRequestValidationError) ErrorName() string {
	return "FindAllLinksRequestValidationError"
}

// Error satisfies the builtin error interface
func (e FindAllLinksRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFindAllLinksRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FindAllLinksRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FindAllLinksRequestValidationError{}

// Validate checks the field values on FindAllLinksResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *FindAllLinksResponse) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetItems() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return FindAllLinksResponseValidationError{
					field:  fmt.Sprintf("Items[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// FindAllLinksResponseValidationError is the validation error returned by
// FindAllLinksResponse.Validate if the designated constraints aren't met.
type FindAllLinksResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FindAllLinksResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FindAllLinksResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FindAllLinksResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FindAllLinksResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FindAllLinksResponseValidationError) ErrorName() string {
	return "FindAllLinksResponseValidationError"
}

// Error satisfies the builtin error interface
func (e FindAllLinksResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFindAllLinksResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FindAllLinksResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FindAllLinksResponseValidationError{}

// Validate checks the field values on SaveLinkRequest with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *SaveLinkRequest) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetLink() == nil {
		return SaveLinkRequestValidationError{
			field:  "Link",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetLink()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SaveLinkRequestValidationError{
				field:  "Link",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// SaveLinkRequestValidationError is the validation error returned by
// SaveLinkRequest.Validate if the designated constraints aren't met.
type SaveLinkRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SaveLinkRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SaveLinkRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SaveLinkRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SaveLinkRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SaveLinkRequestValidationError) ErrorName() string { return "SaveLinkRequestValidationError" }

// Error satisfies the builtin error interface
func (e SaveLinkRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSaveLinkRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SaveLinkRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SaveLinkRequestValidationError{}

// Validate checks the field values on SaveLinkResponse with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *SaveLinkResponse) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// SaveLinkResponseValidationError is the validation error returned by
// SaveLinkResponse.Validate if the designated constraints aren't met.
type SaveLinkResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SaveLinkResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SaveLinkResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SaveLinkResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SaveLinkResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SaveLinkResponseValidationError) ErrorName() string { return "SaveLinkResponseValidationError" }

// Error satisfies the builtin error interface
func (e SaveLinkResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSaveLinkResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SaveLinkResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SaveLinkResponseValidationError{}

// Validate checks the field values on SaveLinksBulkRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *SaveLinksBulkRequest) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetItems() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return SaveLinksBulkRequestValidationError{
					field:  fmt.Sprintf("Items[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// SaveLinksBulkRequestValidationError is the validation error returned by
// SaveLinksBulkRequest.Validate if the designated constraints aren't met.
type SaveLinksBulkRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SaveLinksBulkRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SaveLinksBulkRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SaveLinksBulkRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SaveLinksBulkRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SaveLinksBulkRequestValidationError) ErrorName() string {
	return "SaveLinksBulkRequestValidationError"
}

// Error satisfies the builtin error interface
func (e SaveLinksBulkRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSaveLinksBulkRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SaveLinksBulkRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SaveLinksBulkRequestValidationError{}

// Validate checks the field values on SaveLinksBulkResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *SaveLinksBulkResponse) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// SaveLinksBulkResponseValidationError is the validation error returned by
// SaveLinksBulkResponse.Validate if the designated constraints aren't met.
type SaveLinksBulkResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SaveLinksBulkResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SaveLinksBulkResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SaveLinksBulkResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SaveLinksBulkResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SaveLinksBulkResponseValidationError) ErrorName() string {
	return "SaveLinksBulkResponseValidationError"
}

// Error satisfies the builtin error interface
func (e SaveLinksBulkResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSaveLinksBulkResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SaveLinksBulkResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SaveLinksBulkResponseValidationError{}

// Validate checks the field values on LinkEntity with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *LinkEntity) Validate() error {
	if m == nil {
		return nil
	}

	if utf8.RuneCountInString(m.GetLinkID()) < 1 {
		return LinkEntityValidationError{
			field:  "LinkID",
			reason: "value length must be at least 1 runes",
		}
	}

	if utf8.RuneCountInString(m.GetName()) < 1 {
		return LinkEntityValidationError{
			field:  "Name",
			reason: "value length must be at least 1 runes",
		}
	}

	// no validation rules for Code

	if utf8.RuneCountInString(m.GetCountry()) < 1 {
		return LinkEntityValidationError{
			field:  "Country",
			reason: "value length must be at least 1 runes",
		}
	}

	// no validation rules for Province

	if utf8.RuneCountInString(m.GetCity()) < 1 {
		return LinkEntityValidationError{
			field:  "City",
			reason: "value length must be at least 1 runes",
		}
	}

	// no validation rules for Timezone

	return nil
}

// LinkEntityValidationError is the validation error returned by
// LinkEntity.Validate if the designated constraints aren't met.
type LinkEntityValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LinkEntityValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LinkEntityValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LinkEntityValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LinkEntityValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LinkEntityValidationError) ErrorName() string { return "LinkEntityValidationError" }

// Error satisfies the builtin error interface
func (e LinkEntityValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLinkEntity.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LinkEntityValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LinkEntityValidationError{}
