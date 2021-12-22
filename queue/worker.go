package queue

import (
	"fmt"
	"time"
)

type WebhookJob struct {
	ID string
}

func (t *WebhookJob) Process() {
	fmt.Printf("processing: %s \n", t.ID)
	time.Sleep(1 * time.Second)
}
