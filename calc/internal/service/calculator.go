package service

import (
	"math"
	"strconv"
	"strings"

	"github.com/Wqescs/petPgo/calc/internal/domain"
	"github.com/Wqescs/petPgo/calc/internal/parser"
	"github.com/Wqescs/petPgo/calc/pkg/decimal"
)

type Calculator struct{}

func New() *Calculator {
	return &Calculator{}
}

func (c *Calculator) Calculate(expression string) (*domain.Result, error) {
	expr, err := parser.ParseExpression(expression)
	if err != nil {
		return nil, err
	}
	
	var result float64
	
	switch expr.Operation {
	case domain.Add:
		result = expr.Num1 + expr.Num2
	case domain.Subtract:
		result = expr.Num1 - expr.Num2
	case domain.Multiply:
		result = expr.Num1 * expr.Num2
	case domain.Divide:
		if math.Abs(expr.Num2) < 1e-10 {
			return nil, domain.ErrDivisionByZero
		}
		result = expr.Num1 / expr.Num2
	default:
		return nil, domain.ErrUnknownOperation
	}
	
	precision := c.determinePrecision(expr.Num1, expr.Num2, expr.Operation)
	
	return &domain.Result{
		Value:      decimal.Round(result, precision),
		Expression: expression,
		Precision:  precision,
	}, nil
}

func (c *Calculator) determinePrecision(num1, num2 float64, op domain.Operation) int {
	switch op {
	case domain.Add, domain.Subtract:
		return max(getPrecision(num1), getPrecision(num2))
	case domain.Multiply:
		return getPrecision(num1) + getPrecision(num2)
	case domain.Divide:
		return min(getPrecision(num1)+getPrecision(num2), 10)
	default:
		return 6
	}
}

func getPrecision(num float64) int {
	str := strconv.FormatFloat(math.Abs(num), 'f', -1, 64)
	if idx := strings.Index(str, "."); idx != -1 {
		return len(str) - idx - 1
	}
	return 0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}