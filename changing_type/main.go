package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"sync/atomic"
)

type MyStruct struct {
	SomeColumn string
}

func main() {
	v := atomic.Value{}
	r := strings.NewReader(`{"SomeColumn":"some string"}`)
	m := MyStruct{SomeColumn: "string"}
	v.Store(m)
	newM := v.Load()
	fmt.Println(reflect.TypeOf(newM))
	json.NewDecoder(r).Decode(&newM)
	fmt.Println(reflect.TypeOf(newM))
}
