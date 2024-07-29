package middleware

import "net/http"

type Middleware func(http.Handler) http.Handler

func Chain(m ...Middleware) Middleware {
	return func(h http.Handler) http.Handler {
		for i := range m {
			h = m[i](h)
		}
		return h
	}
}
