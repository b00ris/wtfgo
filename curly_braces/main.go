package main

import (
	"fmt"
)

type array [10]byte

type str struct {
	i int
	j float64
}

func main() {
	var a1 array

	fmt.Println(a1 == array{})
	if a1 == (array{}) {
		fmt.Println("ok")
	}
	/*
		//won't compile
		if a1 == array{} {
			fmt.Println("not ok, for sure)
		}
	*/

	var s1 str

	fmt.Println(s1 == str{})
	if s1 == (str{}) {
		fmt.Println("ok")
	}
	/*
		// won't compile
		if s1 == str{} {
			fmt.Println("not ok, for sure)
		}
	*/
}
