package utils

import "github.com/go-playground/validator"

// RequestValidator exports interface to validate request payloads using structs
type RequestValidator struct {
	PayloadValidator *validator.Validate
}

// Validate invokes go-playground/validator validator using given interface
func (reqValidator *RequestValidator) Validate(i interface{}) error {
	return reqValidator.PayloadValidator.Struct(i)
}
