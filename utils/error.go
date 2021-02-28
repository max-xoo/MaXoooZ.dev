package utils

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func NewError(code int, message string) *Error {
	err := &Error{}

	err.Code = code
	err.Message = message

	return err
}