package main

import (
	"fmt"
	"time"
)

func main() {

	message := make(chan string, 3)

	go func() {
		message <- "send"
		fmt.Println("go func")
	}()
	fmt.Println("main func")
	str := <-message
	fmt.Println(str)

	time.Sleep(time.Nanosecond * 30000)
}
