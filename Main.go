package main

import (
	"fmt"
	"log"
)

func main() {
	booleans()
	signedIntegers()
	simpleArithmetic()
	bitOperators()
	floatingPoint()
	complexNumbers()
	string()
	runePrimitive()
	log.Println("Program exited")
}

func booleans() {
	log.Println("Booleans")
	var n bool = true
	var boolean bool // Will evaluate to 0 (false)
	t := 1 == 1
	f := 1 == 2
	fmt.Printf("%v, %T \n", n, n)
	fmt.Printf("%v, %T \n", t, t)
	fmt.Printf("%v, %T \n", f, f)
	fmt.Printf("%v, %T \n", boolean, boolean)
}

func signedIntegers() {
	log.Println("Signed Integers")
	var x uint8 = 255         // Unsigned int 0 to 255
	var y uint16 = 65535      // Unsigned int 0 to 65,535
	var z uint32 = 4000000000 // Unsigned int 0 to 4,294,967,295
	n := 42                   // Unspecified integer (type int)
	var a int8 = 120          // int8 -128 to 127
	var b int16 = 160         // int16 -32,768 to 32,767
	var c int32 = 32800       // int32 -2,147,483,648 to 2,147,483,647
	var d int64 = 2000000     // int64 approx -9,223,000,000,000,000,000 to approx 9,223,000,000,000,000,000
	fmt.Printf("%v, %T \n", n, n)
	fmt.Printf("%v, %T \n", a, a)
	fmt.Printf("%v, %T \n", b, b)
	fmt.Printf("%v, %T \n", c, c)
	fmt.Printf("%v, %T \n", d, d)
	fmt.Printf("%v, %T \n", x, x)
	fmt.Printf("%v, %T \n", y, y)
	fmt.Printf("%v, %T \n", z, z)
}

func simpleArithmetic() {
	log.Println("Simple Arithmetic")
	a := 30
	b := 15
	fmt.Printf("%v \n", a+b)
	fmt.Printf("%v \n", a-b)
	fmt.Printf("%v \n", a*b)
	fmt.Printf("%v \n", a/b)
	fmt.Printf("%v \n", a%b)
}

func bitOperators() {
	log.Println("Bit Operators")
	a := 10             // 1010
	b := 3              // 0011
	fmt.Println(a & b)  // 0010
	fmt.Println(a | b)  // 1011
	fmt.Println(a ^ b)  // 1001
	fmt.Println(a &^ b) // 0100
	c := 8              // 0100 or 2^3
	fmt.Println(c << 3) // 2^3 * 2^3 = 2^6
	fmt.Println(c >> 3) // 2^3 / 2^3 = 2^0
}

func floatingPoint() {
	log.Println("Floating Point")
	var w float32 = 3.14 // float32 +- 1.18E-38 to +- 3.4E38
	var u float64 = 13.7 // float64 +- 2.23E-308 to +- 1.8E308
	z := 2.14            // Will initialize to float 64 default
	fmt.Printf("%v \n", w)
	fmt.Printf("%v \n", u)
	fmt.Printf("%v \n", z)
	fmt.Println(z + u)
	fmt.Println(z - u)
	fmt.Println(z * u)
	fmt.Println(z / u)
}

func complexNumbers() {
	log.Println("Complex Numbers")
	var n complex64 = 1 + 2i
	var j complex64 = 2i
	var i complex128 = 18i + 20i
	var com complex128 = complex(5, 12)
	fmt.Printf("%v, %T\n", real(n), real(n))
	fmt.Printf("%v, %T\n", imag(n), imag(n))
	fmt.Printf("%v, %T\n", i, i)
	fmt.Printf("%v, %T\n", com, com)
	fmt.Println(n + j)
	fmt.Println(n - j)
	fmt.Println(n * j)
	fmt.Println(n / j)
}

func string() {
	log.Println("Strings")
	s := "simple string"
	b := []byte(s)
	fmt.Printf("%v, %T\n", s, s)
	fmt.Printf("%v, %T\n", s[2], s[2])
	fmt.Printf("%v, %T\n", b, b)
}

func runePrimitive() {
	log.Println("Rune")
	var y rune = 900
	x := 'a'
	fmt.Printf("%v, %T\n", y, y)
	fmt.Printf("%v, %T\n", x, x)
}
