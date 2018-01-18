package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
	s = ""
	// Add for range
	for _, arg := range os.Args[2:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}
