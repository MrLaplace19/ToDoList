package core_http_middleware

import (
	"context"
	"net/http"
	"time"
	core_logger "github.com/MrLaplace19/ToDoList/internal/core/logger"
	core_response "github.com/MrLaplace19/ToDoList/internal/core/transport/http/response"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

const (
	requestIdHeader = "X-Request-Id"
)

func RequestId() Middleware{
	return func(next http.Handler) http.Handler{
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			
			requestID := r.Header.Get(requestIdHeader)
			if requestID == ""{
				requestID = uuid.NewString()
			}

			r.Header.Set(requestIdHeader, requestID)
			w.Header().Set(requestIdHeader, requestID)
			next.ServeHTTP(w,r)
		})
	}
}


func Logger(log *core_logger.Logger) Middleware {
	return func(next http.Handler) http.Handler{
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
			requestID := r.Header.Get(requestIdHeader)

			l := log.With(
				zap.String("request_id", requestID),
				zap.String("url", r.URL.String()),
			)

			ctx := context.WithValue(r.Context(), "log", l)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func Panic() Middleware{
	return func(next http.Handler) http.Handler{
		return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request){
			ctx := r.Context()
			log := core_logger.FromContext(ctx)

			responseHandler := core_response.NewHTTPResponseHandler(log, w)

			defer func(){
				if p := recover(); p!=nil{
					responseHandler.PanicResponse(
						p,
						"During handle HTTP Request got unexpected panic",
					)
				}
			}()

			next.ServeHTTP(w,r)
		})
	}
}


func Trace() Middleware{
	return func(next http.Handler) http.Handler{
		return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request){
			ctx := r.Context()
			log := core_logger.FromContext(ctx)
			rw := core_response.NewResponseWriter(w)
			
			before := time.Now()

			log.Debug(
				">>> Incoming HTTP request",
				zap.Time("time", before.UTC()),
			)

			next.ServeHTTP(rw,r)

			log.Debug(
				"<<< Done HTTP request",
				zap.Int("status_code", rw.GetStatusCodeOrPanic()),
				zap.Duration("latency", time.Since(before)),
			)
		})
	}
}
