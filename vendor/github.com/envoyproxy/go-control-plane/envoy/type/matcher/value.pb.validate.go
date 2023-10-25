// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/type/matcher/value.proto

package matcher

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

// Validate checks the field values on ValueMatcher with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ValueMatcher) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ValueMatcher with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ValueMatcherMultiError, or
// nil if none found.
func (m *ValueMatcher) ValidateAll() error {
	return m.validate(true)
}

func (m *ValueMatcher) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	oneofMatchPatternPresent := false
	switch v := m.MatchPattern.(type) {
	case *ValueMatcher_NullMatch_:
		if v == nil {
			err := ValueMatcherValidationError{
				field:  "MatchPattern",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		oneofMatchPatternPresent = true

		if all {
			switch v := interface{}(m.GetNullMatch()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ValueMatcherValidationError{
						field:  "NullMatch",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ValueMatcherValidationError{
						field:  "NullMatch",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetNullMatch()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ValueMatcherValidationError{
					field:  "NullMatch",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *ValueMatcher_DoubleMatch:
		if v == nil {
			err := ValueMatcherValidationError{
				field:  "MatchPattern",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		oneofMatchPatternPresent = true

		if all {
			switch v := interface{}(m.GetDoubleMatch()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ValueMatcherValidationError{
						field:  "DoubleMatch",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ValueMatcherValidationError{
						field:  "DoubleMatch",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetDoubleMatch()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ValueMatcherValidationError{
					field:  "DoubleMatch",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *ValueMatcher_StringMatch:
		if v == nil {
			err := ValueMatcherValidationError{
				field:  "MatchPattern",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		oneofMatchPatternPresent = true

		if all {
			switch v := interface{}(m.GetStringMatch()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ValueMatcherValidationError{
						field:  "StringMatch",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ValueMatcherValidationError{
						field:  "StringMatch",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetStringMatch()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ValueMatcherValidationError{
					field:  "StringMatch",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *ValueMatcher_BoolMatch:
		if v == nil {
			err := ValueMatcherValidationError{
				field:  "MatchPattern",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		oneofMatchPatternPresent = true
		// no validation rules for BoolMatch
	case *ValueMatcher_PresentMatch:
		if v == nil {
			err := ValueMatcherValidationError{
				field:  "MatchPattern",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		oneofMatchPatternPresent = true
		// no validation rules for PresentMatch
	case *ValueMatcher_ListMatch:
		if v == nil {
			err := ValueMatcherValidationError{
				field:  "MatchPattern",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		oneofMatchPatternPresent = true

		if all {
			switch v := interface{}(m.GetListMatch()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ValueMatcherValidationError{
						field:  "ListMatch",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ValueMatcherValidationError{
						field:  "ListMatch",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetListMatch()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ValueMatcherValidationError{
					field:  "ListMatch",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	default:
		_ = v // ensures v is used
	}
	if !oneofMatchPatternPresent {
		err := ValueMatcherValidationError{
			field:  "MatchPattern",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return ValueMatcherMultiError(errors)
	}

	return nil
}

// ValueMatcherMultiError is an error wrapping multiple validation errors
// returned by ValueMatcher.ValidateAll() if the designated constraints aren't met.
type ValueMatcherMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ValueMatcherMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ValueMatcherMultiError) AllErrors() []error { return m }

// ValueMatcherValidationError is the validation error returned by
// ValueMatcher.Validate if the designated constraints aren't met.
type ValueMatcherValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ValueMatcherValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ValueMatcherValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ValueMatcherValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ValueMatcherValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ValueMatcherValidationError) ErrorName() string { return "ValueMatcherValidationError" }

// Error satisfies the builtin error interface
func (e ValueMatcherValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sValueMatcher.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ValueMatcherValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ValueMatcherValidationError{}

// Validate checks the field values on ListMatcher with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ListMatcher) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListMatcher with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ListMatcherMultiError, or
// nil if none found.
func (m *ListMatcher) ValidateAll() error {
	return m.validate(true)
}

func (m *ListMatcher) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	oneofMatchPatternPresent := false
	switch v := m.MatchPattern.(type) {
	case *ListMatcher_OneOf:
		if v == nil {
			err := ListMatcherValidationError{
				field:  "MatchPattern",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		oneofMatchPatternPresent = true

		if all {
			switch v := interface{}(m.GetOneOf()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ListMatcherValidationError{
						field:  "OneOf",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ListMatcherValidationError{
						field:  "OneOf",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetOneOf()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListMatcherValidationError{
					field:  "OneOf",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	default:
		_ = v // ensures v is used
	}
	if !oneofMatchPatternPresent {
		err := ListMatcherValidationError{
			field:  "MatchPattern",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return ListMatcherMultiError(errors)
	}

	return nil
}

// ListMatcherMultiError is an error wrapping multiple validation errors
// returned by ListMatcher.ValidateAll() if the designated constraints aren't met.
type ListMatcherMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListMatcherMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListMatcherMultiError) AllErrors() []error { return m }

// ListMatcherValidationError is the validation error returned by
// ListMatcher.Validate if the designated constraints aren't met.
type ListMatcherValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListMatcherValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListMatcherValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListMatcherValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListMatcherValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListMatcherValidationError) ErrorName() string { return "ListMatcherValidationError" }

// Error satisfies the builtin error interface
func (e ListMatcherValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListMatcher.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListMatcherValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListMatcherValidationError{}

// Validate checks the field values on ValueMatcher_NullMatch with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ValueMatcher_NullMatch) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ValueMatcher_NullMatch with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ValueMatcher_NullMatchMultiError, or nil if none found.
func (m *ValueMatcher_NullMatch) ValidateAll() error {
	return m.validate(true)
}

func (m *ValueMatcher_NullMatch) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return ValueMatcher_NullMatchMultiError(errors)
	}

	return nil
}

// ValueMatcher_NullMatchMultiError is an error wrapping multiple validation
// errors returned by ValueMatcher_NullMatch.ValidateAll() if the designated
// constraints aren't met.
type ValueMatcher_NullMatchMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ValueMatcher_NullMatchMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ValueMatcher_NullMatchMultiError) AllErrors() []error { return m }

// ValueMatcher_NullMatchValidationError is the validation error returned by
// ValueMatcher_NullMatch.Validate if the designated constraints aren't met.
type ValueMatcher_NullMatchValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ValueMatcher_NullMatchValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ValueMatcher_NullMatchValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ValueMatcher_NullMatchValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ValueMatcher_NullMatchValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ValueMatcher_NullMatchValidationError) ErrorName() string {
	return "ValueMatcher_NullMatchValidationError"
}

// Error satisfies the builtin error interface
func (e ValueMatcher_NullMatchValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sValueMatcher_NullMatch.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ValueMatcher_NullMatchValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ValueMatcher_NullMatchValidationError{}
