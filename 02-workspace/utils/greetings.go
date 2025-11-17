package utils

import "fmt"

func SayHello(name string) string {
	return fmt.Sprintf("Selamat Pagi, %s!", name)
}

func SayGoodMorning(name string) string {
	return greet("pagi", name)
}
func greet(timeofDay, name string) string {
	return fmt.Sprintf("Selamat %s, %s", timeofDay, name)
} //tidak bisa langsung akses dari luar package karena penamaannya menggunakan huruf kecil
