package ping

import (
    "fmt"
    "net/http"
    "time"
)

type Pinger interface {
    Ping() (*Result, error)
}

type Result struct {
    Url      string
    Duration time.Duration
}

type Target struct {
    url string
}

func NewTarget(url string) *Target {
    return &Target{url}
}

func (t *Target) Ping() (*Result, error) {
    startTime := time.Now()
    res, err := http.Get(t.url)
    duration := time.Since(startTime)

    if err != nil {
        return nil, err
    }
    defer res.Body.Close()

    if res.StatusCode != http.StatusOK {
        return nil, fmt.Errorf("ping error from %s bad status: %d", t.url, res.StatusCode)
    }
    return &Result{t.url, duration}, nil
}
