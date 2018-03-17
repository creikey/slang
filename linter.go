package main

import (
	"fmt"
	"strings"
)

func lintForDoubleSpaces(toLint string, line int) *SyntaxError {
	if toLint == "" {
		return nil
	}
	prev := '0'
	for i, val := range toLint {
		if val == ' ' && prev == ' ' {
			return fmt.SyntaxErrorf("Line %d Column %d: Double spaces found", line, i)
		}
		prev = val
	}
	return nil
}

func sniffErrs(errs []SyntaxError) SyntaxError {
	for _, val := range errs {
		if val != nil {
			return val
		}
	}
	return nil
}

// Lint returns an SyntaxError after giving a string of an infinite amount of lines
func Lint(toLint string) SyntaxError {
	lines := strings.Split(toLint, "\n")
	var errs []SyntaxError
	for i, val := range lines {
		errs = append(errs, lintForDoubleSpaces(val, i))
		err := sniffErrs(errs)
		if err != nil {
			return err
		}
	}
	return nil
}
