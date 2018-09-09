package pqerrors_test

import (
	"fmt"
)

import (
	"github.com/dhui/pqerrors"
	"github.com/lib/pq"
)

func ExampleClass() {
	pqErr := pq.Error{Code: pq.ErrorCode("23505")}
	var err error = pqErr // err from database/sql
	if e, ok := err.(pq.Error); ok {
		switch e.Code.Class() {
		case pqerrors.PqErrClassIntegrityConstraintViolation:
			fmt.Println("Success!")
		default:
			fmt.Println("Unexpected error class:", e.Code.Class(), e.Code.Class().Name())
		}
	}
	// Output:
	// Success!
}

func ExampleCode() {
	pqErr := pq.Error{Code: pq.ErrorCode("23505")}
	var err error = pqErr // err from database/sql
	if e, ok := err.(pq.Error); ok {
		switch e.Code {
		case pqerrors.PqErrCodeIntegrityConstraintViolationUniqueViolation:
			fmt.Println("Success!")
		default:
			fmt.Println("Unexpected error code:", e.Code, e.Code.Name())
		}
	}
	// Output:
	// Success!
}