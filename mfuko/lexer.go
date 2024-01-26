package mfuko

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sync"
	"text/scanner"
)

var translations = map[string]string{
	"var":     "var",
	"func":    "kiungo",
	"if":      "kama",
	"else":    "vinginevyo",
	"for":     "kwa",
	"package": "mfuko",
	"import" : "leta",
	"Println" : "Chapisha"
}

func translateToSwahili(englishKeyword string) string {
	if translation, ok := translations[englishKeyword]; ok {
		return translation
	}
	return englishKeyword
}

func processChunk(wg *sync.WaitGroup, chunk []byte) {
	defer wg.Done()

	var s scanner.Scanner
	s.Init(bufio.NewReaderSize(bytes.NewReader(chunk), 1024)) // Use a buffered reader

	tok := s.Scan()
	for tok != scanner.EOF {
		fmt.Printf("Token: %s\n", translateToSwahili(s.TokenText()))
		tok = s.Scan()
	}
}

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

	const chunkSize = 1024 * 1024 // 1 MB chunks
	var wg sync.WaitGroup

	for {
		chunk := make([]byte, chunkSize)
		n, err := file.Read(chunk)
		if n == 0 || err != nil {
			break
		}

		wg.Add(1)
		go processChunk(&wg, chunk[:n])
	}

	wg.Wait()
}
