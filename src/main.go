package main

import (
	"fmt"

	"github.com/gammazero/workerpool"
)

func main() {
	wp := workerpool.New(3)
	requests := []string{"alpha", "beta", "gamma", "delta", "epsilon"}

	for _, r := range requests {
		robj := r
		wp.Submit(func() {
			fmt.Println("Handling request: ", robj)
		})
	}
	wp.StopWait()

}
