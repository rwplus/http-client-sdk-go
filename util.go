package client

import (
	"fmt"
	"net/url"
	"strings"
)

// ParseHost parse host
func ParseHost(host string) (*url.URL, error) {
	hostParts := strings.SplitN(host, "://", 2)
	if len(host) == 1 {
		return nil, fmt.Errorf("unable to parse host %s", host)
	}

	var basePath string
	schema, addr := hostParts[0], hostParts[1]
	if schema == "tcp" {
		parsed, err := url.Parse("tcp://" + addr)
		if err != nil {
			return nil, err
		}

		host = parsed.Host
		basePath = parsed.Path
	}

	return &url.URL{
		Scheme: schema,
		Host:   host,
		Path:   basePath,
	}, nil
}
