%{
package main

import (
	"fmt"
)

var identifiers []string
%}

%token ID

%%
Program:
	Statements

Statements:
	Statements Statement
	| Statement

Statement:
	Identifier ";"
	{
		identifiers = append(identifiers, $1)
		fmt.Printf("Swahili equivalent: %s\n", translateToSwahili($1))
	}

Identifier:
	ID

%%

var translations = map[string]string{
	"var":     "var",
	"func":    "kiungo",
	"if":      "kama",
	"else":    "vinginevyo",
	"for":     "kwa",
	"package": "mfuko",
	"import" : "leta",
	"Println" : "Chapisha"
	// Add more translations as needed
}

func translateToSwahili(englishKeyword string) string {
	if translation, ok := translations[englishKeyword]; ok {
		return translation
	}
	return englishKeyword
}
