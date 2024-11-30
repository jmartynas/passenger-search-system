package main

import "context"

// Validator is an object that can be validated.
type Validator interface {
	// Valid checks the object and returns any
	// problems. If len(problems) == 0 then
	// the object is valid.
	// key - field name
	// value - human-redable explanation of issue
	Valid(ctx context.Context) (problems map[string]string)
}
