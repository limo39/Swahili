package main

import (
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: program filename")
		os.Exit(1)
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Printf("Error opening file: %s\n", err)
		os.Exit(1)
	}
	defer file.Close()

	var s scanner.Scanner
	s.Init(file)
	s.Filename = os.Args[1]

	tok := s.Scan()
	for tok != scanner.EOF {
		yyParse(&s)
		tok = s.Scan()
	}
}
