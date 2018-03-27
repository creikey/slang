package cerror

import "testing"

func TestErrorSyntaxError(t *testing.T) {
	cases := []struct {
		input      SyntaxError
		wantOutput string
	}{
		{SyntaxError{4, 3, "command not found"}, "Line 4 Column 3: command not found"},
	}
	for _, val := range cases {
		errStr := val.input.Error()
		if errStr != val.wantOutput {
			t.Errorf("Wanted errstr (%s) but got (%s) for input (%s)", val.wantOutput, errStr, val.input)
		}
	}
}

func TestStringSyntaxError(t *testing.T) {
	cases := []struct {
		input      SyntaxError
		wantOutput string
	}{
		{SyntaxError{3, 51, "unknown character"}, "(3) (51): unknown character"},
	}
	for _, val := range cases {
		outStr := val.input.String()
		if outStr != val.wantOutput {
			t.Errorf("Wanted str (%s) but got (%s) calling (String()) for input (%s)", val.wantOutput, outStr, val.input)
		}
	}
}
