package lobby

// Worker is an ADT, an item in priority queue, it performs dispatched job,
// and eports back when it's completed.
type Worker struct {
	i       int // index to be used by heap.interface.
	pending int // as a measure of load
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
