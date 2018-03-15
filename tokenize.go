package main

import (
	"errors"
	"strings"
)

// Token is a section of code with info for debugging
type Token struct {
	Content string
	Column  int
}

// Tokenizes input
func tokenize(raw string) ([]Token, error) {
	if raw == "" {
		return nil, errors.New("No input for tokenizing provided")
	}
	if raw[len(raw)-1:] == " " {
		return nil, errors.New("Trailing space")
	}
	rawSlices := strings.Split(raw, " ")
	toReturn := make([]Token, len(rawSlices), len(rawSlices))
	cur := 0
	for i, content := range rawSlices {
		if content != "" {
			toReturn[i] = Token{rawSlices[i], cur}
			cur += len(content) + 1
		}
	}
	return toReturn, nil
}
