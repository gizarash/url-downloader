package main

import (
	"fmt"
	"log"

	"github.com/gizarash/url-downloader/internal/config"
	"github.com/gizarash/url-downloader/internal/downloader"
)

func main() {
	// 1. Парсим конфиг (флаги CLI)
	cfg := config.MustLoad()

	// 2. Получаем список URL
	urls, err := config.LoadURLs(cfg)
	if err != nil {
		log.Fatalf("failed to load urls: %v", err)
	}

	if len(urls) == 0 {
		log.Fatal("no urls provided")
	}

	// 3. Создаём downloader
	d := downloader.New(cfg)

	// 4. Запускаем обработку
	results := d.Run(urls)

	// 5. Выводим результат
	for res := range results {
		if res.Err != nil {
			fmt.Printf("[ERR] %s %v\n", res.URL, res.Err)
			continue
		}

		fmt.Printf("[OK] %s %d %dB %v\n",
			res.URL,
			res.StatusCode,
			res.Size,
			res.Duration,
		)
	}
}