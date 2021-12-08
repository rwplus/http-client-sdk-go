package client

import "net/http"

// RoundTripperFunc is an adaptor
type RoundTripperFunc func(r *http.Request) (*http.Response, error)

// RoundTrip implements the RoundTripper interface.
func (rt RoundTripperFunc) RoundTrip(r *http.Request) (*http.Response, error) {
	return rt(r)
}
