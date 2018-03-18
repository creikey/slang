package main

import (
	"fmt"
	"testing"
)

func TestNoContent(t *testing.T) {
	cases := []struct {
		input     FatalError
		wantError string
	}{
		{FatalError{line: 0, column: 0, Event: NOCONTENT}, fmt.Sprintf("Line: 0 Column: 0 | %s", NOCONTENT)},
	}
	for _, val := range cases {
		errString := val.input.Error()
		if errString != val.wantError {
			t.Errorf("Want error (%s) but got back error (%s) for input (%s)", val.wantError, errString, val.input)
		}
	}
}

func TestString(t *testing.T) {
	cases := []struct {
		input      FatalError
		wantString string
	}{
		{FatalError{line: 0, column: 0, Event: "testing"}, "(0), (0): (testing)"},
	}
	for _, val := range cases {
		gotString := val.input.String()
		if gotString != val.wantString {
			t.Errorf("Want string (%s) but got back string (%s) for input (%s)", val.wantString, gotString, val.input)
		}
	}
}
