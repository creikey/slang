package main

import (
	"errors"
	"testing"
)

func errToBool(err error) bool {
	if err == nil {
		return false
	} else {
		return true
	}
}

func boolToError(bl bool) error {
	if bl == true {
		return errors.New("non-nil error")
	} else {
		return nil
	}
}

func TestDoubleSpace(t *testing.T) {
	cases := []struct {
		input   string
		wantErr bool
		getErr  error
	}{
		{"  ", true, nil},
		{" ", false, nil},
		{"fjdklsfjlkds   fdlksfjlsdkjflds fdslfdslf dsf  slkdj", true, nil},
		{"", false, nil},
		{"            ", true, nil},
	}
	for _, val := range cases {
		val.getErr = lintForDoubleSpaces(val.input, 0)
		errBool := errToBool(val.getErr)
		if errBool != val.wantErr {
			t.Errorf("Wanted (%v) error, but got back (%v) for input %s", boolToError(val.wantErr), val.getErr, val.input)
		}
	}
}

func TestLint(t *testing.T) {
	cases := []struct {
		input   string
		wantErr bool
		getErr  error
	}{
		{"fkjl djklsjdslkjfds  dlkfjdslkfj", true, nil},
	}

	for _, val := range cases {
		val.getErr = Lint(val.input)
		errBool := errToBool(val.getErr)
		if errBool != val.wantErr {
			t.Errorf("Wanted (%v) error, but got back (%v) for input %s", boolToError(val.wantErr), val.getErr, val.input)
		}
	}
}
