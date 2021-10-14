package main

import "fmt"

func main() {

	// version 1
	var names [3]string
	names[0] = "Ahmad"
	names[1] = "Emir"
	names[2] = "Alfatah"
	// names[3] = "hehe" error
	fmt.Println(names[0], names[1], names[2])

	//version 2
	var fruitsV1 = [4]string{"apple", "grape", "banana", "melon"}
	fmt.Println("Jumlah element \t\t", len(fruitsV1))
	fmt.Println("Isi semua element \t", fruitsV1)

	// version 3
	var fruitsV2 = [4]string{
		"apple",
		"grape",
		"banana",
		"melon",
	}
	fmt.Println("Jumlah element \t\t", len(fruitsV2))

	//version 4
	var numbers = [...]int{2, 3, 2, 4, 3}
	fmt.Println("data array \t:", numbers)
	fmt.Println("jumlah elemen \t:", len(numbers))

	//version 5
	var numbers2 = [2][3]int{{3, 2, 3}, {3, 4, 5}}
	fmt.Println("numbers2", numbers2)

	//version 6
	//var fruits = [4]string{"apple", "grape", "banana", "melon"}
	for i := 0; i < len(fruitsV1); i++ {
		fmt.Printf("elemen %d : %s\n", i, fruitsV1[i])
	}
	for i, fruit := range fruitsV1 {
		fmt.Printf("elemen %d : %s\n", i, fruit)
	}
	for _, fruit := range fruitsV1 {
		fmt.Printf("nama buah : %s\n", fruit)
	}
	// for i, _ := range fruitsV1 { }
	// for i := range fruitsV1 { }

	// 7
	var fruits2 = make([]string, 2)
	fruits2[0] = "apple"
	fruits2[1] = "manggo"
	fmt.Println(fruits2)
}
