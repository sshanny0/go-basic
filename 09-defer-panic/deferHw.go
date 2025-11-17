package main

import "fmt"

// jalanin di akhir
func main() {
	// defer ~ execution later,  LIFO (LAST IN FIRST OUT)
	defer fmt.Println("World")
	fmt.Println("Hello")
	fmt.Println()

	// recover
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from", r)
		}
	}()
	// panic
	panic("Ada kesalahan!")

	defer fmt.Println("One")
	defer fmt.Println("Two")
	defer fmt.Println("Three")

}
