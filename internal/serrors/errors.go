package serrors

var (
	// ErrNotFound used when requested resource is not found
	ErrNotFound = newError("err-not-found", "Resource not found").WithCode(404)
	// ErrInternalServer used for any internal errors
	ErrInternalServer = newError("err-internal-server", "Internal server error").WithCode(500)
	// ErrInvalidArgument used when specified an invalid argument
	ErrInvalidArgument = newError("err-invalid-argument", "Invalid argument").WithCode(400)
)
