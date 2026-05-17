package config

import (
	"bufio"
	"flag"
	"os"
)

type Config struct {
	Workers int
	File    string
	Timeout int
	URLs    []string
}

func MustLoad() Config {
	var cfg Config

	flag.IntVar(&cfg.Workers, "w", 5, "number of workers")
	flag.StringVar(&cfg.File, "f", "", "file with urls")
	flag.IntVar(&cfg.Timeout, "t", 5, "timeout in seconds")

	flag.Parse()

	cfg.URLs = flag.Args()

	return cfg
}

func LoadURLs(cfg Config) ([]string, error) {
	if cfg.File == "" {
		return cfg.URLs, nil
	}

	file, err := os.Open(cfg.File)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var urls []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		urls = append(urls, scanner.Text())
	}

	return urls, scanner.Err()
}
