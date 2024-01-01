package main

import (
	"fmt"
	"time"
)

func timer() {
	timer1 := time.NewTimer(2 * time.Second)
	<-timer1.C
	fmt.Println("Timer One is fired")

	timer2 := time.NewTimer(time.Second)

	go func() {
		<-timer2.C
		fmt.Println("2nd Timer is fired")
	}()

	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("2nd timer is stop")
	}

	time.Sleep(2 * time.Second)

}
