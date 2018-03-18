package main

import (
	"bufio"
	"fmt"
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
	for ln := 0; text != ""; text, _ = getInput(reader) {
		tk, err := Tokenize(text, ln)
		if err != nil {
			log.Fatal(err)
		} else {
			fmt.Println(printTokens(tk))
		}
		ln++
	}
}
