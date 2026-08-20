package core_http_response

import "net/http"

var (
	StatusCodeUninitialized = -1
)

type ResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func NewHTTPResponseWriter(w http.ResponseWriter) *ResponseWriter {
	return &ResponseWriter{
		ResponseWriter: w,
		statusCode:     StatusCodeUninitialized,
	}
}

func (rw *ResponseWriter) WriteHeader(code int) {
	rw.ResponseWriter.WriteHeader(code)
	rw.statusCode = code
}

func (rw *ResponseWriter) GetStatusCodeOrPanic() int {
	if rw.statusCode == StatusCodeUninitialized {
		return http.StatusOK
	}

	return rw.statusCode
}
