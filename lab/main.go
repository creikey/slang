package main

import (
	"fmt"
	"log"
)

type numberError struct {
	Offender    rune
	Description string
}

func (e numberError) Error() string {
	return fmt.Sprintf("Offending rune: %v | %s", e.Offender, e.Description)
}

func isNumber(toTest rune) error {
	if toTest >= '0' && toTest <= '9' {
		return nil
	}
	return numberError{toTest, "Not a number"}
}

func main() {
	test := isNumber('d')
	switch test {
	case numberError:
		log.Println("Found a number error")
		break
	default:
		log.Println("Default")
	}
	if test != nil {
		log.Fatal(test)
	}
}
