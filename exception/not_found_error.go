package exception

type NotFoundError struct {
	Error string
}

func NewNotFoudError(error string) NotFoundError {
	return NotFoundError{Error: error}
}
