package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var s, sep string

	// Version 1
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
		fmt.Println(i, os.Args[i])
	}
	fmt.Println(s)

	// Version 2: anothor form of the for loop iterates over a range of values from a data type like a string or a slice

	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}

	fmt.Println(s)

	// Version 3: This is clearer, faster, and avoids manual state management.
	fmt.Println(strings.Join(os.Args[1:], " "))

	fmt.Println(os.Args[0])
}
