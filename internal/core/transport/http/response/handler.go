package core_http_response

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	core_errors "github.com/Mommsent/todoapp-Studying.git/internal/core/errors"
	core_logger "github.com/Mommsent/todoapp-Studying.git/internal/core/logger"
	"go.uber.org/zap"
)

type HTTPResponseHandler struct {
	log            *core_logger.Logger
	responseWriter http.ResponseWriter
}

func NewHTTPResponseHandler(log *core_logger.Logger, responseWriter http.ResponseWriter) *HTTPResponseHandler {
	return &HTTPResponseHandler{
		log:            log,
		responseWriter: responseWriter,
	}
}

func (handler *HTTPResponseHandler) JSONResponse(responseBody any, statusCode int) {
	handler.responseWriter.WriteHeader(statusCode)
	if err := json.NewEncoder(handler.responseWriter).Encode(responseBody); err != nil {
		handler.log.Error("write HTTP response: ", zap.Error(err))
	}
}

func (handler *HTTPResponseHandler) HTMLResponse(html []byte) {
	handler.responseWriter.WriteHeader(http.StatusOK)

	handler.responseWriter.Header().Set("Content-Type", "text/html; charset=utf-8")
	if _, err := handler.responseWriter.Write(html); err != nil {
		handler.log.Error("write HTML HTTP response: ", zap.Error(err))
	}
}

func (handler *HTTPResponseHandler) NoContentResponse() {
	handler.responseWriter.WriteHeader(http.StatusNoContent)
}

func (handler *HTTPResponseHandler) ErrorResponse(err error, msg string) {
	var (
		statusCode int
		logFunc    func(string, ...zap.Field)
	)

	switch {
	case errors.Is(err, core_errors.ErrInvalidArgument):
		statusCode = http.StatusBadRequest
		logFunc = handler.log.Warn

	case errors.Is(err, core_errors.ErrNotFound):
		statusCode = http.StatusNotFound
		logFunc = handler.log.Debug
	case errors.Is(err, core_errors.ErrConflict):
		statusCode = http.StatusConflict
		logFunc = handler.log.Warn
	default:
		statusCode = http.StatusInternalServerError
		logFunc = handler.log.Error
	}

	logFunc(msg, zap.Error(err))

	handler.writeResponse(statusCode, err, msg)
}

func (handler *HTTPResponseHandler) PanicResponse(panic any, msg string) {
	statusCode := http.StatusInternalServerError
	err := fmt.Errorf("unexpected panic: %v", panic)

	handler.log.Error(msg, zap.Error(err))
	handler.writeResponse(statusCode, err, msg)
}

func (handler *HTTPResponseHandler) writeResponse(statusCode int, err error, msg string) {
	response := ErrorResponse{
		Error:   err.Error(),
		Message: msg,
	}

	handler.JSONResponse(response, statusCode)
}
