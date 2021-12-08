package client

import (
	"context"
	"net"
	"net/http"
)

func GetTransport(opt Options) (http.RoundTripper, error) {
	if opt.HTTPTransport == nil {
		return http.DefaultTransport, nil
	}

	tlsCfg, err := GetTLSConfig(opt)
	if err != nil {
		return nil, err
	}

	tsCfg := opt.HTTPTransport

	transport := &http.Transport{
		TLSClientConfig:       tlsCfg,
		Proxy:                 http.ProxyFromEnvironment,
		ResponseHeaderTimeout: tsCfg.ResponseHeaderTimeout,
		TLSHandshakeTimeout:   tsCfg.TLSHandshakeTimeout,
		ExpectContinueTimeout: tsCfg.ExpectContinueTimeout,
		MaxConnsPerHost:       tsCfg.MaxConnsPerHost,
		MaxIdleConns:          tsCfg.MaxIdleConns,
		MaxIdleConnsPerHost:   tsCfg.MaxIdleConnsPerHost,
		IdleConnTimeout:       tsCfg.IdleConnTimeout,
	}

	if transport.DialContext == nil && len(tsCfg.Protocol) == 0 {
		transport.DialContext = (&net.Dialer{
			Timeout:   tsCfg.DialContextTimeout,
			KeepAlive: tsCfg.DialContextKeepAlive,
		}).DialContext
	} else {
		transport.DialContext = func(_ context.Context, _, _ string) (net.Conn, error) {
			return dialContext(tsCfg.Protocol, tsCfg.Addr)
		}
	}

	return applyMiddlewares(opt, transport), nil
}
