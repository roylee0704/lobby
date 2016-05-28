package lobby

import (
	"bytes"
	"container/heap"
	"fmt"
)

// Worker is an ADT, an item in priority queue, it performs dispatched job,
// and eports back when it's completed.
type Worker struct {
	i        int          // index to be used by heap.interface.
	pending  int          // as a measure of load
	requests chan Request // worker has own job queues to be passed into.
}

// work is a [go-routine] that performs any given job function, one at a time,
// upon job completion, it reports back to balancer and
// requester (request-response model).
func (w *Worker) work(done chan *Worker) {
	for r := range w.requests {
		result := r.Fn() // perform job function
		r.C <- result    // response end result to requester
		done <- w        // report job completion to balancer
	}
}

// Pool implements heap.interface, a piority queue based on workload.
type Pool []*Worker

// Len returns number of workers in pool
func (p Pool) Len() int {
	return len(p)
}

// Less ensures min-heap property where lightly-loaded worker stays up.
func (p Pool) Less(i, j int) bool {
	return p[i].pending < p[j].pending
}

// Swap performs in-place ordering between workers in queue
func (p Pool) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
	p[i].i = i
	p[j].i = j
}

// Push inserts *worker in the back of queue, heapify might occur internally.
func (p *Pool) Push(x interface{}) {
	a := *p
	*p = append(a, x.(*Worker)) // in event of malloc within append()
}

// Pop remove last item in the queue (originally from root)
func (p *Pool) Pop() interface{} {
	old := *p
	n := len(old)
	x := old[n-1]
	*p = old[0 : n-1] // in event of len(slice) update.

	return x
}

// Balancer maintains a pool of workers, and a channel for worker to report job
// completion.
type Balancer struct {
	pool Pool
	done chan *Worker
}

// NewBalancer initializes a pool of n workers.Each worker be given a
// `walkie-talkie` before work, so they could report back upon job completion.
func NewBalancer(nWorker int, nRequester int) *Balancer {

	workers := make(Pool, 0, nWorker)
	done := make(chan *Worker)
	for i := 0; i <= nWorker; i++ {
		w := &Worker{
			requests: make(chan Request, nRequester),
			pending:  0,
		}
		go w.work(done) // get ready to work
		heap.Push(&workers, w)
	}

	return &Balancer{
		pool: workers,
		done: done,
	}
}

// Balance [go-routine] multiplexes job requests and job completions
func (b Balancer) Balance(work chan Request) {
	for {
		select {
		case r := <-work:
			b.dispatch(r)
		case w := <-b.done:
			b.complete(w)
		}
	}
}

// String outputs balancer in string
func (b Balancer) String() string {
	var buf bytes.Buffer

	for _, w := range b.pool {
		buf.WriteString(fmt.Sprintf("%2d", w.pending))
	}
	return buf.String()
}

func (b Balancer) dispatch(r Request) {
	w := heap.Pop(&b.pool).(*Worker)
	w.pending++
	w.requests <- r
	heap.Push(&b.pool, w)

	fmt.Println(b)
}

func (b Balancer) complete(w *Worker) {

	w = heap.Remove(&b.pool, w.i).(*Worker)
	w.pending--
	heap.Push(&b.pool, w)
	fmt.Println(b)

}
