package main

import (
	"02-workspace/utils"
	"fmt"
)

func main() {
	fmt.Println("Demo Package")
	greeting := utils.SayHello("Hanny")
	fmt.Println(greeting)
	fmt.Println()

	fmt.Println("Demo Greet")
	greetingDynamic := utils.SayGoodMorning("Hanny")
	fmt.Println(greetingDynamic)
	fmt.Println()

	// deklarasikan variabel
	fmt.Println("Parameter Sprintf")
	var sfstring = fmt.Sprintf("Name: %s", "Hanny")
	fmt.Println(sfstring)
	// deklarasikan angka
	var sfnumber = fmt.Sprintf("Age: %d", 22)
	fmt.Println(sfnumber)
	// deklarasikan decimal
	var phi = fmt.Sprintf("Phi: %f", 3.14)
	fmt.Println(phi)
}
