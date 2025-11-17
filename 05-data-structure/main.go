package main

import "fmt"

func main() {
	fmt.Println("=========================")
	fmt.Println("BELAJAR STRUKTUR DATA")
	fmt.Println("=========================")
	fmt.Println()

	demonstrasiArray()
	demonstrasiSlice()
	demonstrasiMap()
}

func demonstrasiArray() {
	fmt.Println("1. Array")
	fmt.Println("=====================")

	// array dengan ukuran tetap
	var angka [5]int = [5]int{1, 2, 3, 4, 5}
	buah := [3]string{"apel", "jeruk", "mangga"}

	// ukuran otomatis
	warna := [...]string{"Merah", "Kuning", "Hijau", "Biru"}

	fmt.Printf("Array angka: %v\n", angka)
	fmt.Printf("Array buah: %v\n", buah)
	fmt.Printf("Array warna: %v\n", warna)
	fmt.Printf("Panjang array warna: %d\n", len(warna))
	fmt.Printf("Buah pertama: %v\n", buah[0])
	fmt.Printf("Buah kedua: %v\n", buah[len(buah)-1])
	fmt.Println()
}

func demonstrasiSlice() {
	fmt.Println("2. Slice")
	fmt.Println("=====================")

	// \membuat slice
	var bahasa []string
	fmt.Printf("Kondisi awal: %v (length: %d, capacity: %d)\n", bahasa, len(bahasa), cap(bahasa))

	// append element
	bahasa = append(bahasa, "Go")
	bahasa = append(bahasa, "Javascript", "C++")
	fmt.Printf("Setelah append: %v (length: %d, capacity: %d)\n", bahasa, len(bahasa), cap(bahasa))
	fmt.Println()
}

func demonstrasiMap() {
	fmt.Println("3. Map")
	fmt.Println("=====================")

	// contoh
	m := map[string]int{
		"Hanny":  22,
		"Soobin": 15,
	}
	fmt.Printf("Map: %v\n", m)

	// contoh map interface
	mapInter := map[string]interface{}{
		"nama":     "Shafira Shalehanny",
		"age":      60,
		"isActive": true,
	}
	fmt.Printf("Map Interface: %v\n", mapInter)

	// membuat map
	populasi := make(map[string]int)
	populasi["Beji"] = 1000000
	populasi["Pengasinan"] = 600000
	populasi["Bekasi"] = 92842400
	populasi["Parung"] = 120000

	// iterasi
	fmt.Println("Data Populasi Kota:")
	for kota, jumlah := range populasi {
		fmt.Printf("%s: %d Jiwa\n", kota, jumlah)
	}
	fmt.Println()
}
