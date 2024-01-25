// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: settings/v1/settings.proto

package settingsv1

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

// Validate checks the field values on OIDCConfig with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *OIDCConfig) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on OIDCConfig with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in OIDCConfigMultiError, or
// nil if none found.
func (m *OIDCConfig) ValidateAll() error {
	return m.validate(true)
}

func (m *OIDCConfig) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Name

	// no validation rules for Issuer

	// no validation rules for ClientId

	// no validation rules for CliClientId

	if len(errors) > 0 {
		return OIDCConfigMultiError(errors)
	}

	return nil
}

// OIDCConfigMultiError is an error wrapping multiple validation errors
// returned by OIDCConfig.ValidateAll() if the designated constraints aren't met.
type OIDCConfigMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m OIDCConfigMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m OIDCConfigMultiError) AllErrors() []error { return m }

// OIDCConfigValidationError is the validation error returned by
// OIDCConfig.Validate if the designated constraints aren't met.
type OIDCConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OIDCConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OIDCConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OIDCConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OIDCConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OIDCConfigValidationError) ErrorName() string { return "OIDCConfigValidationError" }

// Error satisfies the builtin error interface
func (e OIDCConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOIDCConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OIDCConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OIDCConfigValidationError{}

// Validate checks the field values on SettingsRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *SettingsRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SettingsRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// SettingsRequestMultiError, or nil if none found.
func (m *SettingsRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *SettingsRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return SettingsRequestMultiError(errors)
	}

	return nil
}

// SettingsRequestMultiError is an error wrapping multiple validation errors
// returned by SettingsRequest.ValidateAll() if the designated constraints
// aren't met.
type SettingsRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SettingsRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SettingsRequestMultiError) AllErrors() []error { return m }

// SettingsRequestValidationError is the validation error returned by
// SettingsRequest.Validate if the designated constraints aren't met.
type SettingsRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SettingsRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SettingsRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SettingsRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SettingsRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SettingsRequestValidationError) ErrorName() string { return "SettingsRequestValidationError" }

// Error satisfies the builtin error interface
func (e SettingsRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSettingsRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SettingsRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SettingsRequestValidationError{}

// Validate checks the field values on SettingsResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *SettingsResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SettingsResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// SettingsResponseMultiError, or nil if none found.
func (m *SettingsResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *SettingsResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Url

	if all {
		switch v := interface{}(m.GetOidcConfig()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, SettingsResponseValidationError{
					field:  "OidcConfig",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, SettingsResponseValidationError{
					field:  "OidcConfig",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetOidcConfig()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SettingsResponseValidationError{
				field:  "OidcConfig",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return SettingsResponseMultiError(errors)
	}

	return nil
}

// SettingsResponseMultiError is an error wrapping multiple validation errors
// returned by SettingsResponse.ValidateAll() if the designated constraints
// aren't met.
type SettingsResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SettingsResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SettingsResponseMultiError) AllErrors() []error { return m }

// SettingsResponseValidationError is the validation error returned by
// SettingsResponse.Validate if the designated constraints aren't met.
type SettingsResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SettingsResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SettingsResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SettingsResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SettingsResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SettingsResponseValidationError) ErrorName() string { return "SettingsResponseValidationError" }

// Error satisfies the builtin error interface
func (e SettingsResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSettingsResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SettingsResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SettingsResponseValidationError{}