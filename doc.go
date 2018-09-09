// Package pqerrors provides constants for pq's (github.com/lib/pq) ErrorClass and ErrorCode
//
// ErrorClass constants are in the pqerrcls package and ErrorCode constants are in the pqerrcode package
package pqerrors

//go:generate mkdir -p pqerrcls pqerrcode
//go:generate go run generate.go
