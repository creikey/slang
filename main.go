package main

import (
	"bufio"
	"log"
	"os"
)

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
		tk, err := tokenize(text)
		if err != nil {
			log.Fatal(err)
		} else {
			printTokens(tk)
		}
	}
}
