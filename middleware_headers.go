package client

import "net/http"

func SetHeaderMiddleware(opt Options, next http.RoundTripper) http.RoundTripper {
	if len(opt.CustomHeaders) == 0 {
		return next
	}
	rt := RoundTripperFunc(func(r *http.Request) (*http.Response, error) {
		for k, v := range opt.CustomHeaders {
			r.Header.Set(k, v)
		}
		return next.RoundTrip(r)
	})
	return rt
}
