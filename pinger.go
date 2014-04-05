package main

import (
    "log"
    "sync"

    "pinger/ping"
)

func pinger(work chan ping.Pinger, result chan *ping.Result, wg *sync.WaitGroup) {
    for {
        w, ok := <-work
        if !ok {
            break
        }
        res, err := w.Ping()
        if err != nil {
            log.Printf("pinger: %s", err)
            continue
        }
        result <- res
    }
    close(result)
    wg.Done()
}

func printer(result chan *ping.Result, wg *sync.WaitGroup) {
    for {
        res, ok := <-result
        if !ok {
            break
        }
        log.Printf("ping %s %s", res.Url, res.Duration)
    }
    wg.Done()
}
