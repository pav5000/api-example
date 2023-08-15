package userbyid

import (
	"context"
	"errors"

	"github.com/pav5000/api-example/internal/httperrors"
)

type Handler struct{}

func New() *Handler {
	return &Handler{}
}

// Request is a well defined struct for a request data
// you can always see the input contract of your handler
type Request struct {
	ID int64 `json:"id"`
}

// Validate is a function that must exist on any request struct
// it returns error if request data is invalid and wrapper will return the code 400 (bad request)
func (r Request) Validate() error {
	if r.ID <= 0 {
		return ErrIdShouldBePositive
	}
	return nil
}

// Response is a well defined struct for a response data
// you can always see the output contract of your handler
type Response struct {
	Name string `json:"name"`
	Age  int32  `json:"age"`
}

// Better to keep all errors as variables for a performance reason and for a documentation reason
// except the errors which should contain some data in the text, they may be generated on-the-fly
var (
	ErrIdShouldBePositive = errors.New("id should be a positive integer")

	// this error if wrapped in a special type to tell the wrapper to return the code 404
	// the error will propagate up to the wrapper even if you use errors.WithMessage on it
	ErrNoSuchUser = httperrors.NotFound(errors.New("couldn't find user with this id"))
)

// Handle is called when the handler should be processed
func (h *Handler) Handle(ctx context.Context, request Request) (Response, error) {
	// using request.ID to get user from storage......
	if request.ID == 42 {
		// found a user
		return Response{
			Name: "Douglas",
			Age:  40,
		}, nil
	}

	// not found a user
	return Response{}, ErrNoSuchUser
}
