package main

import (
	"fmt"
)

func main() {
	fmt.Println("Main")
	// grades()
	// array()
	// slice()
	// other()
	operator()
}

func grades() {
	grades := [...]int{97, 85, 83}
	var students [3]string
	fmt.Printf("Grades: %v \n", grades)
	fmt.Printf("Students: %v \n", students)
	students[0] = "David"
	students[1] = "Divad"
	students[2] = "Animal"
	fmt.Printf("Students: %v \n", students)
	fmt.Printf("Students: %v \n", len(students))
}

func array() {
	fmt.Println("================")
	a := [...]int{1, 2, 3}
	b := a  // A copy of array a is made (not a pointer referencing original)
	c := &a // Array c pointer referencing original Array a
	b[1] = 5
	a[0] = 30
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println("================")
}

func slice() {
	fmt.Println("================")
	a := []int{1, 2, 3}
	fmt.Println(a)
	aa := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	b := aa[:]   // Slice of all element
	c := aa[3:]  // Slice from 4th element to end
	d := aa[:6]  // Slice from first to 6th element
	e := aa[3:6] // Slice from 4th to 6th element
	fmt.Println(aa)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	make := make([]int, 3, 100)
	fmt.Println(make)
	fmt.Println("================")
}

func other() {
	fmt.Println("================")
	a := []int{}
	fmt.Println(a)
	fmt.Printf("Length: %v\n", len(a))
	fmt.Printf("Capacity: %v\n", cap(a))
	fmt.Println(a)
	a = append(a, 1)
	a = append(a, []int{2, 3, 4, 5}...)
	fmt.Println(a)
	fmt.Printf("Length: %v\n", len(a))
	fmt.Printf("Capacity: %v\n", cap(a))
	fmt.Println("================")
}

func operator() {
	a := []int{1, 2, 3, 4, 5, 6, 7}
	// b := a[1:] // Remove from start
	// b := a[:len(a)-1] // Remove from end
	b := append(a[:2], a[3:]...)
	fmt.Println(a)
	fmt.Println(b)
}
