package main

import (
    "sync"

    "pinger/ping"
)

func main() {
    targets := []string{
        "http://google.com",
        "http://puppetlabs.com",
        "http://newrelic.com",
    }
    work := make(chan ping.Pinger)
    result := make(chan *ping.Result)

    var wg sync.WaitGroup
    wg.Add(2)
    go pinger(work, result, &wg)
    go printer(result, &wg)

    for _, t := range targets {
        work <- ping.NewTarget(t)
    }
    close(work)
    wg.Wait()
}
