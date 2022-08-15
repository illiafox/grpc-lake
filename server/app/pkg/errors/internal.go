package errors

type InternalError struct {
	Err   error
	Scope string
}

func Convert(err error) (internal InternalError, ok bool) {
	internal, ok = err.(InternalError)

	return
}

func NewInternal(scope string, err error) error {
	return InternalError{
		Err:   err,
		Scope: scope,
	}
}
