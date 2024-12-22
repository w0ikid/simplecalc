package errors

import "errors"

var (
	ErrInvalidExpression = errors.New("некорректное выражение")
	ErrInvalidCharacter  = errors.New("некорректный символ")
	ErrDivideByZero      = errors.New("деление на ноль")
)
