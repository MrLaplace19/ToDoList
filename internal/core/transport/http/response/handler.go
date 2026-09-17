package core_http_response

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	core_errors "github.com/MrLaplace19/ToDoList/internal/core/errors"
	core_logger "github.com/MrLaplace19/ToDoList/internal/core/logger"
	"go.uber.org/zap"
)

type HTTPResponseHandler struct{
	log *core_logger.Logger
	w http.ResponseWriter
	
}

func NewHTTPResponseHandler(log *core_logger.Logger, w http.ResponseWriter) *HTTPResponseHandler{
	return &HTTPResponseHandler{
		log: log,
		w: w,
	}
}

func (h *HTTPResponseHandler) errorResponse(statuscode int, err error, msg string){
	

	response := map[string]string{
		"message": msg,
		"error": err.Error(),
	}

	h.JsonResponse(response, statuscode)
}

func (h *HTTPResponseHandler) PanicResponse(p any, msg string){
	StatusCode := http.StatusInternalServerError
	err := fmt.Errorf("unexpected panic: %v", p)

	h.log.Error(msg, zap.Error(err))

	h.errorResponse(
		StatusCode,
		err,
		msg,
	)
}

func (h *HTTPResponseHandler) ErrorResponse(err error, msg string){
	
	var (
		statusCode int
		logFunc func(string, ...zap.Field)
	)

	switch  {
	case errors.Is(err, core_errors.ErrInvalidArgument):
		statusCode = http.StatusBadRequest
		logFunc = h.log.Warn

	case errors.Is(err, core_errors.ErrNotFound):
		statusCode = http.StatusNotFound
		logFunc = h.log.Debug
	
	case errors.Is(err, core_errors.ErrConflict):
		statusCode = http.StatusConflict
		logFunc = h.log.Warn
	
	default:
		statusCode = http.StatusInternalServerError
		logFunc = h.log.Error
	}

	logFunc(msg, zap.Error(err))

	h.errorResponse(
		statusCode,
		err,
		msg,
	)

}

func (h *HTTPResponseHandler) JsonResponse(responseBody any, statuscode int){
	h.w.WriteHeader(statuscode)

	if err := json.NewEncoder(h.w).Encode(responseBody); err != nil{
		h.log.Error("write HTTP response", zap.Error(err))
	}
}
