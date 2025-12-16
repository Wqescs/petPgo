package main

import (
	"flag"
	"fmt"

	"github.com/Wqescs/petPgo/calc/internal/handlers"
)

func main() {
	expression := flag.String("expr", "", "Expression to calculate")
	interactive := flag.Bool("i", false, "Interactive mode")
	flag.Parse()
	
	handler := handlers.NewCLIHandler()
	
	if *expression != "" {
		handler.ProcessSingle(*expression)
		return
	}
	
	handler.Run()
}