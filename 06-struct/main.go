package main

import "fmt"

// structs mahasiswa
type Mahasiswa struct {
	Nama    string
	NIM     string
	Jurusan string
	IPK     float64
}

func main() {
	demonstrasiStruct()
}

func (m Mahasiswa) PrintInfoMahasiswa() {
	fmt.Println("Data Mahasiswa")
	fmt.Println("Nama Mahasiswa:", m.Nama)
	fmt.Println("NIM Mahasiswa:", m.NIM)
	fmt.Println("Jurusan Mahasiswa:", m.Jurusan)
	fmt.Printf("IPK Mahasiswa: %.2f\n", m.IPK)
}

// receiver
func (m Mahasiswa) GetFullInfo() string {
	return fmt.Sprintf("%s (NIM: %s) - %s", m.Nama, m.NIM, m.Jurusan)
}

// *(Pointer)
func (m *Mahasiswa) updateIPK(ipkBaru float64) {
	m.IPK = ipkBaru
}

// cek kelulusan
func (m Mahasiswa) IsLulus() (float64, string) {
	if m.IPK >= 2.75 {
		return 2.75, "LULUS"
	} else {
		return 2.75, "TIDAK LULUS"
	}
}

// constructor
func NewMahasiswa(nama, nim, jurusan string, ipk float64) *Mahasiswa {
	return &Mahasiswa{
		Nama:    nama,
		NIM:     nim,
		Jurusan: jurusan,
		IPK:     ipk,
	}
}

func demonstrasiStruct() {
	fmt.Println("Struct")
	fmt.Println("=====================")

	mahasiswa := Mahasiswa{
		Nama:    "Shafira Shalehanny",
		NIM:     "183112700650145",
		Jurusan: "Sawangan-Bojong Sari",
		IPK:     3.63,
	}

	// tampilkan
	mahasiswa.PrintInfoMahasiswa()
	fmt.Println("Info Lengkap: ", mahasiswa.GetFullInfo())
	println()

	// Update IPK: Pointer
	fmt.Println("Update IPK")
	fmt.Println("=====================")
	mahasiswa.updateIPK(3.56)
	fmt.Printf("Update IPK: %.2f\n", mahasiswa.IPK)

	// & (ampersand) -> mengambil alamat memori
	x := 7
	y := &x
	println("X: ", x)
	println("Y: ", y)

	// Demo alasan dibutuhkannya & dan *
	num := 25
	ubahNilaiX(&num)
	fmt.Println("Num:", num)

	// method cek kelulusan, min IPK 2.75
	minIPK, status := mahasiswa.IsLulus()
	fmt.Printf("Min IPK: (%.2f) | Hasil: %s\n", minIPK, status)

	fmt.Println("Constructor")
	fmt.Println("=====================")
	mhs1 := NewMahasiswa(
		"Choi Soobin",
		"12345",
		"Teknik Industri",
		3.65,
	)
	mhs2 := NewMahasiswa(
		"Choi Yeonjun",
		"678910",
		"Teknik Sipil",
		3.89,
	)

	fmt.Println("Mahasiswa 1: ", mhs1.GetFullInfo())
	fmt.Println("Mahasiswa 2: ", mhs2.GetFullInfo())
}

// Demo alasan dibutuhkannya & dan *
func ubahNilaiX(x *int) {
	*x = 100
}
