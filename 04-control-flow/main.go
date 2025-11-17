package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// if-else if-else
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Need one or more arguments")
		return
	}

	score, err := strconv.ParseFloat(arguments[1], 64)

	// apabila ada error
	if err != nil {
		fmt.Println(err)
	}

	if score >= 95 {
		fmt.Println("Cumlaude")
	} else if score >= 75 {
		fmt.Println("Selamat anda lulus!")
	} else {
		fmt.Println("Maaf anda belum lulus!")
	}

	// switch
	day := 4
	switch day {
	case 1:
		fmt.Println("Senin")

	case 2:
		fmt.Println("Selasa")

	case 3:
		fmt.Println("Rabu")

	case 4:
		fmt.Println("Kamis")

	case 5:
		fmt.Println("Jumat")

	case 6:
		fmt.Println("Sabtu")

	case 7:
		fmt.Println("Minggu")
	default:
		fmt.Println("Hari tidak diketahui")
	}

	// for loop
	fmt.Println("For Loop")
	for i := 0; i < 10; i++ {
		if i == 5 {
			continue
		}
		if i == 8 {
			break
		}

		fmt.Println(i)
	}
}
