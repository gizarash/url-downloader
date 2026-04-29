package downloader

import (
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
	// TODO:
	// 1. сделать GET запрос
	// 2. обработать ошибку
	// 3. не забыть закрыть body
	// 4. посчитать размер ответа

	return 0, 0, nil
}