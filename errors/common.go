package errors

type ValidationError struct {
}

func (e *ValidationError) Error() string {
	return "Validation Error"
}
