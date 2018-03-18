package main

import (
	"reflect"
	"testing"

	"github.com/creikey/thelper"
)

func TestDoubleSpace(t *testing.T) {
	cases := []thelper.ErrNameCase{
		{
			InputDesc: " ",
			ErrFunc: func() error {
				err := lintForDoubleSpaces(" ", 0)
				return err
			},
			WantErrName: "nil",
		},
	}
	thelper.CheckErrName(cases, t)
	/*cases := []struct {
		input       string
		wantErrName string
	}{
		{"  ", "SyntaxError"},
		{" ", "nil"},
		{"fjdklsfjlkds   fdlksfjlsdkjflds fdslfdslf dsf  slkdj", "SyntaxError"},
		{"", "nil"},
		{"            ", "SyntaxError"},
	}
	for _, val := range cases {
		err := lintForDoubleSpaces(val.input, 0)
		if err == nil {
			if val.wantErrName != "nil" {
				t.Errorf("Wanted (%s) error name, but got back (nil) for input %s", val.wantErrName, val.input)
			}
		} else {
			errType := reflect.TypeOf(err)
			if errType.Name() != val.wantErrName {
				t.Errorf("Wanted (%s) error name, but got back (%s) for input %s", val.wantErrName, errType.Name(), val.input)
			}
		}
	}*/
}

/*func TestTrailingSpace(t *testing.T) {
	cases := []struct {
		input       string
		wantErrName string
	}{
		{"dlkfjldsfj", "nil"},
		{"fkldjfljfldksjfds ", "SyntaxError"},
	}

}*/

func TestLint(t *testing.T) {
	cases := []struct {
		input       string
		wantErrName string
	}{
		{"fkjl djklsjdslkjfds  dlkfjdslkfj", "SyntaxError"},
	}

	for _, val := range cases {
		err := lintForDoubleSpaces(val.input, 0)
		if err == nil && val.wantErrName != "nil" {
			t.Errorf("Wanted (%s) error name, but got back (nil) for input %s", val.wantErrName, val.input)
		}
		errType := reflect.TypeOf(err)
		if errType.Name() != val.wantErrName {
			t.Errorf("Wanted (%s) error name, but got back (%s) for input %s", val.wantErrName, errType.Name(), val.input)
		}
	}
}
