package apperrors

type ErrCode string

const (
	Unknown ErrCode = "U000"

	// sql error
	SelectErr       ErrCode = "SQL001"
	SelectNoRowsErr ErrCode = "SQL002"
	InsertErr       ErrCode = "SQL003"
	UpdateErr       ErrCode = "SQL004"
)

func (code ErrCode) Wrap(message string, err error) error {
	return &AppError{
		ErrCode: code,
		Message: message,
		Err:     err,
	}
}
