package cerror

import "fmt"

type SyntaxError struct {
	Line   int
	Column int
	Desc   string
}

func (e SyntaxError) Error() string {
	return fmt.Sprintf("Line %d Column %d: %s", e.Line, e.Column, e.Desc)
}
