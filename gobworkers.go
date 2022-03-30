package gobworkers

type WorkerPool interface {
	AddWork(do func())
	WaitForWorkToBeDone()
}

type workerpool struct {
	jobChan         chan func()
	doneChan        chan bool
	amountOfWorkers int
}

func New(amountOfWorkers int) WorkerPool {
	jobChan := make(chan func(), amountOfWorkers)
	doneChan := make(chan bool)
	wp := &workerpool{
		amountOfWorkers: amountOfWorkers,
		jobChan:         jobChan,
		doneChan:        doneChan,
	}

	for i := 0; i < amountOfWorkers; i++ {
		go worker{
			jobChan:  jobChan,
			doneChan: doneChan,
		}.Start()
	}

	return wp
}

func (wp *workerpool) AddWork(do func()) {
	wp.jobChan <- do
}

func (wp *workerpool) WaitForWorkToBeDone() {
	for {
		if len(wp.jobChan) == 0 {
			close(wp.jobChan)
			break
		}
	}

	for i := 0; i < wp.amountOfWorkers; i++ {
		<-wp.doneChan
	}

}

type worker struct {
	jobChan  <-chan func()
	doneChan chan<- bool
}

func (w worker) Start() {
	defer func() {
		w.doneChan <- true
	}()

	for workReceived := range w.jobChan {
		workReceived()
	}
}
