// The actual issue is https://github.com/golang/go/issues/4876.
// You also can read "Bugs" section in https://golang.org/pkg/reflect/#pkg-note-BUG

package main

import (
	"fmt"
	"reflect"

	"./wheel"
)

type Bodywork struct {
	Seats int
	wight int
	metal int
}

type Car struct {
	Bodywork
	wheel.Wheel
}

func main() {
	c := Car{}
	c.wight = 1000

	fmt.Println(reflect.TypeOf(c.wight))

	_, ok := reflect.TypeOf(c).FieldByName("metal")
	fmt.Println(ok)

	_, ok = reflect.TypeOf(c).FieldByName("wight")
	fmt.Println(ok)
}
