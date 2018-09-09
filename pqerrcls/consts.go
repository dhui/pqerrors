// Package pqerrcls provides consts for pq.ErrorClass
//
// DO NOT EDIT
//
// Code is auto-generated

package pqerrcls

import (
	"github.com/lib/pq"
)

const (
	SuccessfulCompletion                    = pq.ErrorClass("00")
	Warning                                 = pq.ErrorClass("01")
	NoData                                  = pq.ErrorClass("02")
	SQLStatementNotYetComplete              = pq.ErrorClass("03")
	ConnectionException                     = pq.ErrorClass("08")
	TriggeredActionException                = pq.ErrorClass("09")
	FeatureNotSupported                     = pq.ErrorClass("0A")
	InvalidTransactionInitiation            = pq.ErrorClass("0B")
	LocatorException                        = pq.ErrorClass("0F")
	InvalidGrantor                          = pq.ErrorClass("0L")
	InvalidRoleSpecification                = pq.ErrorClass("0P")
	DiagnosticsException                    = pq.ErrorClass("0Z")
	CaseNotFound                            = pq.ErrorClass("20")
	CardinalityViolation                    = pq.ErrorClass("21")
	DataException                           = pq.ErrorClass("22")
	IntegrityConstraintViolation            = pq.ErrorClass("23")
	InvalidCursorState                      = pq.ErrorClass("24")
	InvalidTransactionState                 = pq.ErrorClass("25")
	InvalidSQLStatementName                 = pq.ErrorClass("26")
	TriggeredDataChangeViolation            = pq.ErrorClass("27")
	InvalidAuthorizationSpecification       = pq.ErrorClass("28")
	DependentPrivilegeDescriptorsStillExist = pq.ErrorClass("2B")
	InvalidTransactionTermination           = pq.ErrorClass("2D")
	SQLRoutineException                     = pq.ErrorClass("2F")
	InvalidCursorName                       = pq.ErrorClass("34")
	ExternalRoutineException                = pq.ErrorClass("38")
	ExternalRoutineInvocationException      = pq.ErrorClass("39")
	SavepointException                      = pq.ErrorClass("3B")
	InvalidCatalogName                      = pq.ErrorClass("3D")
	InvalidSchemaName                       = pq.ErrorClass("3F")
	TransactionRollback                     = pq.ErrorClass("40")
	SyntaxErrorOrAccessRuleViolation        = pq.ErrorClass("42")
	WithCheckOptionViolation                = pq.ErrorClass("44")
	InsufficientResources                   = pq.ErrorClass("53")
	ProgramLimitExceeded                    = pq.ErrorClass("54")
	ObjectNotInPrerequisiteState            = pq.ErrorClass("55")
	OperatorIntervention                    = pq.ErrorClass("57")
	SystemError                             = pq.ErrorClass("58")
	ConfigFileError                         = pq.ErrorClass("F0")
	FdwError                                = pq.ErrorClass("HV")
	PlpgsqlError                            = pq.ErrorClass("P0")
	InternalError                           = pq.ErrorClass("XX")
)
