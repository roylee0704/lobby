package lobby

import (
	"math/rand"
	"time"
)

// Request is a job item to be sent to worker.
type Request struct {
	Fn func() int // unit of work
	C  chan int   // a channel to report job result.
}

// Requester creates job request (every s interval) to be sent to its receiver
// via request-channel.
//
// note:
// - new job only will be created upon job completion.
// - this is a go-routine.
func Requester(workFn func() int, work chan<- Request) {
	for {

		time.Sleep(time.Duration(rand.Int63n(2 * 2e9))) // spend time
		c := make(chan int)                             //
		work <- Request{Fn: workFn, C: c}               // send work via chan

		<-c // wait for job completion
		//fmt.Println("job done!")
	}
}
