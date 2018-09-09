package pqerrors_test

import (
	"fmt"
)

import (
	"github.com/dhui/pqerrors/pqerrcls"
	"github.com/dhui/pqerrors/pqerrcode"
	"github.com/lib/pq"
)

func Example() {
	pqErr := pq.Error{Code: pq.ErrorCode("23505")}
	var err error = pqErr // err from database/sql
	if e, ok := err.(pq.Error); ok {
		// Example usage of pqerrcls
		switch e.Code.Class() {
		case pqerrcls.IntegrityConstraintViolation:
			fmt.Println("Class success!")
		default:
			fmt.Println("Unexpected error class:", e.Code.Class(), e.Code.Class().Name())
		}

		// Example usage of pqerrcode
		switch e.Code {
		case pqerrcode.IntegrityConstraintViolationUniqueViolation:
			fmt.Println("Code success!")
		default:
			fmt.Println("Unexpected error code:", e.Code, e.Code.Name())
		}
	}
	// Output:
	// Class success!
	// Code success!
}
