package main

import (
	"errors"
	"fmt"
	"strings"
)

// Token is a section of code with info for debugging
type Token struct {
	Content string
	Column  int
}

func (t Token) String() string {
	return fmt.Sprintf("(%s, %d)", t.Content, t.Column)
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

func printTokens(tok []Token) (string, error) {
	if tok == nil {
		return "", errors.New("Pass a non-nil slice of tokens")
	}
	toReturn := ""
	for _, val := range tok {
		toReturn += fmt.Sprintf("%v ", val)
	}
	return toReturn, nil
}
