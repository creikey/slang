package cerror

import "fmt"

// SyntaxError is an error based on syntax alone, like double spaces or unknown tokens
type SyntaxError struct {
	Line   int
	Column int
	Desc   string
}

func (e SyntaxError) Error() string {
	return fmt.Sprintf("Line %d Column %d: %s", e.Line, e.Column, e.Desc)
}

func (e SyntaxError) String() string {
	return fmt.Sprintf("(%d) (%d): %s", e.Line, e.Column, e.Desc)
}
