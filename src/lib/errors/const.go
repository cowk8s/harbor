package errors

const (
	// NotFoundCode is code for the error of no object found
	NotFoundCode = "NOT_FOUND"
	// ConflictCode ...
	ConflictCode = "CONFLICT"
	// UnAuthorizedCode ...
	UnAuthorizedCode = "UNAUTHORIZED"
	// GeneralCode ...
	GeneralCode = "UNKNOWN"
)

// UnknownError ...
func UnknownError(err error) *Error {
	return New("unknown").WithCode(GeneralCode).WithCause(err)
}
