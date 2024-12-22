package service

import (
	"github.com/w0ikid/simplecalc/pkg/errors"
	"testing"
)

func TestCalculate(t *testing.T) {
	tests := []struct {
		name       string
		expression string
		expected   string
		hasError   bool
		errType    error
	}{
		// Успешные вычисления
		{"Simple Addition", "1+2", "3", false, nil},
		{"Simple Multiplication", "2*3", "6", false, nil},
		{"Simple Division", "4/2", "2", false, nil},
		{"Simple Subtraction", "10-3", "7", false, nil},
		{"Mixed Operators", "3+5*2", "13", false, nil},
		{"Parentheses Priority", "(3+5)*2", "16", false, nil},
		{"Negative Numbers", "-5+3", "-2", false, nil},
		{"Negative in Parentheses", "(-5+3)*2", "-4", false, nil},
		{"Nested Parentheses", "(1+(2*3))", "7", false, nil},
		{"Double Parentheses", "((2+3)*4)", "20", false, nil},
		{"Floating Point Numbers", "2.5+3.1", "5.6", false, nil},
		{"Division with Floating Point", "5/2", "2.5", false, nil},
		{"Spaces in Expression", " 2 + 3 ", "5", false, nil},

		// Ошибки
		{"Invalid Characters", "2+*", "", true, errors.ErrInvalidExpression},
		{"Unmatched Parentheses", "(1+2", "", true, errors.ErrInvalidExpression},
		{"Empty Expression", "", "", true, errors.ErrInvalidExpression},
		{"Invalid Symbol", "2+@3", "", true, errors.ErrInvalidCharacter},
		{"Divide by Zero", "10/0", "", true, errors.ErrDivideByZero},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Calculate(tt.expression)

			if (err != nil) != tt.hasError {
				t.Errorf("expected error: %v, got: %v", tt.hasError, err)
			}

			if err != nil && tt.errType != nil && err != tt.errType {
				t.Errorf("expected error type: %v, got: %v", tt.errType, err)
			}

			if result != tt.expected {
				t.Errorf("expected result: %v, got: %v", tt.expected, result)
			}
		})
	}
}
