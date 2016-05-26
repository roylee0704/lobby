package main

import (
	"fmt"

	"github.com/roylee0704/lobby"
)

func main() {

	work := make(chan lobby.Request)

	f := func() int {
		fmt.Println("working on job request!")
		return 1
	}

	go lobby.Requester(f, work)

	func(req chan lobby.Request) {
		for {
			select {
			case r := <-req:
				r.C <- r.Fn()
			}
		}

	}(work)
}
