package model

import "time"

type Job struct {
	URL string
}

type Result struct {
	URL        string
	StatusCode int
	Size       int
	Duration   time.Duration
	Err        error
}