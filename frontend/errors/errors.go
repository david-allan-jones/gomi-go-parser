package errors

import "fmt"

type Kind string

const (
	UnrecognizedError Kind = "UnrecognizedError"
	EofError          Kind = "EofError"
)

type Error struct {
	Line   int
	Column int
	Op     string
	Kind   Kind
	Err    error
}

func (e Error) Error() string {
	return fmt.Sprintf("%v errors with kind %v at %v:%v", e.Op, e.Kind, e.Line, e.Column)
}
