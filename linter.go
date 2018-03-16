package main

import (
	"fmt"
	"strings"
)

func lintForDoubleSpaces(toLint string, line int) error {
	if toLint == "" {
		return nil
	}
	prev := '0'
	for i, val := range toLint {
		if val == ' ' && prev == ' ' {
			return fmt.Errorf("Line %d Column %d: Double spaces found", line, i)
		}
		prev = val
	}
	return nil
}

func sniffErrs(errs []error) error {
	for _, val := range errs {
		if val != nil {
			return val
		}
	}
	return nil
}

// Lint returns an error after giving a string of an infinite amount of lines
func Lint(toLint string) error {
	lines := strings.Split(toLint, "\n")
	var errs []error
	for i, val := range lines {
		errs = append(errs, lintForDoubleSpaces(val, i))
		err := sniffErrs(errs)
		if err != nil {
			return err
		}
	}
	return nil
}
