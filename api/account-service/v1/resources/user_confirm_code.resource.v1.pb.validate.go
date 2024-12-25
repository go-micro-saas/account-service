// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: api/account-service/v1/resources/user_confirm_code.resource.v1.proto

package resourcev1

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"

	enumv1 "github.com/go-micro-saas/account-service/api/account-service/v1/enums"
)

// ensure the imports are used
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
	_ = anypb.Any{}
	_ = sort.Sort

	_ = enumv1.UserVerifyTypeEnum_UserVerifyType(0)
)

// Validate checks the field values on UserConfirmCode with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *UserConfirmCode) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UserConfirmCode with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UserConfirmCodeMultiError, or nil if none found.
func (m *UserConfirmCode) ValidateAll() error {
	return m.validate(true)
}

func (m *UserConfirmCode) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for CreatedTime

	// no validation rules for UpdatedTime

	// no validation rules for UserIdentify

	// no validation rules for ConfirmType

	// no validation rules for ConfirmCode

	// no validation rules for ConfirmStatus

	// no validation rules for ConfirmTime

	// no validation rules for CancelTime

	if len(errors) > 0 {
		return UserConfirmCodeMultiError(errors)
	}

	return nil
}

// UserConfirmCodeMultiError is an error wrapping multiple validation errors
// returned by UserConfirmCode.ValidateAll() if the designated constraints
// aren't met.
type UserConfirmCodeMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UserConfirmCodeMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UserConfirmCodeMultiError) AllErrors() []error { return m }

// UserConfirmCodeValidationError is the validation error returned by
// UserConfirmCode.Validate if the designated constraints aren't met.
type UserConfirmCodeValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UserConfirmCodeValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UserConfirmCodeValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UserConfirmCodeValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UserConfirmCodeValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UserConfirmCodeValidationError) ErrorName() string { return "UserConfirmCodeValidationError" }

// Error satisfies the builtin error interface
func (e UserConfirmCodeValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUserConfirmCode.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UserConfirmCodeValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UserConfirmCodeValidationError{}
