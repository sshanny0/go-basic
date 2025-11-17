package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("======================")
	fmt.Println("BELAJAR ERROR HANDLING")
	fmt.Println("======================")
	fmt.Println()

	basicError()
	customError()
}

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("Pembagian dengan nol tidak diperbolehkan")
	}
	return a / b, nil
}

func basicError() {
	fmt.Println("1. BASIC ERROR HANDLING")
	fmt.Println("=======================")

	r, err := divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Hasil pembagian: %.2f\n", r)
	}

	_, err = divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println()
}

func customError() {
	fmt.Println("2. CUSTOM ERROR")
	fmt.Println("=======================")

	if err := ValidateAge(25); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Validasi berhasil untuk umur 25 tahun!")
	}

	if err := ValidateAge(-5); err != nil {
		fmt.Println("Error:", err)
	}

	if err := ValidateAge(200); err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println()
}

// custom error
type ValidationError struct {
	Field   string
	Message string
	isValid bool
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("Validasi gagal pada field '%s': %s status: %t", e.Field, e.Message, e.isValid)
}

func ValidateAge(age int) error {
	if age < 0 {
		return &ValidationError{
			Field:   "age",
			Message: "Umur tidak boleh kurang daari 0",
			isValid: false}
	}

	if age > 150 {
		return &ValidationError{
			Field:   "age",
			Message: "Umur tidak boleh lebih daari 150",
			isValid: false}
	}

	return nil
}
