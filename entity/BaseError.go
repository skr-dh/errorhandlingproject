package entity

import "encoding/json"

type BaseError struct {
	Cause   error  `json:"detail"` // debatable to have it or not, but for the sake of idea presenting
	Message string `json:"message"`
	Status  int    `json:"-"`
}

func (e *BaseError) Error() string {
	// return both the message and original error
	return "Message : " + e.Message + " " + e.Cause.Error()
}

func (e *BaseError) ErrorResponseBody() []byte {
	// construct an error response body
	body, er := json.Marshal(e)
	if er != nil {
		// log
	}
	return body
}

func (e *BaseError) ErrorResponseHeaders() map[int]string {
	// get the status of the custom error thrown, header response type is redundant here, but
	// thought about what if it is streams
	return map[int]string{
		e.Status: "application/json; charset=utf-8",
	}
}
