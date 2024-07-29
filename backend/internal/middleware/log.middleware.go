package middleware

import (
	logging "backend/internal/logger"
	"net/http"
)

type wrappedWriter struct {
	http.ResponseWriter
	statusCode int
}

func (w *wrappedWriter) WriteHeader(statusCode int) {
	w.statusCode = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		// log the request

		wrapped := &wrappedWriter{res, http.StatusOK}
		// call the next middleware/handler
		query := req.URL.Query()
		if len(query) == 0 {
			logging.Logger.LogInfo().Fields(map[string]interface{}{
				"method": req.Method,
				"uri":    req.URL.Path,
				"query":  "No query",
				"status": wrapped.statusCode,
			}).Msg("Request")
		} else {
			logging.Logger.LogInfo().Fields(map[string]interface{}{
				"method": req.Method,
				"uri":    req.URL.Path,
				"query":  req.URL.Query(),
				"status": wrapped.statusCode,
			}).Msg("Request")
		}

		next.ServeHTTP(res, req)
	})
}
