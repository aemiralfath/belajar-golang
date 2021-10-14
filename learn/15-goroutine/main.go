package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(2)

	var messages = make(chan string)

	go print(5, "concurrent", messages)
	go print(10, "sequential", messages)
	go print(15, "parallel", messages)

	var message1 = <-messages
	fmt.Println(message1)

	var message2 = <-messages
	fmt.Println(message2)

	var message3 = <-messages
	fmt.Println(message3)
}

func print(till int, message string, messages chan string) {
	result := ""
	for i := 0; i < till; i++ {
		result += fmt.Sprintf("hello %d %s \n", (i + 1), message)
	}
	messages <- result
}
