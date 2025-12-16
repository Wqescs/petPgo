package parser

import (
	"strconv"
	"strings"

	"github.com/Wqescs/petPgo/calc/internal/domain"
)

func ParseExpression(expr string) (*domain.Expression, error) {
	if strings.TrimSpace(expr) == "" {
		return nil, domain.ErrEmptyExpression
	}

	expr = normalizeExpression(expr)
	
	operator, operatorIndex := findOperator(expr)
	if operatorIndex == -1 {
		return nil, domain.ErrMissingOperator
	}
	
	num1Str := expr[:operatorIndex]
	num2Str := expr[operatorIndex+1:]
	
	if num1Str == "" || num2Str == "" {
		return nil, domain.ErrInvalidExpression
	}
	
	num1, err := strconv.ParseFloat(num1Str, 64)
	if err != nil {
		return nil, domain.ErrInvalidNumber
	}
	
	num2, err := strconv.ParseFloat(num2Str, 64)
	if err != nil {
		return nil, domain.ErrInvalidNumber
	}
	
	return &domain.Expression{
		Num1:      num1,
		Num2:      num2,
		Operation: domain.Operation(operator),
	}, nil
}

func normalizeExpression(expr string) string {
	expr = strings.ReplaceAll(expr, ",", ".")
	expr = strings.ReplaceAll(expr, " ", "")
	return expr
}

func findOperator(expr string) (string, int) {
	operators := []string{"*", "/", "+", "-"}
	
	for i := len(expr) - 1; i >= 0; i-- {
		char := string(expr[i])
		for _, op := range operators {
			if char == op {
				if char == "-" && i == 0 {
					continue
				}
				if char == "-" && strings.ContainsAny(string(expr[i-1]), "+-*/") {
					continue
				}
				return char, i
			}
		}
	}
	
	return "", -1
}