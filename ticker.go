package main

import (
	"fmt"
	"time"
)

func ticker() {
	ticker := time.NewTicker(500 * time.Microsecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Run Ticker at : ", t)
			}
		}
	}()
	<-ticker.C
	fmt.Println("Running Ticker")

	time.Sleep(2600 * time.Microsecond)
	ticker.Stop()
	fmt.Println("ticker Stop")
}
