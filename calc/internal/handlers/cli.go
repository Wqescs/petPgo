package handlers

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Wqescs/petPgo/calc/internal/service"
	"github.com/Wqescs/petPgo/calc/pkg/decimal"
)

type CLIHandler struct {
	calculator *service.Calculator
}

func NewCLIHandler() *CLIHandler {
	return &CLIHandler{
		calculator: service.New(),
	}
}

func (h *CLIHandler) Run() {
	scanner := bufio.NewScanner(os.Stdin)
	
	fmt.Println("🔢 Профессиональный калькулятор")
	fmt.Println("===============================")
	fmt.Println("Поддерживаемые операции: + - * /")
	fmt.Println("Примеры:")
	fmt.Println("  15.264-10.15365")
	fmt.Println("  3.14159 * 2.71828")
	fmt.Println("  100 / 3")
	fmt.Println("Введите 'help' для справки")
	fmt.Println("Введите 'exit' для выхода")
	fmt.Println()
	
	for {
		fmt.Print(">>> ")
		if !scanner.Scan() {
			break
		}
		
		input := strings.TrimSpace(scanner.Text())
		
		switch input {
		case "":
			continue
		case "exit", "quit", "q":
			fmt.Println("Выход из калькулятора")
			return
		case "help", "?":
			h.printHelp()
		default:
			h.processExpression(input)
		}
	}
}

func (h *CLIHandler) ProcessSingle(expr string) {
	result, err := h.calculator.Calculate(expr)
	if err != nil {
		fmt.Printf("❌ Ошибка: %v\n", err)
		return
	}
	
	formatted := decimal.Format(result.Value, result.Precision)
	fmt.Printf("✅ %s = %s\n", result.Expression, formatted)
}

func (h *CLIHandler) processExpression(expr string) {
	h.ProcessSingle(expr)
}

func (h *CLIHandler) printHelp() {
	helpText := `
Доступные команды:
  <выражение>  - вычислить выражение
  help, ?      - показать эту справку
  exit, quit, q - выйти

Формат выражений:
  • Числа могут быть целыми или десятичными
  • Используйте точку или запятую как разделитель
  • Поддерживаются пробелы в любом количестве
  • Поддерживаются отрицательные числа

Примеры:
  5 + 3.14
  -10 * 2.5
  100/3
  15,264 - 10,15365
`
	fmt.Println(helpText)
}