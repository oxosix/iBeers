package errors

type HTTPError interface {
	StatusCode() int
	Error() string
}
