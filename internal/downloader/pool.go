package downloader

import (
	"sync"

    "github.com/gizarash/url-downloader/internal/config"
    "github.com/gizarash/url-downloader/internal/model"
)

type Downloader struct {
	cfg    config.Config
	client *HTTPClient
}

func New(cfg config.Config) *Downloader {
	return &Downloader{
		cfg:    cfg,
		client: NewHTTPClient(cfg.Timeout),
	}
}

func (d *Downloader) Run(urls []string) <-chan model.Result {
	jobs := make(chan model.Job)
	results := make(chan model.Result)

	var wg sync.WaitGroup

	// 1. Запуск воркеров
	for i := 0; i < d.cfg.Workers; i++ {
		wg.Add(1)

		go func(id int) {
			defer wg.Done()
			worker(id, jobs, results, d.client)
		}(i)
	}

	// 2. Отправка задач
	go func() {
		for _, url := range urls {
			jobs <- model.Job{URL: url}
		}
		close(jobs)
	}()

	// 3. Закрытие results после завершения воркеров
	go func() {
		wg.Wait()
		close(results)
	}()

	return results
}