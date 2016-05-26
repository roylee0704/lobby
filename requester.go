package lobby

import (
	"fmt"
	"math/rand"
)

// Request is a job item to be sent to worker.
type Request struct {
	fn func() int // unit of work
	c  chan int   // a channel to report job result.
}

// Requester creates job request (every s interval) to be sent to its receiver
// via request-channel.
//
// note:
// - new job only will be created upon job completion.
// - ultimately, a go-routine.
func Requester(work chan Request) {
	fmt.Print(rand.Int63n(2e9))
}
