package main

func main() {
	for i := 1; i < 10; i++ {
		defer panic(i)
	}
}
