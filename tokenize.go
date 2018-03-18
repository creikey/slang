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
func Tokenize(raw string, line int) ([]Token, error) {
	if raw == "" {
		return nil, errors.New("No input for tokenizing provided")
	}
	if raw[len(raw)-1:] == " " {
		return nil, fmt.Errorf("Trailing space at line %d", line)
	}
	err := Lint(raw)
	if err != nil {
		return nil, err
	}
	rawSlices := strings.Split(raw, " ")
	toReturn := make([]Token, 0, len(rawSlices))
	cur := 0
	for _, content := range rawSlices {
		/*if content != "" && content != " " {
			toReturn[i] = Token{rawSlices[i], cur}
			cur += len(content) + 1
		}*/
		if content != "" {
			if content == " " {
				cur++
				continue
			}
			toReturn = append(toReturn, Token{content, cur})
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
	toReturn = toReturn[:len(toReturn)-1]
	return toReturn, nil
}
