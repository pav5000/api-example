package wrapper

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/pav5000/api-example/internal/httperrors"
	"github.com/pkg/errors"
)

type Validator interface {
	Validate() error
}

// Wrap converts any function which looks like Handle(ctx, request) (response, error)
// to httprouter handler
// it does json marshal-unmarshal and request validation inside
func Wrap[
	Request Validator, Response any,
](
	cb func(context.Context, Request) (Response, error),
) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		err := handleRequest(w, r, cb)
		if err == nil {
			return
		}

		setJsonContentType(w)

		// Custom error processing here
		// feel free to add any other custom errors that you need
		{
			target := httperrors.HttpError{}
			// we try to cast common error type to a specific error
			// if cast succeeds we get true here and can use custom fields in target
			if errors.As(err, &target) {
				w.WriteHeader(target.Code)
			}
		}

		type StructuredError struct {
			Error string `json:"error"`
		}
		structError := StructuredError{
			Error: err.Error(),
		}
		rawJSON, _ := json.Marshal(structError)
		_, _ = w.Write(rawJSON)
		_, _ = w.Write([]byte("\n")) // just for easier life with curl
	}
}

func setJsonContentType(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}

// handleRequest is separated for easy error handling
// we just return errors from the function
func handleRequest[
	Request Validator, Response any,
](
	w http.ResponseWriter,
	r *http.Request,
	cb func(context.Context, Request) (Response, error),
) error {
	decoder := json.NewDecoder(r.Body)
	var request Request
	err := decoder.Decode(&request)
	if err != nil {
		return httperrors.New(
			errors.WithMessage(err, "decoding JSON"),
			http.StatusBadRequest,
		)
	}

	err = request.Validate()
	if err != nil {
		return httperrors.New(
			errors.WithMessage(err, "validating request"),
			http.StatusBadRequest,
		)
	}

	response, err := cb(r.Context(), request)
	if err != nil {
		return errors.WithMessage(err, "processing response")
	}

	setJsonContentType(w)
	encoder := json.NewEncoder(w)
	err = encoder.Encode(response)
	if err != nil {
		// use any log lib that you like here
		// we cannot return this error in response because we already wrote the header
		log.Println("error marshaling JSON:", err)
	}

	return nil
}
