package main

import "fmt"

func kirimData(ch chan string, data string) {
	ch <- data
}

func terimaData(ch chan string) {
	data := <-ch
	fmt.Println("Menerima:", data)
}

func main() {
	ch := make(chan string, 5)

	go kirimData(ch, "Hello World!")
	go terimaData(ch)

	var input string
	fmt.Scanln(&input)
}
