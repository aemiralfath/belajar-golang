package main

import "fmt"
import "strings"

type FilterCallback func(string) bool

func main() {

	var data = []string{"Ahmad", "Emir", "Alfatah", "owo"}
	fmt.Println("data asli \t\t:", data)

	//version 1
	var dataContainsO = filter(data, func(each string) bool {
		return strings.Contains(each, "o")
	})
	fmt.Println("filter ada huruf \"o\"\t:", dataContainsO)

	//version 2
	var dataLenght5 = filter2(data, func(each string) bool {
		return len(each) == 5
	})
	fmt.Println("filter jumlah huruf \"5\"\t:", dataLenght5)

}

func filter(data []string, callback func(string) bool) []string {
	var result []string
	for _, each := range data {
		if filtered := callback(each); filtered {
			result = append(result, each)
		}
	}
	return result
}

func filter2(data []string, callback FilterCallback) []string {
	var result []string
	for _, each := range data {
		if filtered := callback(each); filtered {
			result = append(result, each)
		}
	}
	return result
}
