package main

import (
	"fmt"
	"os"
	"text/scanner"
)

func main() {
	var s scanner.Scanner
	s.Init(os.Stdin)
	s.Filename = "stdin"

	tok := s.Scan()
	for tok != scanner.EOF {
		fmt.Printf("Token: %s\n", s.TokenText())
		tok = s.Scan()
	}
}
