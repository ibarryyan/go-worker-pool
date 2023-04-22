package internal

// Step 1: Define the Task
// A task that accepts a URL and returns the extracted data as a string.
type Task func(url string) (string, error)

// Step 2: Create the Worker
// A worker is a goroutine that processes tasks and sends the results through a channel.
type Worker struct {
	id         int
	task       Task
	taskQueue  <-chan string
	resultChan chan<- Result
}

func (w *Worker) Start() {
	go func() {
		for url := range w.taskQueue {
			data, err := w.task(url) // Perform the web scraping task
			w.resultChan <- Result{workerID: w.id, url: url, data: data, err: err}
		}
	}()
}

// Step 3: Implement the Worker Pool
// The worker pool manages the workers, distributes tasks, and collects results.
type WorkerPool struct {
	taskQueue   chan string
	resultChan  chan Result
	task        Task
	workerCount int
}

type Result struct {
	workerID int
	url      string
	data     string
	err      error
}

func NewWorkerPool(workerCount int, task Task) *WorkerPool {
	return &WorkerPool{
		taskQueue:   make(chan string),
		resultChan:  make(chan Result),
		task:        task,
		workerCount: workerCount,
	}
}

func (wp *WorkerPool) Start() {
	for i := 0; i < wp.workerCount; i++ {
		worker := Worker{id: i, taskQueue: wp.taskQueue, resultChan: wp.resultChan}
		worker.Start()
	}
}

func (wp *WorkerPool) Submit(url string) {
	wp.taskQueue <- url
}

func (wp *WorkerPool) GetResult() Result {
	return <-wp.resultChan
}
