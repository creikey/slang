package cerror

//TODO refactor entire library so that there's a standard error command and not a janky unit test library

import (
	"fmt"
)

// FatalEvent enum for describing fatal errors
type FatalEvent string

const (
	// NOCONTENT is when nothing is provided to be interpreted
	NOCONTENT = FatalEvent("nothing is provided to be interpreted")
)

// FatalError is a catastrophic error where the interpreter cannot continue
type FatalError struct {
	line   int
	column int
	Event  FatalEvent
}

func (fe FatalError) Error() string {
	return fmt.Sprintf("Line: %d Column: %d | %s", fe.line, fe.column, fe.Event)
}

func (fe FatalError) String() string {
	return fmt.Sprintf("(%d), (%d): (%s)", fe.line, fe.column, fe.Event)
}
