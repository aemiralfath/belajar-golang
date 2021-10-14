package main

import (
	"fmt"
	"strings"
)

func main() {
	// version 1
	var avgV1 = calculate2(2, 4, 3, 5, 4, 3, 3, 5, 5, 3)
	//mengembalikan nilai dalam bentuk string (fmt.Sprint() dan fmt.Sprintln())
	var msgV1 = fmt.Sprintf("Rata-rata : %.2f", avgV1)
	fmt.Println(msgV1)

	// version 2
	var numbers = []int{2, 4, 3, 5, 4, 3, 3, 5, 5, 3}
	var avgV2 = calculate2(numbers...)
	var msgV2 = fmt.Sprintf("Rata-rata : %.2f", avgV2)
	fmt.Println(msgV2)

	// version 3
	var hobbies = []string{"sleeping", "eating"}
	yourHobbies("emir", "code", "bucin")
	yourHobbies("emir", hobbies...)
}

//fungsi variadic
func yourHobbies(name string, hobbies ...string) {
	var hobbiesAsString = strings.Join(hobbies, ", ")
	fmt.Printf("Hello, my name is: %s\n", name)
	fmt.Printf("My hobbies are: %s\n", hobbiesAsString)
}

func calculate2(numbers ...int) float64 {
	var total int = 0
	for _, number := range numbers {
		total += number
	}
	var avg = float64(total) / float64(len(numbers))
	return avg
}
