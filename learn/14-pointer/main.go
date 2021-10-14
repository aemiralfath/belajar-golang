package main

import "fmt"

func main() {

	//version 1
	var numberA int = 4
	var numberB *int = &numberA
	fmt.Println("numberA (value)   :", numberA)  // 4
	fmt.Println("numberA (address) :", &numberA) // 0xc20800a220
	fmt.Println("numberB (value)   :", *numberB) // 4
	fmt.Println("numberB (address) :", numberB)  // 0xc20800a220
	fmt.Println("")

	//version 2
	numberA = 5
	fmt.Println("numberA (value)   :", numberA)  // 5
	fmt.Println("numberA (address) :", &numberA) // 0xc20800a220
	fmt.Println("numberB (value)   :", *numberB) // 5
	fmt.Println("numberB (address) :", numberB)  // 0xc20800a220

	//version 3
	var number = 4
	fmt.Println("before", number) // 4
	change(&number, 10)
	fmt.Println("after", number) // 10
}

func change(original *int, value int) {
	*original = value
}
