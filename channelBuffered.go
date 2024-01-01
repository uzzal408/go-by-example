package main

import "fmt"

func channelBuffered() {
	messages := make(chan string, 2)

	messages <- "Buffered"
	messages <- "Channel"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
