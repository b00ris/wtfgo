package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Println("Cpu num: ", runtime.NumCPU())
	for i := 0; i < runtime.NumCPU(); i++ {
		go func(id int) {
			fmt.Println(id)
			for {
			}
		}(i)
	}
	time.Sleep(1 * time.Second)
	// and you could'n continue anymore
}
