package serrors

import (
	"errors"
	"strings"
)

type ServiceError struct {
	name    string
	message string

	details  string
	wrapped  error
	httpCode int
}

func (s ServiceError) Error() string {
	parts := []string{s.message}
	if s.details != "" {
		parts = append(parts, s.details)
	}
	if s.wrapped != nil {
		parts = append(parts, s.wrapped.Error())
	}
	return strings.Join(parts, ": ")
}

func newError(name string, message string) ServiceError {
	return ServiceError{name: name, message: message}
}

func (s ServiceError) WithDetails(details string) ServiceError {
	s.details = details
	return s
}

func (s ServiceError) WithWrapped(wrapped error) ServiceError {
	s.wrapped = wrapped
	return s
}

func (s ServiceError) WithCode(code int) ServiceError {
	s.httpCode = code
	return s
}

func (s ServiceError) GetHttpCode() int {
	return s.httpCode
}

func (s ServiceError) Unwrap() error {
	return s.wrapped
}

func (s ServiceError) Is(err error) bool {
	if err == nil {
		return false
	}
	var sErr ServiceError
	if ok := errors.As(err, &sErr); !ok {
		return false
	}
	return s.name == sErr.name
}
