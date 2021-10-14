package main

import "fmt"

func main() {
	var fruits = []string{"apple", "grape", "banana", "melon"}
	fmt.Println(fruits[0])

	var newFruits = fruits[0:2]
	fmt.Println(newFruits)
	fmt.Println(fruits[:])
	fmt.Println(fruits[2:])
	fmt.Println(fruits[:2])

	//pembuktian reference
	fmt.Println(fruits) // [apple grape banana melon]

	var aFruits = fruits[0:3]
	fmt.Println(aFruits) // [apple grape banana]

	var bFruits = fruits[1:4]
	fmt.Println(bFruits) // [grape banana melon]

	var aaFruits = aFruits[1:2]
	fmt.Println(aaFruits) // [grape]

	var bbFruits = bFruits[0:1]
	fmt.Println(bbFruits) // [grape]

	// Buah "grape" diubah menjadi "pinnaple"
	bbFruits[0] = "pineaple"

	// len and cap
	fmt.Println(fruits)       // [apple pinnaple banana melon]
	fmt.Println(aFruits)      // [apple pinnaple banana]
	fmt.Println(bFruits)      // [pinnaple banana melon]
	fmt.Println(aaFruits)     // [pinnaple]
	fmt.Println(bbFruits)     // [pinnaple]
	fmt.Println(len(fruits))  // len: 4
	fmt.Println(cap(fruits))  // cap: 4
	fmt.Println(len(aFruits)) // len: 3
	fmt.Println(cap(aFruits)) // cap: 4
	fmt.Println(len(bFruits)) // len: 3
	fmt.Println(cap(bFruits)) // cap: 3

	//append
	var cFruits = append(fruits, "papaya")
	fmt.Println(cFruits)

	//copy
	dst := make([]string, 3)
	src := []string{"watermelon", "pinnaple", "apple", "orange"}
	n := copy(dst, src)
	fmt.Println(dst) // watermelon pinnaple apple
	fmt.Println(src) // watermelon pinnaple apple orange
	fmt.Println(n)   // 3

	dst = []string{"potato", "potato", "potato"}
	src = []string{"watermelon", "pinnaple"}
	n = copy(dst, src)

	fmt.Println(dst) // watermelon pinnaple potato
	fmt.Println(src) // watermelon pinnaple
	fmt.Println(n)   // 2

	//last
	var fruits2 = []string{"apple", "grape", "banana"}
	var aFruits2 = fruits2[0:2]
	var bFruits2 = fruits2[0:2:2]
	fmt.Println(fruits2)       // ["apple", "grape", "banana"]
	fmt.Println(len(fruits2))  // len: 3
	fmt.Println(cap(fruits2))  // cap: 3
	fmt.Println(aFruits2)      // ["apple", "grape"]
	fmt.Println(len(aFruits2)) // len: 2
	fmt.Println(cap(aFruits2)) // cap: 3
	fmt.Println(bFruits2)      // ["apple", "grape"]
	fmt.Println(len(bFruits2)) // len: 2
	fmt.Println(cap(bFruits2)) // cap: 2
}
