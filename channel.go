package main

import "fmt"

func channel() {
	message := make(chan string)
	go func() {
		message <- "Ping"
	}()

	msg := <-message

	fmt.Println("Channel Passs Message: ", msg)

}
