package main

import (
	"fmt"
	"log"
	"reflect"
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
	v := reflect.TypeOf(test)
	if v.Name() == "numberError" {
		fmt.Println("Num error")
	} else {
		fmt.Println(fmt.Sprintf("Not num error. Is %s", v.Name()))
	}
	if test != nil {
		log.Fatal(test)
	}
}
