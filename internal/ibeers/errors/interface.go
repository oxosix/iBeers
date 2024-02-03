package errors

type HTTPError interface {
	error
	StatusCode() int      // Retorna o código de status HTTP
	ErrorMessage() string // Retorna a mensagem de erro
}
