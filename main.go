package main

import (
	"bufio"
	"errors"
	"log"
	"os"
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
	rawSlices := strings.Split(raw, " ")
	toReturn := make([]Token, len(rawSlices), len(rawSlices))
	cur := 0
	for i, content := range rawSlices {
		toReturn[i] = Token{rawSlices[i], cur}
		cur += len(content)
	}
	return toReturn, nil
}

// Gets a line of input
func getInput(rdr *bufio.Reader) (string, error) {
	out, err := rdr.ReadString('\n')
	if err != nil {
		return "", err
	}
	return out[:len(out)-1], nil
}

func main() {
	// Define reader
	reader := bufio.NewReader(os.Stdin)
	text, err := getInput(reader)
	if err != nil {
		panic(err)
	}
	// Loop over until no more input
	for ; text != ""; text, _ = getInput(reader) {
		log.Println(text)
	}
}
