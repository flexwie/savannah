package queue

import (
	"log"
	"sync"
	"time"

	"felixwie.com/savannah/config"
	"felixwie.com/savannah/queue/worker"
	"github.com/google/uuid"
)

type Job interface {
	Process()
}

type Worker struct {
	done             sync.WaitGroup
	readyPool        chan chan Job
	assignedJobQueue chan Job

	quit chan bool
}

type JobQueue struct {
	internalQueue     chan Job
	readyPool         chan chan Job
	workers           []*Worker
	dispatcherStopped sync.WaitGroup
	workersStopped    sync.WaitGroup

	quit chan bool
}

var queue *JobQueue
var scheduler Scheduler

func init() {
	queue = NewJobQueue(1)
	scheduler = Scheduler{}

	cfg := config.GetConfig().Source

	for _, s := range cfg {
		if s.Polling != nil {
			log.Printf("adding polling worker for %s", s.Name)
			scheduler.Every(&worker.PollingJob{
				ID:         uuid.NewString(),
				Repository: s.URL,
				Branch:     s.Branch,
				Folder:     s.Folder,
			}, time.Duration(s.Polling.Interval)*time.Second)
		}
	}
}

func GetQueue() *JobQueue {
	return queue
}

// QUEUES
func (q *JobQueue) Start() {
	for i := 0; i < len(q.workers); i++ {
		q.workers[i].Start()
	}
	go q.dispatch()
}

func (q *JobQueue) Submit(job Job) {
	q.internalQueue <- job
}

func (q *JobQueue) Stop() {
	q.quit <- true
	q.dispatcherStopped.Wait()
}
func (q *JobQueue) dispatch() {
	q.dispatcherStopped.Add(1)

	for {
		select {
		case job := <-q.internalQueue:
			workerChannel := <-q.readyPool
			workerChannel <- job
		case <-q.quit:
			for i := 0; i < len(q.workers); i++ {
				q.workers[i].Stop()
			}
			q.workersStopped.Wait()
			q.dispatcherStopped.Done()
			return
		}
	}
}

// WORKERS
func (w *Worker) Start() {
	go func() {
		w.done.Add(1)
		for {
			w.readyPool <- w.assignedJobQueue

			select {
			case job := <-w.assignedJobQueue:
				job.Process()
			case <-w.quit:
				w.done.Done()
				return
			}
		}
	}()
}

func (w *Worker) Stop() {
	w.quit <- true
}

func NewJobQueue(maxWorkers int) *JobQueue {
	workersStopped := sync.WaitGroup{}
	readyPool := make(chan chan Job, maxWorkers)
	workers := make([]*Worker, maxWorkers)

	for i := 0; i < maxWorkers; i++ {
		workers[i] = NewWorker(readyPool, workersStopped)
	}

	return &JobQueue{
		internalQueue:     make(chan Job),
		readyPool:         readyPool,
		workers:           workers,
		dispatcherStopped: sync.WaitGroup{},
		workersStopped:    workersStopped,
		quit:              make(chan bool),
	}
}

func NewWorker(readyPool chan chan Job, done sync.WaitGroup) *Worker {
	return &Worker{
		done:             done,
		readyPool:        readyPool,
		assignedJobQueue: make(chan Job),
		quit:             make(chan bool),
	}
}
