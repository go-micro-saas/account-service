// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: api/account-service/v1/enums/account.enum.v1.proto

package enumv1

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
)

// Validate checks the field values on AccountInitEnum with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *AccountInitEnum) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AccountInitEnum with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// AccountInitEnumMultiError, or nil if none found.
func (m *AccountInitEnum) ValidateAll() error {
	return m.validate(true)
}

func (m *AccountInitEnum) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return AccountInitEnumMultiError(errors)
	}

	return nil
}

// AccountInitEnumMultiError is an error wrapping multiple validation errors
// returned by AccountInitEnum.ValidateAll() if the designated constraints
// aren't met.
type AccountInitEnumMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AccountInitEnumMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AccountInitEnumMultiError) AllErrors() []error { return m }

// AccountInitEnumValidationError is the validation error returned by
// AccountInitEnum.Validate if the designated constraints aren't met.
type AccountInitEnumValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AccountInitEnumValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AccountInitEnumValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AccountInitEnumValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AccountInitEnumValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AccountInitEnumValidationError) ErrorName() string { return "AccountInitEnumValidationError" }

// Error satisfies the builtin error interface
func (e AccountInitEnumValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAccountInitEnum.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AccountInitEnumValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AccountInitEnumValidationError{}

// Validate checks the field values on UserStatusEnum with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *UserStatusEnum) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UserStatusEnum with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in UserStatusEnumMultiError,
// or nil if none found.
func (m *UserStatusEnum) ValidateAll() error {
	return m.validate(true)
}

func (m *UserStatusEnum) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return UserStatusEnumMultiError(errors)
	}

	return nil
}

// UserStatusEnumMultiError is an error wrapping multiple validation errors
// returned by UserStatusEnum.ValidateAll() if the designated constraints
// aren't met.
type UserStatusEnumMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UserStatusEnumMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UserStatusEnumMultiError) AllErrors() []error { return m }

// UserStatusEnumValidationError is the validation error returned by
// UserStatusEnum.Validate if the designated constraints aren't met.
type UserStatusEnumValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UserStatusEnumValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UserStatusEnumValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UserStatusEnumValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UserStatusEnumValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UserStatusEnumValidationError) ErrorName() string { return "UserStatusEnumValidationError" }

// Error satisfies the builtin error interface
func (e UserStatusEnumValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUserStatusEnum.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UserStatusEnumValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UserStatusEnumValidationError{}

// Validate checks the field values on UserGenderEnum with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *UserGenderEnum) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UserGenderEnum with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in UserGenderEnumMultiError,
// or nil if none found.
func (m *UserGenderEnum) ValidateAll() error {
	return m.validate(true)
}

func (m *UserGenderEnum) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return UserGenderEnumMultiError(errors)
	}

	return nil
}

// UserGenderEnumMultiError is an error wrapping multiple validation errors
// returned by UserGenderEnum.ValidateAll() if the designated constraints
// aren't met.
type UserGenderEnumMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UserGenderEnumMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UserGenderEnumMultiError) AllErrors() []error { return m }

// UserGenderEnumValidationError is the validation error returned by
// UserGenderEnum.Validate if the designated constraints aren't met.
type UserGenderEnumValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UserGenderEnumValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UserGenderEnumValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UserGenderEnumValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UserGenderEnumValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UserGenderEnumValidationError) ErrorName() string { return "UserGenderEnumValidationError" }

// Error satisfies the builtin error interface
func (e UserGenderEnumValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUserGenderEnum.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UserGenderEnumValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UserGenderEnumValidationError{}

// Validate checks the field values on UserRegisterTypeEnum with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UserRegisterTypeEnum) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UserRegisterTypeEnum with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UserRegisterTypeEnumMultiError, or nil if none found.
func (m *UserRegisterTypeEnum) ValidateAll() error {
	return m.validate(true)
}

func (m *UserRegisterTypeEnum) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return UserRegisterTypeEnumMultiError(errors)
	}

	return nil
}

// UserRegisterTypeEnumMultiError is an error wrapping multiple validation
// errors returned by UserRegisterTypeEnum.ValidateAll() if the designated
// constraints aren't met.
type UserRegisterTypeEnumMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UserRegisterTypeEnumMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UserRegisterTypeEnumMultiError) AllErrors() []error { return m }

// UserRegisterTypeEnumValidationError is the validation error returned by
// UserRegisterTypeEnum.Validate if the designated constraints aren't met.
type UserRegisterTypeEnumValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UserRegisterTypeEnumValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UserRegisterTypeEnumValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UserRegisterTypeEnumValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UserRegisterTypeEnumValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UserRegisterTypeEnumValidationError) ErrorName() string {
	return "UserRegisterTypeEnumValidationError"
}

// Error satisfies the builtin error interface
func (e UserRegisterTypeEnumValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUserRegisterTypeEnum.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UserRegisterTypeEnumValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UserRegisterTypeEnumValidationError{}

// Validate checks the field values on UserVerifyTypeEnum with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UserVerifyTypeEnum) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UserVerifyTypeEnum with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UserVerifyTypeEnumMultiError, or nil if none found.
func (m *UserVerifyTypeEnum) ValidateAll() error {
	return m.validate(true)
}

func (m *UserVerifyTypeEnum) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return UserVerifyTypeEnumMultiError(errors)
	}

	return nil
}

// UserVerifyTypeEnumMultiError is an error wrapping multiple validation errors
// returned by UserVerifyTypeEnum.ValidateAll() if the designated constraints
// aren't met.
type UserVerifyTypeEnumMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UserVerifyTypeEnumMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UserVerifyTypeEnumMultiError) AllErrors() []error { return m }

// UserVerifyTypeEnumValidationError is the validation error returned by
// UserVerifyTypeEnum.Validate if the designated constraints aren't met.
type UserVerifyTypeEnumValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UserVerifyTypeEnumValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UserVerifyTypeEnumValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UserVerifyTypeEnumValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UserVerifyTypeEnumValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UserVerifyTypeEnumValidationError) ErrorName() string {
	return "UserVerifyTypeEnumValidationError"
}

// Error satisfies the builtin error interface
func (e UserVerifyTypeEnumValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUserVerifyTypeEnum.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UserVerifyTypeEnumValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UserVerifyTypeEnumValidationError{}

// Validate checks the field values on UserVerifyStatusEnum with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UserVerifyStatusEnum) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UserVerifyStatusEnum with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UserVerifyStatusEnumMultiError, or nil if none found.
func (m *UserVerifyStatusEnum) ValidateAll() error {
	return m.validate(true)
}

func (m *UserVerifyStatusEnum) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return UserVerifyStatusEnumMultiError(errors)
	}

	return nil
}

// UserVerifyStatusEnumMultiError is an error wrapping multiple validation
// errors returned by UserVerifyStatusEnum.ValidateAll() if the designated
// constraints aren't met.
type UserVerifyStatusEnumMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UserVerifyStatusEnumMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UserVerifyStatusEnumMultiError) AllErrors() []error { return m }

// UserVerifyStatusEnumValidationError is the validation error returned by
// UserVerifyStatusEnum.Validate if the designated constraints aren't met.
type UserVerifyStatusEnumValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UserVerifyStatusEnumValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UserVerifyStatusEnumValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UserVerifyStatusEnumValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UserVerifyStatusEnumValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UserVerifyStatusEnumValidationError) ErrorName() string {
	return "UserVerifyStatusEnumValidationError"
}

// Error satisfies the builtin error interface
func (e UserVerifyStatusEnumValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUserVerifyStatusEnum.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UserVerifyStatusEnumValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UserVerifyStatusEnumValidationError{}
