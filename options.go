package client

import "time"

type Options struct {
	BasicAuth     *BasicAuthOption  `json:"basicAuth"`
	TLSOption     *TLSOption        `json:"tls"`
	HTTPTransport *TransportOption  `json:"httpTransport"`
	Middlewares   []Middleware      `json:"-"`
	CustomHeaders map[string]string `json:"customHeaders"`

	BearerToken string `json:"bearerToken"`
}

type BasicAuthOption struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type TLSOption struct {
	InsecureSkipVerify bool
	ServerName         string
	CAFile             string `json:"caFile"`
	CertFile           string `json:"certFile"`
	KeyFile            string `json:"keyFile"`
}

type TransportOption struct {
	// unix,tcp,stmp,file
	Protocol string `json:"protocol"`
	Addr     string `json:"addr"`

	DialContextTimeout    time.Duration `json:"dialContextTimeout"`
	DialContextKeepAlive  time.Duration `json:"dialContextKeepAlive"`
	ResponseHeaderTimeout time.Duration `json:"responseHeaderTimeout"`
	TLSHandshakeTimeout   time.Duration `json:"tlsHandshakeTimeout"`
	ExpectContinueTimeout time.Duration `json:"expectContinueTimeout"`
	MaxConnsPerHost       int           `json:"maxConnsPerHost"`
	MaxIdleConns          int           `json:"maxIdleConns"`
	MaxIdleConnsPerHost   int           `json:"maxIdleConnsPerHost"`
	IdleConnTimeout       time.Duration `json:"idleConnTimeout"`
}
