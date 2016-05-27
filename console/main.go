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
	lobby.NewBalancer(5, 50).Balance(work)
}
