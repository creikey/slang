package main

import (
	"testing"

	"github.com/creikey/thelper"
)

func compareTokens(tok1, tok2 []Token) (bool, int) {
	for i := range tok1 {
		if tok1[i] != tok2[i] {
			return false, i
		}
	}
	return true, 0
}

func TestTell(t *testing.T) {
	cases := []struct {
		input    string
		line     int
		wantToks []Token
	}{
		{
			"TELL hello",
			0,
			[]Token{
				{"TELL", 0},
				{"hello", 5},
			},
		},
	}
	for _, val := range cases {
		toks, err := Tokenize(val.input, val.line)
		if err != nil {
			t.Error(err)
		}
		eq, index := compareTokens(val.wantToks, toks)
		if eq == false {
			t.Errorf("Wanted token (%v), but got token (%v) for input (%s)", val.wantToks[index], toks[index], val.input)
		}
	}
}

func TestTokenError(t *testing.T) {
	cases := []thelper.ErrNameCase{
		{
			InputDesc: "",
			ErrFunc: func() error {
				err := lintForDoubleSpaces("", 0)
				return err
			},
			WantErrName: "FatalError",
		},
		{
			InputDesc: " ",
			ErrFunc: func() error {
				err := lintForDoubleSpaces(" ", 0)
				return err
			},
			WantErrName: "SyntaxError",
		},
		{
			InputDesc: "fds",
			ErrFunc: func() error {
				err := lintForDoubleSpaces("fds", 0)
				return err
			},
			WantErrName: "nil",
		},
		{
			InputDesc: " d",
			ErrFunc: func() error {
				err := lintForDoubleSpaces(" d", 0)
				return err
			},
			WantErrName: "nil",
		},
	}
	thelper.CheckErrName(cases, t)
}

func TestTokenPrint(t *testing.T) {
	cases := []struct {
		input []Token
		want  string
		get   string
	}{
		{[]Token{{"test", 0}, {"other", 5}}, "(test, 0) (other, 5)", ""},
		{nil, "", ""},
	}
	for _, val := range cases {
		var err error
		val.get, err = printTokens(val.input)
		if err != nil {
			if val.input != nil {
				t.Errorf("Expected no error but got '%s'", err.Error())
			}
		}
		if val.get != val.want {
			t.Errorf("wanted string '%s', got '%s'", val.want, val.get)
		}
	}
}
