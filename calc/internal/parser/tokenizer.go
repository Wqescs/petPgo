package parser

import (
	"strings"
	"unicode"

	"github.com/Wqescs/calculator/internal/domain"
)

type TokenType int

const (
	Number TokenType = iota
	Operator
	ParenthesisOpen
	ParenthesisClose
)

type Token struct {
	Type  TokenType
	Value string
}

func Tokenize(expression string) ([]Token, error) {
	expression = normalizeExpression(expression)
	
	var tokens []Token
	var currentNumber strings.Builder
	
	for i, char := range expression {
		switch {
		case unicode.IsDigit(char) || char == '.' || char == ',':
			if char == ',' {
				currentNumber.WriteRune('.')
			} else {
				currentNumber.WriteRune(char)
			}
			
		case isOperator(char):
			if currentNumber.Len() > 0 {
				tokens = append(tokens, Token{Type: Number, Value: currentNumber.String()})
				currentNumber.Reset()
			}
			
			if char == '-' && (i == 0 || isOperator(rune(expression[i-1])) || expression[i-1] == '(') {
				currentNumber.WriteRune(char)
			} else {
				tokens = append(tokens, Token{Type: Operator, Value: string(char)})
			}
			
		case char == '(':
			if currentNumber.Len() > 0 {
				tokens = append(tokens, Token{Type: Number, Value: currentNumber.String()})
				currentNumber.Reset()
			}
			tokens = append(tokens, Token{Type: ParenthesisOpen, Value: "("})
			
		case char == ')':
			if currentNumber.Len() > 0 {
				tokens = append(tokens, Token{Type: Number, Value: currentNumber.String()})
				currentNumber.Reset()
			}
			tokens = append(tokens, Token{Type: ParenthesisClose, Value: ")"})
			
		case unicode.IsSpace(char):
			continue
			
		default:
			return nil, domain.ErrInvalidExpression
		}
	}
	
	if currentNumber.Len() > 0 {
		tokens = append(tokens, Token{Type: Number, Value: currentNumber.String()})
	}
	
	return tokens, nil
}

func isOperator(char rune) bool {
	return char == '+' || char == '-' || char == '*' || char == '/'
}

func normalizeExpression(expr string) string {
	
	expr = strings.ReplaceAll(expr, " ", "")
	
	return expr
}