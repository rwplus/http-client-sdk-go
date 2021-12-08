package client

import (
	"net"
	"time"
)

const defaultTimeout = 5 * time.Second

func dialContext(network, address string) (net.Conn, error) {
	return net.DialTimeout(network, address, defaultTimeout)
}
