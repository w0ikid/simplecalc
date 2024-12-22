package service

import (
	"github.com/w0ikid/simplecalc/pkg/errors"
	"strconv"
	"strings"
	"unicode"
)

func Calculate(expression string) (string, error) {
	// Удаляем пробелы из выражения
	expression = strings.ReplaceAll(expression, " ", "")

	var nums []float64 // Стек чисел
	var ops []rune // Стек операций

	// Выполняет операцию из стека
	applyOperation := func() error {
		if len(nums) < 2 || len(ops) == 0 {
			return errors.ErrInvalidExpression
		}

		b := nums[len(nums)-1]
		a := nums[len(nums)-2]
		op := ops[len(ops)-1]
		nums = nums[:len(nums)-2]
		ops = ops[:len(ops)-1]

		switch op {
		case '+':
			nums = append(nums, a+b)
		case '-':
			nums = append(nums, a-b)
		case '*':
			nums = append(nums, a*b)
		case '/':
			if b == 0 {
				return errors.ErrDivideByZero
			}
			nums = append(nums, a/b)
		default:
			return errors.ErrInvalidExpression
		}
		return nil
	}

	// Определяет приоритет операций
	precedence := func(op rune) int {
		switch op {
		case '+', '-':
			return 1
		case '*', '/':
			return 2
		default:
			return 0
		}
	}

	// Парсинг числа из строки
	parseNumber := func(i *int) (float64, error) {
		start := *i
		for *i < len(expression) && (unicode.IsDigit(rune(expression[*i])) || expression[*i] == '.') {
			*i++
		}
		num, err := strconv.ParseFloat(expression[start:*i], 64)
		if err != nil {
			return 0, errors.ErrInvalidExpression
		}
		return num, nil
	}

	for i := 0; i < len(expression); i++ {
		ch := rune(expression[i])

		if (ch == '-' && (i == 0 || expression[i-1] == '(')) || unicode.IsDigit(ch) || ch == '.' {
			// Обрабатываем отрицательное или обычное число
			if ch == '-' && (i == 0 || expression[i-1] == '(') {
				i++
				if i >= len(expression) || (!unicode.IsDigit(rune(expression[i])) && expression[i] != '.') {
					return "", errors.ErrInvalidExpression
				}
				num, err := parseNumber(&i)
				if err != nil {
					return "", err
				}
				nums = append(nums, -num)
				i--
			} else {
				num, err := parseNumber(&i)
				if err != nil {
					return "", err
				}
				nums = append(nums, num)
				i--
			}
		} else if ch == '(' {
			ops = append(ops, ch)
		} else if ch == ')' {
			// Обрабатываем все операции до открывающей скобки
			for len(ops) > 0 && ops[len(ops)-1] != '(' {
				if err := applyOperation(); err != nil {
					return "", err
				}
			}
			if len(ops) == 0 || ops[len(ops)-1] != '(' {
				return "", errors.ErrInvalidExpression
			}
			ops = ops[:len(ops)-1] // Убираем '('
		} else if ch == '+' || ch == '-' || ch == '*' || ch == '/' {
			// Обрабатываем операции с более высоким приоритетом
			for len(ops) > 0 && precedence(ops[len(ops)-1]) >= precedence(ch) {
				if err := applyOperation(); err != nil {
					return "", err
				}
			}
			ops = append(ops, ch)
		} else {
			return "", errors.ErrInvalidCharacter
		}
	}

	// Выполняем оставшиеся операции
	for len(ops) > 0 {
		if err := applyOperation(); err != nil {
			return "", err
		}
	}

	// Проверяем корректность результата
	if len(nums) != 1 {
		return "", errors.ErrInvalidExpression
	}

	return strconv.FormatFloat(nums[0], 'f', -1, 64), nil
}
