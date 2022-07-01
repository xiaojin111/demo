package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/gammazero/workerpool"
)

func TestAAA(t *testing.T) {
	wp := workerpool.New(2)
	requests := []string{"alpha", "beta", "gamma", "delta", "epsilon"}

	for _, r := range requests {
		r := r
		wp.Submit(func() {
			time.Sleep(1*time.Second)
			fmt.Println("Handling request:", r)
		})
	}

	wp.StopWait()

}
