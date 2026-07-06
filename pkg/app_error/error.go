package app_error

type AppError struct {
	Message string
	Status  int
}

func NewAppError(message string, status int) *AppError {
	return &AppError{
		Message: message,
		Status:  status,
	}
}

func (e *AppError) Error() string {
	return e.Message
}
