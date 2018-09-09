// DO NOT EDIT
//
// Code is auto-generated

package pqerrors

import (
	"github.com/lib/pq"
)

const (
	PqErrClassSuccessfulCompletion                    = pq.ErrorClass("00")
	PqErrClassWarning                                 = pq.ErrorClass("01")
	PqErrClassNoData                                  = pq.ErrorClass("02")
	PqErrClassSQLStatementNotYetComplete              = pq.ErrorClass("03")
	PqErrClassConnectionException                     = pq.ErrorClass("08")
	PqErrClassTriggeredActionException                = pq.ErrorClass("09")
	PqErrClassFeatureNotSupported                     = pq.ErrorClass("0A")
	PqErrClassInvalidTransactionInitiation            = pq.ErrorClass("0B")
	PqErrClassLocatorException                        = pq.ErrorClass("0F")
	PqErrClassInvalidGrantor                          = pq.ErrorClass("0L")
	PqErrClassInvalidRoleSpecification                = pq.ErrorClass("0P")
	PqErrClassDiagnosticsException                    = pq.ErrorClass("0Z")
	PqErrClassCaseNotFound                            = pq.ErrorClass("20")
	PqErrClassCardinalityViolation                    = pq.ErrorClass("21")
	PqErrClassDataException                           = pq.ErrorClass("22")
	PqErrClassIntegrityConstraintViolation            = pq.ErrorClass("23")
	PqErrClassInvalidCursorState                      = pq.ErrorClass("24")
	PqErrClassInvalidTransactionState                 = pq.ErrorClass("25")
	PqErrClassInvalidSQLStatementName                 = pq.ErrorClass("26")
	PqErrClassTriggeredDataChangeViolation            = pq.ErrorClass("27")
	PqErrClassInvalidAuthorizationSpecification       = pq.ErrorClass("28")
	PqErrClassDependentPrivilegeDescriptorsStillExist = pq.ErrorClass("2B")
	PqErrClassInvalidTransactionTermination           = pq.ErrorClass("2D")
	PqErrClassSQLRoutineException                     = pq.ErrorClass("2F")
	PqErrClassInvalidCursorName                       = pq.ErrorClass("34")
	PqErrClassExternalRoutineException                = pq.ErrorClass("38")
	PqErrClassExternalRoutineInvocationException      = pq.ErrorClass("39")
	PqErrClassSavepointException                      = pq.ErrorClass("3B")
	PqErrClassInvalidCatalogName                      = pq.ErrorClass("3D")
	PqErrClassInvalidSchemaName                       = pq.ErrorClass("3F")
	PqErrClassTransactionRollback                     = pq.ErrorClass("40")
	PqErrClassSyntaxErrorOrAccessRuleViolation        = pq.ErrorClass("42")
	PqErrClassWithCheckOptionViolation                = pq.ErrorClass("44")
	PqErrClassInsufficientResources                   = pq.ErrorClass("53")
	PqErrClassProgramLimitExceeded                    = pq.ErrorClass("54")
	PqErrClassObjectNotInPrerequisiteState            = pq.ErrorClass("55")
	PqErrClassOperatorIntervention                    = pq.ErrorClass("57")
	PqErrClassSystemError                             = pq.ErrorClass("58")
	PqErrClassConfigFileError                         = pq.ErrorClass("F0")
	PqErrClassFdwError                                = pq.ErrorClass("HV")
	PqErrClassPlpgsqlError                            = pq.ErrorClass("P0")
	PqErrClassInternalError                           = pq.ErrorClass("XX")
)
