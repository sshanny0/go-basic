package main

import (
	"fmt"
	"time"
)

func cetakAngka(angka int) {
	for i := 0; i < angka; i++ {
		fmt.Println(i)
	}
}
func main() {
	fmt.Println("Sync")
	cetakAngka(3)
	cetakAngka(3)
	fmt.Println()

	fmt.Println("Async")
	go cetakAngka(3)
	go cetakAngka(3)
	fmt.Println()

	go func(message string) {
		fmt.Println(message)
	}("going")

	time.Sleep(time.Second)
	println("DONE!")
}
