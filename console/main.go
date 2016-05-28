package main

import (
	"flag"
	"math/rand"
	"time"

	"github.com/roylee0704/lobby"
)

var (
	nRequester = 0
	nWorker    = 0
)

func main() {

	flag.IntVar(&nRequester, "r", 50, "Number of requesters")
	flag.IntVar(&nWorker, "w", 10, "Number of workers")

	work := make(chan lobby.Request)

	f := func() int {
		time.Sleep(time.Duration(rand.Int63n(5e9)))
		return 1
	}

	for i := 0; i < nRequester; i++ {
		go lobby.Requester(f, work) // one worker only.
	}

	lobby.NewBalancer(nWorker, nRequester).Balance(work)

}
