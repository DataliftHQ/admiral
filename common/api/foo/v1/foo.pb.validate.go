// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: foo/v1/foo.proto

package foov1

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

// Validate checks the field values on GetFooRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *GetFooRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetFooRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in GetFooRequestMultiError, or
// nil if none found.
func (m *GetFooRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetFooRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return GetFooRequestMultiError(errors)
	}

	return nil
}

// GetFooRequestMultiError is an error wrapping multiple validation errors
// returned by GetFooRequest.ValidateAll() if the designated constraints
// aren't met.
type GetFooRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetFooRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetFooRequestMultiError) AllErrors() []error { return m }

// GetFooRequestValidationError is the validation error returned by
// GetFooRequest.Validate if the designated constraints aren't met.
type GetFooRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetFooRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetFooRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetFooRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetFooRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetFooRequestValidationError) ErrorName() string { return "GetFooRequestValidationError" }

// Error satisfies the builtin error interface
func (e GetFooRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetFooRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetFooRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetFooRequestValidationError{}

// Validate checks the field values on GetFooResponse with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *GetFooResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetFooResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in GetFooResponseMultiError,
// or nil if none found.
func (m *GetFooResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetFooResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Foo

	if len(errors) > 0 {
		return GetFooResponseMultiError(errors)
	}

	return nil
}

// GetFooResponseMultiError is an error wrapping multiple validation errors
// returned by GetFooResponse.ValidateAll() if the designated constraints
// aren't met.
type GetFooResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetFooResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetFooResponseMultiError) AllErrors() []error { return m }

// GetFooResponseValidationError is the validation error returned by
// GetFooResponse.Validate if the designated constraints aren't met.
type GetFooResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetFooResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetFooResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetFooResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetFooResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetFooResponseValidationError) ErrorName() string { return "GetFooResponseValidationError" }

// Error satisfies the builtin error interface
func (e GetFooResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetFooResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetFooResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetFooResponseValidationError{}
