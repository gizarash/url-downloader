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

		statusCode, contentLenght, err := client.Fetch(job.URL)

		res := model.Result{
			URL:        job.URL,
			StatusCode: statusCode,
			Size:       contentLenght,
			Duration:   time.Since(start),
			Err:        err,
		}

		results <- res
	}
}
