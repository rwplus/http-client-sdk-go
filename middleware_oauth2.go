package client

import "net/http"

func OauthMiddleware(opt Options, next http.RoundTripper) http.RoundTripper {
	if opt.BasicAuth == nil {
		return next
	}

	rt := RoundTripperFunc(func(r *http.Request) (*http.Response, error) {
		if r.Header.Get("Authorization") == "" && len(opt.BearerToken) > 0 {
			r.Header.Add("Authorization", "Bearer "+opt.BearerToken)
		}

		return next.RoundTrip(r)
	})

	return rt
}
