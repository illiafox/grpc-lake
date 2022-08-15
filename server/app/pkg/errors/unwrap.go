package errors

func (i InternalError) Unwrap() error {
	return i.Err
}

func (i InternalError) Cause() error {
	return i.Err
}
