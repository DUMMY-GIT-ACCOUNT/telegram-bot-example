package main

/*
#include "factorial.c"
*/
import "C"
import (
	"fmt"
)

func main() {
	fmt.Print("Hello World")
	n := 5
	result := C.factorial(C.int(n))
	fmt.Printf("Factorial of %d is %d\n", n, result)
}
