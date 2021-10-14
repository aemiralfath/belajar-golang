package main

import "fmt"

func main() {

	// version 1
	var chicken map[string]int
	chicken = map[string]int{}
	chicken["januari"] = 50
	chicken["februari"] = 40

	chicken = map[string]int{
		"januari":  50,
		"februari": 40,
		"maret":    34,
		"april":    67,
	}

	fmt.Println("januari", chicken["januari"])
	fmt.Println("mei", chicken["mei"])

	for key, val := range chicken {
		fmt.Println(key, "  \t:", val)
	}

	//version 2
	var chickenV2 = map[string]int{"januari": 50, "februari": 40}
	fmt.Println(len(chickenV2))
	fmt.Println(chickenV2)

	//delete()
	delete(chickenV2, "januari")
	fmt.Println(len(chickenV2))
	fmt.Println(chickenV2)

	//isexist
	var value, isExist = chickenV2["mei"]
	if isExist {
		fmt.Println(value)
	} else {
		fmt.Println("item is not exist")
	}

	//map and slice
	var chickens = []map[string]string{
		{"name": "chicken blue", "gender": "male", "color": "blue"},
		{"name": "chicken red", "gender": "male"},
		{"name": "chicken yellow", "gender": "female"},
	}
	for _, chicken := range chickens {
		fmt.Println(chicken["gender"], chicken["name"])
	}

	// var chicken4 = make(map[string]int)
	// var chicken5 = *new(map[string]int)
}
