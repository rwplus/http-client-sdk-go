package client

import "net/http"

func BasicAuthenticationMiddleware(opt Options, next http.RoundTripper) http.RoundTripper {
	if opt.BasicAuth == nil {
		return next
	}

	rt := RoundTripperFunc(func(r *http.Request) (*http.Response, error) {
		if r.Header.Get("Authorization") == "" {
			r.SetBasicAuth(opt.BasicAuth.Username, opt.BasicAuth.Password)
		}

		return next.RoundTrip(r)
	})

	return rt
}
