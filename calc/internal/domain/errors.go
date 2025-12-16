package domain

import "errors"

var (
	ErrInvalidExpression   = errors.New("invalid expression format")
	ErrInvalidNumber       = errors.New("invalid number format")
	ErrDivisionByZero      = errors.New("division by zero")
	ErrUnknownOperation    = errors.New("unknown operation")
	ErrEmptyExpression     = errors.New("empty expression")
	ErrMissingOperator     = errors.New("missing operator")
)