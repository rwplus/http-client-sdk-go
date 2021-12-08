package client

import "net/http"

// Middleware
type Middleware interface {
	Use(opt Options, next http.RoundTripper) http.RoundTripper
}

// MiddlewareFunc is an adaptor
type MiddlewareFunc func(opt Options, next http.RoundTripper) http.RoundTripper

// MiddlewareFunc Use method implements middleware interface
func (fn MiddlewareFunc) Use(opt Options, next http.RoundTripper) http.RoundTripper {
	return fn(opt, next)
}

func applyMiddlewares(opt Options, finalTrip http.RoundTripper) http.RoundTripper {
	// clone finalTrip,avoid mutate it
	next := finalTrip
	for _, m := range opt.Middlewares {
		next = m.Use(opt, next)
	}

	return next
}
