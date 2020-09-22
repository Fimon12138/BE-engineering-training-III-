package errors

type InternalException struct {
	message string
}

func (e *InternalException) Error() string {
	return e.message
}

func SetUpError(info string) *InternalException {
	return &InternalException{info}
}
