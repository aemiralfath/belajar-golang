package main

import (
	"fmt"
)

func main() {
	//version 1
	var getMinMax = func(n []int) (min int, max int) {
		for i, e := range n {
			switch {
			case i == 0:
				max, min = e, e
			case e > max:
				max = e
			case e < min:
				min = e
			}
		}
		return
	}

	var numbers = []int{2, 3, 4, 3, 4, 2, 3}
	var min, max = getMinMax(numbers)
	//%v untuk menampilkan segala jenis data
	fmt.Printf("data : %v\nmin  : %v\nmax  : %v\n", numbers, min, max)


	//version 2
	var numbers2 = []int{2, 3, 0, 4, 3, 2, 0, 4, 2, 0, 3}
	var newNumbers = func(min int) []int {
		var r []int
		for _, e := range numbers2 {
			if e < min {
				continue
			}
			r = append(r, e)
		}
		return r
	}(3)

	fmt.Println("original number :", numbers)
	fmt.Println("filtered number :", newNumbers)


	//version 3
	var max2 = 3
	var howMany, getNumbers = findMax(numbers2, max2)
	var theNumbers = getNumbers()
	fmt.Println("numbers\t:", numbers2)
	fmt.Printf("find \t: %d\n\n", max2)
	fmt.Println("found \t:", howMany)    // 9
	fmt.Println("value \t:", theNumbers) // [2 3 0 3 2 0 2 0 3]
}

//closure
func findMax(numbers []int, max int) (int, func() []int) {
	var res []int
	for _, e := range numbers {
		if e <= max {
			res = append(res, e)
		}
	}
	return len(res), func() []int {
		return res
	}
}
