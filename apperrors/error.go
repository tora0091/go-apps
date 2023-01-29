package apperrors

type AppError struct {
	ErrCode `json:"error_code"`
	Message string `json:"error_message"`
	Err     error  `json:"-"`
}

func (a *AppError) Error() string {
	return a.Err.Error()
}

func (a *AppError) Unwrap() error {
	return a.Err
}
