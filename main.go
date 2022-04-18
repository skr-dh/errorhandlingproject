package main

import (
	"context"
	"encoding/json"
	"errorhandlingproject/entity"
	"errors"
	"net/http"
)

type Api interface {
	CreateData(ctx context.Context) ([]string, error)
}

type baseHandler func(w http.ResponseWriter, r *http.Request) error

func (b baseHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	err := b(writer, request)
	// here we are returning since there is nothing to be done if there are no errors
	if err == nil {
		return
	}

	if errors.As(err, entity.CustomError.Error) {
		customError, _ := err.(entity.CustomError)

		// log here once since all errors go through here
		// better place to ingest logging for errors and stacktrace in DD logs
		body := customError.ErrorResponseBody()
		writer.Write(body)

		// same for status headers as well
		headerInfo := customError.ErrorResponseHeaders()
		for k, v := range headerInfo {
			writer.Header().Set("Content-Type", v)
			writer.WriteHeader(k)
		}
	} else {
		// general/unhandled errors that might occur
		// log here as well, common ingestion place
		writer.WriteHeader(500)
		// may be construct a response body a common one for all Internal server errors
		return
	}

}

func testHandler(w http.ResponseWriter, r *http.Request) error {
	// call service
	// need to understand how a context is initialised, and what are the params involved in creating one
	var srv interface{}
	data, err := srv.(Api).CreateData(nil)
	if err != nil {
		return err
	} else {
		// write success data , maye be add header here as well since we have response writer
		w.WriteHeader(200)
		body, err := json.Marshal(data)
		if err == nil {
			w.Write(body)
		} else {
			return &entity.BaseError{
				Cause:   err,
				Status:  500,
				Message: "error occured while marshalling data during constructing response",
			}
		}
	}
	return nil
}

func main() {
	http.Handle("/test", baseHandler(testHandler))
	http.ListenAndServe(":8080", nil)
}
