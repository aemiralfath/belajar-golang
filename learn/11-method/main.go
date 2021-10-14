package main

import "fmt"
import "strings"

func main() {
	//version 1
	var s1 = student{"ahmad emir", 21}
	s1.sayHello()

	// version 2
	var name = s1.getNameAt(2)
	fmt.Println("nama panggilan :", name)

	// version 3
	fmt.Println("s1 before", s1.name)
	s1.changeName1("jeju juju")
	fmt.Println("s1 after changeName1", s1.name) //ahmad emir

	// version 3
	s1.changeName2("ethan hunt")
	fmt.Println("s1 after changeName2", s1.name) //ethan hunt
}

func (s student) sayHello() {
	fmt.Println("halo", s.name)
}

func (s student) getNameAt(i int) string {
	return strings.Split(s.name, " ")[i-1]
}

func (s student) changeName1(name string) {
	fmt.Println("---> on changeName1, name changed to", name)
	s.name = name
}

func (s *student) changeName2(name string) {
	fmt.Println("---> on changeName2, name changed to", name)
	s.name = name
}

type student struct {
	name  string
	grade int
}
