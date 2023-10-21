package main

import (
  "time"
  "github.com/robfig/cron/v3"
)

func main() {
	loc, err := time.LoadLocation("America/Sao_Paulo")
  if err != nil {
    panic(err)
  }

  cronJob := cron.WithLocation(loc)

  cronJob.AddFunc("* * * * * *", func() {
  })
}
