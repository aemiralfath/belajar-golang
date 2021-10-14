package main

import "fmt"

func main() {

	// version 1
	var firstNameV1 string = "Emir"
	lastNameV1 := "Alfatah"
	fmt.Printf("Halo %s %s!\n", firstNameV1, lastNameV1)

	// version 2
	var firstNameV2 = "Emir"
	fmt.Println(firstNameV2)

	// version 3
	var firstNameV3, lastNameV3 = "Emir", "Alfatah"
	fmt.Println("Halo", firstNameV3, lastNameV3+"!")

	name := new(string)
	fmt.Println(name)
	fmt.Println(*name)

	exist := true
	fmt.Printf("exits? %t \n", exist)

	message := `Nama saya "Ahmad Emir Alfatah".
	Teknik Informatika.
	Universitas Sriwijaya.`
	fmt.Println(message)

	var _ = "hehe"
	const phi = 3.14 //tidak bisa diubah
	fmt.Println(phi)
}
