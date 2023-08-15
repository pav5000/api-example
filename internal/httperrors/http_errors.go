package httperrors

import "net/http"

type HttpError struct {
	Code int
	Err  error
}

func (e HttpError) Error() string {
	return e.Err.Error()
}

func New(err error, code int) HttpError {
	return HttpError{
		Code: code,
		Err:  err,
	}
}

func BadRequest(err error) HttpError {
	return New(err, http.StatusBadRequest)
}

func NotFound(err error) HttpError {
	return New(err, http.StatusNotFound)
}

// add some other custom error types if needed
