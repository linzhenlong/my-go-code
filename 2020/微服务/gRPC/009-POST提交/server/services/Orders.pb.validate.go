// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: Orders.proto

package services

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
	_ = ptypes.DynamicAny{}
)

// define the regex for a UUID once up-front
var _orders_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on OrderRequest with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *OrderRequest) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetOrderMain()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return OrderRequestValidationError{
				field:  "OrderMain",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// OrderRequestValidationError is the validation error returned by
// OrderRequest.Validate if the designated constraints aren't met.
type OrderRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OrderRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OrderRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OrderRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OrderRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OrderRequestValidationError) ErrorName() string { return "OrderRequestValidationError" }

// Error satisfies the builtin error interface
func (e OrderRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOrderRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OrderRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OrderRequestValidationError{}

// Validate checks the field values on OrderResponse with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *OrderResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for ErrCode

	// no validation rules for ErrMsg

	for idx, item := range m.GetData() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return OrderResponseValidationError{
					field:  fmt.Sprintf("Data[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// OrderResponseValidationError is the validation error returned by
// OrderResponse.Validate if the designated constraints aren't met.
type OrderResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OrderResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OrderResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OrderResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OrderResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OrderResponseValidationError) ErrorName() string { return "OrderResponseValidationError" }

// Error satisfies the builtin error interface
func (e OrderResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOrderResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OrderResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OrderResponseValidationError{}
