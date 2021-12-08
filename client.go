package client

import (
	"net/http"
)

func NewClientWithOpts(opt Options) (*http.Client, error) {
	if opt.HTTPTransport == nil {
		return http.DefaultClient, nil
	}

	transport, err := GetTransport(opt)
	if err != nil {
		return nil, err
	}

	tc := &http.Client{
		Transport: transport,
		Timeout:   opt.HTTPTransport.DialContextTimeout,
	}

	return tc, nil
}
