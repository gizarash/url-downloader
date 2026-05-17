package downloader

import (
	"io"
	"net/http"
	"time"
)

type HTTPClient struct {
	client *http.Client
}

func NewHTTPClient(timeout int) *HTTPClient {
	return &HTTPClient{
		client: &http.Client{
			Timeout: time.Duration(timeout) * time.Second,
		},
	}
}

func (c *HTTPClient) Fetch(url string) (int, int, error) {
	resp, err := c.client.Get(url)
	if err != nil {
		return 0, 0, err
	}
	defer resp.Body.Close()

	size, err := io.Copy(io.Discard, resp.Body)
	if err != nil {
		return 0, 0, err
	}

	return resp.StatusCode, int(size), nil
}
