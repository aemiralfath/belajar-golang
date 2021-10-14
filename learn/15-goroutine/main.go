package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(2)
	
	timeStart := time.Now()

	var messages = make(chan string)

	go print(3, "concurrent", messages)
	go print(2, "sequential", messages)
	go print(5, "parallel", messages)

	var message1 = <-messages
	fmt.Println(message1)

	var message2 = <-messages
	fmt.Println(message2)

	var message3 = <-messages
	fmt.Println(message3)

	fmt.Println(time.Since(timeStart))
}

func print(till int, message string, messages chan string) {
	result := ""
	for i := 0; i < till; i++ {
		result += fmt.Sprintf("hello %d %s \n", (i + 1), message)
		time.Sleep(time.Second * 1)
	}
	messages <- result
}
