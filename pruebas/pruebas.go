package main

import "fmt"

func main() {
	// fmt.Printf("%d - %b - %x \n", 42, 42, 42)

	for i := 0; i < 100; i++ {
		fmt.Printf("%d - %b - %x - %q \n", i, i, i, i)

	}
}
