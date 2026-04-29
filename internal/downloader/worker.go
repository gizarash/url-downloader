package downloader

import (
	"time"

	"github.com/gizarash/url-downloader/internal/model"
)

func worker(
	id int,
	jobs <-chan model.Job,
	results chan<- model.Result,
	client *HTTPClient,
) {
	for job := range jobs {
		start := time.Now()

		// TODO:
		// вызвать client.Fetch
		// обработать ошибку

		res := model.Result{
			URL: job.URL,
			// TODO: заполнить поля
			Duration: time.Since(start),
		}

		results <- res
	}
}