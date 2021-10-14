package main

import "fmt"

func main() {
	//version 1
	var s1 student
	s1.name = "Ahmad Emir Alfatah"
	s1.grade = 1
	fmt.Println("name :", s1.name)
	fmt.Println("grade :", s1.grade)

	//version 2
	var s2 = student2{}
	s2.name = "emir"
	s2.age = 21
	s2.grade = 1
	s2.person.name = "jeju"
	s2.person.age = 22
	fmt.Println("student 2 :", s2.name)
	fmt.Println("age   :", s2.age)
	fmt.Println("grade :", s2.grade)
	fmt.Println("name  :", s2.person.name)
	fmt.Println("age   :", s2.person.age)

	//version 3
	var s3 = student3{person{"jeje", 22}, 1, []string{"read"}}
	fmt.Println("student 3 :", s3.person.name)

	//version 4
	var s4 *student = &s1
	fmt.Println("student 1, name :", s1.name)
	fmt.Println("student 4, name :", s4.name)
	s4.name = "ethan"
	fmt.Println("student 1, name :", s1.name)
	fmt.Println("student 4, name :", s4.name)

	//version 5
	var p1 = person{name: "alfath", age: 21}
	var s5 = student2{person: p1, grade: 2}
	fmt.Println("name  :", s5.name)
	fmt.Println("age   :", s5.age)
	fmt.Println("grade :", s5.grade)

	//version 6
	var s6 = struct {
		person
		grade int
	}{}
	s6.person = person{"wikwik", 21}
	s6.grade = 2
	fmt.Println("name  :", s6.person.name)
	fmt.Println("age   :", s6.person.age)
	fmt.Println("grade :", s6.grade)

	//version 7
	var s7 = struct {
		person
		grade int
	}{
		person: person{"wick", 21},
		grade:  2,
	}
	fmt.Println("name  :", s7.person.name)
	fmt.Println("age   :", s7.person.age)
	fmt.Println("grade :", s7.grade)

	//version 8
	var allStudents = []person{
		{name: "emir", age: 23},
		{name: "Ethan", age: 23},
		{name: "bourne", age: 22},
	}
	for _, student8 := range allStudents {
		fmt.Println(student8.name, "age is", student8.age)
	}

	//version 9
	var allStudents2 = []struct {
		person
		grade int
	}{
		{person: person{"wick", 21}, grade: 2},
		{person: person{"wick2", 21}, grade: 3},
		{person: person{"wick3", 21}, grade: 3},
	}
	for _, student9 := range allStudents2 {
		fmt.Println(student9)
	}

	//version 10
	var student10 struct {
		person
		grade int
	}
	student10.person = person{"wick", 21}
	student10.grade = 2

	//version 11
	var p2 = struct {
		name string
		age  int
	}{age: 22, name: "emir"}
	var p3 = struct {
		name string
		age  int
	}{"hehe", 23}
	fmt.Println(p2)
	fmt.Println(p3)

	//version 12
	var p4 = person{"hehe", 21}
	fmt.Println(p4)
	var p5 = People{"wick", 21}
	fmt.Println(p5)

	//version 13
	var p6 = person2{}
	p6.name = "hai"
	p6.age = 1
	p6.hobbies = []string{"mandi"}
	fmt.Println(p6)
}

type student struct {
	name  string
	grade int
}

type student2 struct {
	grade int
	person
}

type student3 struct {
	person struct {
		name string //`tag1`
		age  int    //`tag2`
	}
	grade   int
	hobbies []string
}

type person struct {
	name string
	age  int
}

type person2 struct {
	name    string
	age     int
	hobbies []string
}

type People = person
