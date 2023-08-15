package createuser

import (
	"context"
	"errors"
)

type Handler struct{}

func New() *Handler {
	return &Handler{}
}

// Request is a well defined struct for a request data
// you can always see the input contract of your handler
type Request struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// Validate is a function that must exist on any request struct
// it returns error if request data is invalid and wrapper will return the code 400 (bad request)
func (r Request) Validate() error {
	if r.ID <= 0 {
		return ErrIdShouldBePositive
	}
	if r.Name == "" {
		return ErrEmptyName
	}
	if r.Age <= 0 {
		return ErrAgeShouldBePositive
	}
	return nil
}

// Response is a well defined struct for a response data
// here we just use an empty struct because we don't need to return anything
type Response struct {
}

// Better to keep all errors as variables for a performance reason and for a documentation reason
// except the errors which should contain some data in the text, they may be generated on-the-fly
var (
	ErrIdShouldBePositive  = errors.New("id should be a positive integer")
	ErrEmptyName           = errors.New("name should not be empty")
	ErrAgeShouldBePositive = errors.New("age should be a positive integer")
)

// Handle is called when the handler should be processed
func (h *Handler) Handle(ctx context.Context, request Request) (Response, error) {
	return Response{}, nil
}
