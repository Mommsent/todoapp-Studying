package core_http_response

import "net/http"

var (
	StatusCodeUninitialized = -1
)

type ResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func NewResponseWriter(rwriter http.ResponseWriter) *ResponseWriter {
	return &ResponseWriter{
		ResponseWriter: rwriter,
		statusCode:     StatusCodeUninitialized,
	}
}

func (rw *ResponseWriter) WriteHeader(statusCode int) {
	rw.ResponseWriter.WriteHeader(statusCode)
	rw.statusCode = statusCode
}

func (rw *ResponseWriter) GetSatusCodeOrPanic() int {

	if rw.statusCode == StatusCodeUninitialized {
		panic("no status code set")
	}

	return rw.statusCode
}
