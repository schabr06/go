package main

import (
	"fmt"
	"math"
)

const (
	a = iota // 0
	b        // iota
	c        // iota
	//b = iota // 1
	//c = iota // 2
)

const (
	a2 = iota
)

const (
	_ = iota // In case checking against an unassigned var which would be 0
	catSpecialist
	dogSpecialist
	snakeSpecialist
)

const (
	_  = iota // ignore first value by assigning to blank identifier
	KB = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

const (
	isAdmin = 1 << iota
	isHeadquarters
	canSeeFinances
	canSeeAfrica
	canSeeAsia
	canSeeEurope
	canSeeNorthAmerica
	canSeeSouthAmerica
)

func main() {
	contants()
	example()
	file()
	roles()
	fmt.Println("Iota")
	fmt.Printf("%v, %T\n", a, a)
	fmt.Printf("%v, %T\n", b, b)
	fmt.Printf("%v, %T\n", c, c)
	fmt.Printf("%v, %T\n", a2, a2)
}

func contants() {
	const myConst = 64
	// Compiler will throw error
	// const con float64 = math.Sin(1.57)
	const b int16 = 27
	const pi float64 = math.Pi
	fmt.Printf("%v, %T\n", myConst, myConst)
	fmt.Printf("%v, %T\n", pi, pi)
	fmt.Printf("%v, %T\n", myConst+b, myConst+b)
}

func example() {
	fmt.Println("========================")
	var specialistType int = catSpecialist
	fmt.Printf("%v\n", specialistType == catSpecialist)
	fmt.Println("========================")
}

func file() {
	fmt.Println("========================")
	fileSize := 4000000000.
	fmt.Printf("%.2fGB\n", fileSize/GB)
	fmt.Println("========================")
}

func roles() {
	fmt.Println("========================")
	var roles byte = isAdmin | canSeeFinances | canSeeEurope
	fmt.Printf("%b\n", roles)
	fmt.Printf("Is Admin? %v\n", isAdmin&roles == isAdmin)
	fmt.Printf("Is at Headquarters? %v\n", isHeadquarters&roles == isHeadquarters)
	fmt.Println("========================")
}
