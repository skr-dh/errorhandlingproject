package entity

import "encoding/json"

type ValidationErrors struct {
	Errors  map[string]string `json:"validation_errors"`
	Status  int               `json:"-"`
	Message string            `json:"message"`
}

func (e *ValidationErrors) Error() string {
	// return only the original message here
	return "Message : " + e.Message + " "
}

func (e *ValidationErrors) ErrorResponseBody() []byte {
	// construct an error response body
	body, er := json.Marshal(e)
	if er != nil {
		// log
	}
	return body
}

func (e *ValidationErrors) ErrorResponseHeaders() map[int]string {
	// get the status of the custom error thrown, header response type is redundant here, but
	// thought about what if it is streams
	return map[int]string{
		e.Status: "application/json; charset=utf-8",
	}
}
