package client

import (
	"fmt"
	"net/http"
)

func CheckRedirect(req *http.Request, via []*http.Request) error {
	if via[0].Method == http.MethodGet {
		return http.ErrUseLastResponse
	}
	return fmt.Errorf("non-GET Method unexpected redirect in response")
}
