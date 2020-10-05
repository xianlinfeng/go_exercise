package main

import (
	"fmt"
	"strings"
)

func main() {
	var str1 = "This is a string variable"
	fmt.Println(str1)

	// multiple line string
	fmt.Printf("This is a multiple \n line string\n")

	// string index
	s := "go string example"
	for k, v := range s {
		fmt.Printf("k %d, v :%c == %d \n", k, v, v)
	}

	// string join
	a := "hel" + "lo, "
	a += "world!"
	fmt.Println(a)

	b := strings.Join([]string{"hello", ", "}, "world")
	fmt.Println(b)
}
