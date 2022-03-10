package main

import (
	"fmt"
	"strconv"
)

var i int = 69

// Grouping variables / cleaner version of mass assigning variables
// These variable are global and will be available if this package is used
// some where else
var (
	ACTOR_NAME string = "David Davidson"
	FRIEND     string = "Nosdivad Divda"
	DOCTOR     int    = 10
	SEASON     int    = 14
)

func main() {
	var x int      // Assigning variable way #1
	x = 20         // Assigning variable way #1
	var y int = 40 // Assigning variable way #2
	z := 20        // Assigning variable way #3
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	fmt.Printf("%v, %T, %b, \n", z, z, z)
	fmt.Printf("%v, %T, \n", i, i)
	// Int to string conversion
	str := strconv.Itoa(z)
	fmt.Println(str)
	fmt.Printf("Convert function %v \n", convert(z))
}

func convert(x int) float32 {
	return float32(x)
}
