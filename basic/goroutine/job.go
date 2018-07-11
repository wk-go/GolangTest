package main
//可控的go程并发实践
//需要注意的是worker的go程是需要阻塞的不然会出现fatal error: all goroutines are asleep - deadlock!错误
import (
	"log"
	"time"
	"os"
	"strconv"
)

////////////////////////////////////////////////////////////////////////////////
//Job接口
type Job interface {
	Do() error
}

// Global variables
//执行job(工人)
type Worker struct {
	Name		string
	WorkerPool chan *Worker
	JobChannel chan Job
	quit       chan bool
}

func NewWorker(name string, workerPool chan *Worker) *Worker {
	return &Worker{
		Name:name,
		WorkerPool: workerPool,
		JobChannel: make(chan Job),
		quit:       make(chan bool),
	}
}

func (w *Worker) Start() {
	go func() {
		log.Printf("worker[%s] starting\n", w.Name)
		for {
			w.WorkerPool <- w
			select {
			case job := <-w.JobChannel:
				if err := job.Do(); err != nil {
					log.Printf("excute job failed with err: %v", err)
				}
			case <-w.quit:
				log.Printf("worker[%s] stoped\n", w.Name)
				return
			//default://一定要阻塞不然会报错

			}
		}
	}()
}

func (w *Worker) AddJob(job Job) {
	w.JobChannel <- job
}

func (w *Worker) Stop() {
	w.quit <- true
}

//调度器(车间主任)
type Dispatcher struct {
	Name string
	WorkerPool        chan *Worker
	WorkerList			[]*Worker
	JobQueue          chan Job
	MaxWorkerPoolSize int
	MaxJobQueueSize   int
	quit              chan bool
}

func NewDispatcher(name string) *Dispatcher {
	return &Dispatcher{
		Name:name,
		MaxWorkerPoolSize: 1000,
		MaxJobQueueSize:   1000,
		quit:              make(chan bool),
	}
}

func (d *Dispatcher) Run() {
	d.WorkerPool = make(chan *Worker, d.MaxWorkerPoolSize)
	d.JobQueue = make(chan Job, d.MaxJobQueueSize)
	for i := 0; i < d.MaxWorkerPoolSize; i++ {
		worker := NewWorker("w"+strconv.Itoa(i),d.WorkerPool)
		d.WorkerList = append(d.WorkerList, worker)
		worker.Start()
	}
	go d.dispatch()
}

func (d *Dispatcher) dispatch() {
	for {
		select {
		case job := <-d.JobQueue:
			worker := <-d.WorkerPool
			worker.AddJob(job)
		case <-d.quit:
			log.Printf("Dispatcher[%s] stoped!\n",d.Name)
			return
		default:
			//log.Println("dipathcer waiting!")
			//time.Sleep(time.Millisecond*10)
		}
	}
}
func (d *Dispatcher) AddJob(job Job) {
	d.JobQueue <- job
}

func (d *Dispatcher) Stop() {
	go func() {
		for _, w := range d.WorkerList {
			w.Stop()
			time.Sleep(time.Millisecond * 10)
		}
		d.quit <- true
	}()
}

///////////////////////////////////////////////////////////////

/////////////////test add/////////////////
type Job1 struct {
	Name  string
	Count int
}

func (j *Job1) Do() error {
	j.Count++
	log.Printf("%v:%d", j.Name, j.Count)
	return nil
}

func main() {
	maxWorkerPoolSize := 6
	if len(os.Args) >1 {
		maxWorkerPoolSize,_ = strconv.Atoi(os.Args[1])
	}
	log.Println("maxWorkPoolSize:", maxWorkerPoolSize)

	dispatcher := NewDispatcher("default")
	//worker的队列的数量跟任务数有一定的比例，当前例子的情况4、5、6、7效果最佳,8/8反而不好，超过以后处理的计数反而变少了
	dispatcher.MaxWorkerPoolSize = maxWorkerPoolSize

	//job的队列要足够大
	dispatcher.MaxJobQueueSize = 100000
	dispatcher.Run()
	t1 := time.NewTimer(time.Millisecond * 1)
	t2 := time.NewTimer(time.Millisecond * 1)
	t3 := time.NewTimer(time.Millisecond * 1)
	t4 := time.NewTimer(time.Millisecond * 1)
	t5 := time.NewTimer(time.Millisecond * 1)
	t6 := time.NewTimer(time.Millisecond * 1)
	t7 := time.NewTimer(time.Millisecond * 1)
	t8 := time.NewTimer(time.Millisecond * 1)

	job1 := &Job1{Name: "job-obj-1", Count: 0,}
	job2 := &Job1{Name: "job-obj-2", Count: 0,}
	job3 := &Job1{Name: "job-obj-3", Count: 0,}
	job4 := &Job1{Name: "job-obj-4", Count: 0,}
	job5 := &Job1{Name: "job-obj-5", Count: 0,}
	job6 := &Job1{Name: "job-obj-6", Count: 0,}
	job7 := &Job1{Name: "job-obj-7", Count: 0,}
	job8 := &Job1{Name: "job-obj-8", Count: 0,}

	tX := time.NewTimer(time.Second * 1)
For:
	for {
		select {
		case <-t1.C:
			dispatcher.AddJob(job1)
			t1.Reset(time.Millisecond * 1)

		case <-t2.C:
			dispatcher.AddJob(job2)
			t2.Reset(time.Millisecond * 1)

		case <-t3.C:
			dispatcher.AddJob(job3)
			t3.Reset(time.Millisecond * 1)

		case <-t4.C:
			dispatcher.AddJob(job4)
			t4.Reset(time.Millisecond * 1)

		case <-t5.C:
			dispatcher.AddJob(job5)
			t5.Reset(time.Millisecond * 1)

		case <-t6.C:
			dispatcher.AddJob(job6)
			t6.Reset(time.Millisecond * 1)

		case <-t7.C:
			dispatcher.AddJob(job7)
			t7.Reset(time.Millisecond * 1)

		case <-t8.C:
			dispatcher.AddJob(job8)
			t8.Reset(time.Millisecond * 1)

		case <-tX.C:
			dispatcher.Stop()
			break For

		default:
			//log.Println("main waiting!")
			//time.Sleep(time.Millisecond*10)

		}
	}

	job1.Do()
	job2.Do()
	job3.Do()
	job4.Do()
	job5.Do()
	job6.Do()
	job7.Do()
	job8.Do()
	log.Println("maxWorkPoolSize:", maxWorkerPoolSize)
}
