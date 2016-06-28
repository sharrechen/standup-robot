package main

import (
    "github.com/robfig/cron"
    "jobs"
    "log"
)

func main() {
    i := 0
    c := cron.New()
    c.AddFunc("*/5 * * * * ?", func() {
        i++
        jobs.Standup("STAND UPPPP!! ", i)
    })
    c.Start()
    select{}
    // c.Stop()  // Stop the scheduler (does not stop any jobs already running).
}
