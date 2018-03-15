package main

import (
	"fmt"
	"testing"
)

func TestTell(t *testing.T) {
	rawStr := "TELL hello testing more !! 3212j"
	wanted := []Token{{"TELL", 0}, {"hello", 5}, {"testing", 11}, {"more", 19}, {"!!", 24}, {"3212j", 27}}
	tok, err := tokenize(rawStr)
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
