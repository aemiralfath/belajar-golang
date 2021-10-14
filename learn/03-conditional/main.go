package main

import "fmt"

func main() {

	// version 1
	var point = 8
	if point == 10 {
		fmt.Println("Lulus dengan nilai sempurrna")
	} else if point > 5 {
		fmt.Println("Lulus")
	} else if point == 4 {
		fmt.Println("Hampir lulus")
	} else {
		fmt.Printf("Tidak lulus. nilai anda %d\n", point)
	}

	// version 2
	var point2 = 8840.0
	if percent := point2 / 100; percent >= 100 {
		fmt.Printf("%.1f%s perfect!\n", percent, "%")
	} else if percent >= 70 {
		fmt.Printf("%.1f%s good\n", percent, "%")
	} else {
		fmt.Printf("%.1f%s not bad\n", percent, "%")
	}

	// version 3
	var point3 = 6
	switch { // point3
	case point3 == 8: // 8
		fmt.Println("Perfect")
	case (point3 < 8) && (point3 > 3): // 7, 6, 5 ,4
		fmt.Println("Awesome")
		fallthrough //untuk meneruskan
	case point3 < 5:
		fmt.Println("you need to learn more")
	default:
		{
			fmt.Println("Not bad")
			fmt.Println("you can be better!")
		}
	}

	// version 4
	var point4 = 10
	if point > 7 {
		switch point4 {
		case 10:
			fmt.Println("perfect!")
		default:
			fmt.Println("nice!")
		}
	} else {
		if point == 5 {
			fmt.Println("not bad")
		} else if point == 3 {
			fmt.Println("kepp trying")
		} else {
			fmt.Println("you can do it")
			if point == 0 {
				fmt.Println("try harder!")
			}
		}
	}

}
