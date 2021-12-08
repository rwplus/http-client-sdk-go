package client

import (
	"io"
	"log"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestBasicAuthClient(t *testing.T) {
	client, err := NewClientWithOpts(Options{
		BasicAuth: &BasicAuthOption{
			Username: "weekndCN",
			// github using personal token
			Password: os.Getenv("GITHUB_PERSONAL_TOKEN"),
		},
		HTTPTransport: &TransportOption{
			DialContextTimeout: 30 * time.Second,
		},
		Middlewares: []Middleware{
			MiddlewareFunc(BasicAuthenticationMiddleware),
		},
	})

	assert.NoError(t, err)
	res, err := client.Get("https://api.github.com/user")
	assert.NoError(t, err)
	defer res.Body.Close()
	assert.Equal(t, res.StatusCode/100, 2)
}

func TestUnixClient(t *testing.T) {
	client, err := NewClientWithOpts(Options{
		HTTPTransport: &TransportOption{
			Protocol:           "unix",
			Addr:               "/var/run/docker.sock",
			DialContextTimeout: 30 * time.Second,
		},
	})
	assert.NoError(t, err)
	resp, err := client.Get("http://localhost/v1.41/containers/json")
	assert.NoError(t, err)
	defer resp.Body.Close()
	res, err := io.ReadAll(resp.Body)
	assert.NoError(t, err)
	log.Println(string(res))
}
