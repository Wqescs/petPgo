package decimal

import (
	"math"
	"strconv"
	"strings"
)

func ParseFloatWithPrecision(s string, bitSize int) (float64, int, error) {
	s = strings.ReplaceAll(s, ",", ".")
	s = strings.TrimSpace(s)
	
	val, err := strconv.ParseFloat(s, bitSize)
	if err != nil {
		return 0, 0, err
	}
	
	precision := 0
	if idx := strings.Index(s, "."); idx != -1 {
		precision = len(s) - idx - 1
	}
	
	return val, precision, nil
}

func Round(value float64, precision int) float64 {
	if precision <= 0 {
		return math.Round(value)
	}
	
	pow := math.Pow(10, float64(precision))
	return math.Round(value*pow) / pow
}

func Format(value float64, precision int) string {
	str := strconv.FormatFloat(value, 'f', precision, 64)
	str = strings.TrimRight(str, "0")
	str = strings.TrimRight(str, ".")
	return str
}