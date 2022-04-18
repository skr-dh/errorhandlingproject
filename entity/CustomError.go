package entity

type CustomError interface {
	Error() string
	ErrorResponseBody() []byte
	ErrorResponseHeaders() map[int]string
}
