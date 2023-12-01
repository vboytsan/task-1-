package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

func main() {

	message := make(chan string)
	go func() {
		message <- "Hello World"
	}()
	msg := <-message
	fmt.Println(msg)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	done := make(chan bool)
	go func() {
		select {
		case <-time.After(10 * time.Second):
			fmt.Println("Goodbye World!")
			done <- true
		case <-interrupt:
			fmt.Println("Stopped by the user after x seconds")
			done <- true
		}
	}()
	<-done
}
