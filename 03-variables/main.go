package main

import "fmt"

const (
	AppName    = "Golang Basic"
	AppVersion = "1.0.0"
	MaxUsers   = 1000
)

func main() {
	fmt.Println("=========================")
	fmt.Println("BELAJAR VARIABEL DASAR GO")
	fmt.Println("=========================")
	fmt.Println()

	// deklarasi var
	demoVariabel()

	// Test Konstanta
	demoConstant()

	// Tipe data
	demoTipeData()
}

func demoVariabel() {
	fmt.Println("1. Deklarasi Variabel")
	fmt.Println("=====================")

	// Cara 1: explisit
	var nama string = "Shafira Shalehanny"
	var umur int = 24
	var tinggi float64 = 145.9

	fmt.Println("Menggunakan var:")
	fmt.Printf("- Nama: %s\n", nama)
	fmt.Printf("- Umur: %d\n", umur)
	fmt.Printf("- Tinggi: %.2f cm\n", tinggi)
	fmt.Println()

	// Cara 2: Type reference (Go menebak tipe datanya)
	var negara = "Australia"

	// Cara 3: Short declaration (hanya bisa di dalam fungsi)
	kota := "Ciseeng City"
	isActive := true

	fmt.Println("Menggunakan := (short declaration)")
	fmt.Printf("- Kota: %s\n", negara)
	fmt.Printf("- Negara: %s\n", kota)
	fmt.Printf("- Status: %.t\n", isActive)
	fmt.Println()
}

func demoConstant() {
	fmt.Println("2. Deklarasi Konstanta")
	fmt.Println("======================")

	const Phi = 3.14159
	const dayOfWeek = 7

	fmt.Println("Menggunakan Konstanta:")
	fmt.Printf("- Nama Aplikasi: %s\n", AppName)
	fmt.Printf("- Versi: %s\n", AppVersion)
	fmt.Printf("- Users: %d\n", MaxUsers)
	fmt.Printf("- Phi: %.5f\n", Phi)
	fmt.Printf("- Jumlah hari dalam seminggu: %d\n", dayOfWeek)
	fmt.Println()

	// Menghitung luas lingkaran
	var radius = 7.0 // Untuk 7.0 otomatis menjadi float
	luas := Phi * radius * radius

	fmt.Printf("Luas lingkaran: Radius %.1f cm: %.2f cm2\n", radius, luas)
	fmt.Println()
}

func demoTipeData() {
	fmt.Println("3. Tipe Data")
	fmt.Println("======================")

	// Integer
	var i8 int8 = 127
	var i16 int16 = 32767
	var i32 int32 = 2147483647

	fmt.Println("Integer:")
	fmt.Printf("- int8: -128 hingga %d\n", i8)
	fmt.Printf("- int16: -32768 hingga %d\n", i16)
	fmt.Printf("- int32: -2147483648 hingga %d\n", i32)
	fmt.Println()

	// float
	var gaji float64 = 15000000.50
	var nilai float32 = 85.25

	fmt.Println("Float:")
	fmt.Printf("- Gaji: Rp.%.2f\n", gaji)
	fmt.Printf("- Nilai: %.2f\n", nilai)
	fmt.Println()

	// Boolean
	var lulus bool = true
	fmt.Println("Boolean:")
	fmt.Printf("- Status Kelulusan: Rp.%t\n", lulus)
	fmt.Println()

	// Type safe, penjumlahan ini akan error
	var x int = 10
	var y float32 = 20.5

	// z := x+y ~ invalid operation: x+y (mismatched type int and float32)
	z := float32(x) + y
	fmt.Printf("Peenjumlahan %d + %.2f = %.2f\n", x, y, z)

	// string to byte slice
	teks := "Hai"
	bytes := []byte(teks)
	fmt.Printf("String ke []byte: %s -> %v\n", teks, bytes)
	fmt.Println()
}
