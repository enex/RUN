package token

import (
	"fmt"
)

// Error represents an error in the source code. It consists of a position
// within the source files and message text describing the error
type Error struct {
	pos Position
	msg string
}

// Error generates an error string to satisfy the error interface
func (e Error) Error() string {
	return fmt.Sprint(e.pos, " ", e.msg)
}

// ErrorList is a slice of Error pointers
type ErrorList []*Error

// Count returns the number of errors within the list
func (el ErrorList) Count() int {
	return len(el)
}

// Add a new error to the list at the given position p.
func (el *ErrorList) Add(p Position, args ...interface{}) {
	*el = append(*el, &Error{pos: p, msg: fmt.Sprint(args...)})
}

func (el *ErrorList) cleanup() {
	var last Position
	i := 0
	for _, v := range *el {
		if v.pos != last {
			last = v.pos
			(*el)[i] = v
			i++
		}
	}
	(*el) = (*el)[:i]
}

// Error returns a string containing all the errors in the error list
func (el ErrorList) Error() string {
	var msg string
	el.cleanup()
	for i, err := range el {
		if i >= 10 {
			msg += fmt.Sprintln("More than 10 errors,", len(el)-10, "more not shown")
			break
		}
		msg += fmt.Sprintln(err)
	}
	return msg
}

// Print outputs a message containing all the errors in the list
func (el ErrorList) Print() {
	for _, err := range el {
		fmt.Println(err)
	}
}
