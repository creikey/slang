package main

import (
	"fmt"
	"testing"
)

type testCase struct {
	want string
	get  string
}

func TestTell(t *testing.T) {
	rawStr := "TELL hello testing more !! 3212j"
	wanted := []Token{{"TELL", 0}, {"hello", 5}, {"testing", 11}, {"more", 19}, {"!!", 24}, {"3212j", 27}}
	tok, err := tokenize(rawStr, 0)
	if err != nil {
		t.Error(err)
	}
	for i := range tok {
		if tok[i].Content != wanted[i].Content {
			t.Error(fmt.Errorf("token %s not equal to token %s at %d", tok[i].Content, wanted[i].Content, i))
		}
		if tok[i].Column != wanted[i].Column {
			t.Error(fmt.Errorf("token column %d for token %s not equal to token column %d for token %s", tok[i].Column, tok[i].Content, wanted[i].Column, wanted[i].Content))
		}
	}
}

func TestTokenError(t *testing.T) {
	cases := []struct {
		input      string
		wantNotNil bool
		getNotNil  bool
	}{
		{"", true, false},
		{" ", true, false},
		{"fds", false, false},
		{"fddsfdks ", true, false},
		{"  d", true, false},
	}
	for i, val := range cases {
		_, err := tokenize(val.input, i)
		if err != nil {
			val.getNotNil = true
		} else {
			val.getNotNil = false
		}
		if val.getNotNil != val.wantNotNil {
			t.Errorf("expected '%v', got '%v' with input '%s' and error '%s'", val.wantNotNil, val.getNotNil, val.input, err.Error())
		}
	}
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
