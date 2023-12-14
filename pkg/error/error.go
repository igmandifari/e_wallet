package error

type Err struct {
	httpStatusCode int
	message        string
}

func (e *Err) Error() string {
	return e.message
}

func (e *Err) GetHTTPStatusCode() int {
	return e.httpStatusCode
}

func NewErr(httpStatusCode int, message string) *Err {
	return &Err{
		httpStatusCode: httpStatusCode,
		message:        message,
	}
}
