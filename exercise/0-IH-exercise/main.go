package main

import (
	"fmt"
	"runtime"
	"time"
)


func main() {
	// use channel

	runtime.GOMAXPROCS(2)

	timeStart := time.Now()

	var messages = make(chan string)

	go taskOne(3, messages)
	go taskTwo(2, messages)
	go taskThree(5, messages)

	var task1 = <-messages
	fmt.Println(task1)

	var task2 = <-messages
	fmt.Println(task2)

	var task3 = <-messages
	fmt.Println(task3)

	fmt.Println(time.Since(timeStart))
}

func taskOne(till int, messages chan string) {
	for i:= 0; i < till; i++ {
		time.Sleep(time.Second * 1)
	}
	messages <- fmt.Sprintf("Task 1 with %d second", till)
}

func taskTwo(till int, messages chan string) {
	for i:= 0; i < till; i++ {
		time.Sleep(time.Second * 1)
	}
	messages <- fmt.Sprintf("Task 2 with %d second", till)
}

func taskThree(till int, messages chan string) {
	for i:= 0; i < till; i++ {
		time.Sleep(time.Second * 1)
	}
	messages <- fmt.Sprintf("Task 3 with %d second", till)
}