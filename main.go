package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	// Ticker to write a log every 100ms.
	ticker := time.NewTicker(100 * time.Millisecond)
	go func() {
		for t := range ticker.C {
			n := int(math.Round(float64(t.Nanosecond())/1000/1000/100)) + 1
			fmt.Println("Tick #", fmt.Sprintf("%02d", n), "in second")
		}
	}()

	fmt.Println("burn some CPU")
	go func() {
		for {
			math.Sqrt(1456789509685765673854965098574.0)
		}
	}()

	select {}
}
