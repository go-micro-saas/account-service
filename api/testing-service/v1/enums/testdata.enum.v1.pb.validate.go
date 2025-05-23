// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: api/testing-service/v1/enums/testdata.enum.v1.proto

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

// Validate checks the field values on TestdataInitEnum with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *TestdataInitEnum) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on TestdataInitEnum with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// TestdataInitEnumMultiError, or nil if none found.
func (m *TestdataInitEnum) ValidateAll() error {
	return m.validate(true)
}

func (m *TestdataInitEnum) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return TestdataInitEnumMultiError(errors)
	}

	return nil
}

// TestdataInitEnumMultiError is an error wrapping multiple validation errors
// returned by TestdataInitEnum.ValidateAll() if the designated constraints
// aren't met.
type TestdataInitEnumMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m TestdataInitEnumMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m TestdataInitEnumMultiError) AllErrors() []error { return m }

// TestdataInitEnumValidationError is the validation error returned by
// TestdataInitEnum.Validate if the designated constraints aren't met.
type TestdataInitEnumValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TestdataInitEnumValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TestdataInitEnumValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TestdataInitEnumValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TestdataInitEnumValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TestdataInitEnumValidationError) ErrorName() string { return "TestdataInitEnumValidationError" }

// Error satisfies the builtin error interface
func (e TestdataInitEnumValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTestdataInitEnum.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TestdataInitEnumValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TestdataInitEnumValidationError{}
