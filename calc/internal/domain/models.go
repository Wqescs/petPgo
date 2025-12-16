package domain

type Operation string

const (
	Add      Operation = "+"
	Subtract Operation = "-"
	Multiply Operation = "*"
	Divide   Operation = "/"
)

type Expression struct {
	Num1      float64
	Num2      float64
	Operation Operation
}

type Result struct {
	Value      float64
	Expression string
	Precision  int
}