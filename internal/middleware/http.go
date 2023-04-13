package middleware

import (
	"encoding/json"
	"net/http"
	"time"
)

type Handler func(w http.ResponseWriter, r *http.Request) (interface{}, error)

func ErrorMiddleware(h Handler) http.Handler {
	return corsMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()

		response, err := h(w, r)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			HandleHTTPMiddlewareError(
				w,
				r,
				startTime,
				err,
			)
			return
		}

		if response != nil {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(response)
		}
	}))
}

type httpErrorResponse struct {
	Code    string
	Message string
	Debug   string
}

func HandleHTTPMiddlewareError(
	w http.ResponseWriter,
	r *http.Request,
	startTime time.Time,
	err error,
) {
	json.NewEncoder(w).Encode(httpErrorResponse{
		Code:    err.Error(),
		Message: err.Error(),
		Debug:   err.Error(),
	})
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		allowedorigins := "http://localhost:3000"
		if len(allowedorigins) != 0 {
			h := w.Header()
			h.Set("Access-Control-Allow-Origin", allowedorigins)
			h.Set("Access-Control-Allow-Headers", "Origin, Content-Type")
			h.Set("Access-Control-Allow-Credentials", "true")
			h.Set("Access-Control-Allow-Methods", "POST")
		}

		if r.Method == http.MethodOptions {
			return
		}

		next.ServeHTTP(w, r)
	})
}
