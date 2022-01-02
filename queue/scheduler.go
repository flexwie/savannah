package queue

import (
	"log"
	"time"

	"felixwie.com/savannah/queue/worker"
)

type Scheduler struct {
	tickers []Ticker
}

type Ticker struct {
	ticker *time.Ticker
	quit   chan bool
}

func (s *Scheduler) Every(job *worker.PollingJob, interval time.Duration) error {
	t := time.NewTicker(interval)

	ticker := &Ticker{
		ticker: t,
		quit:   make(chan bool),
	}

	s.tickers = append(s.tickers, *ticker)

	go func() {
		for {
			select {
			case <-t.C:
				queue.Submit(job)
			case <-ticker.quit:
				log.Printf("quiting scheduled worker for %s", job.ID)
				return
			}
		}
	}()

	return nil
}
