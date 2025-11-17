package main

import (
	"fmt"
	"net/http"
	"os"
)

// Panic with Defer Execution
func panicWithDefer() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()
	panic("A panic occurred!")
}

// File Processing with Defer Cleanup
func fileProcessing() {
	file, err := os.Open("sample.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()
	fmt.Println("File opened successfully")
}

// Transaction Rollback Pattern
func transactionRollback() {
	defer func() {
		fmt.Println("Transaction rolled back")
	}()
	fmt.Println("Starting transaction")
	panic("Transaction error")
}

// HTTP Handler Panic Recovery Pattern
func handler(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if r := recover(); r != nil {
			http.Error(w, "Recovered from panic", http.StatusInternalServerError)
		}
	}()
	panic("Simulated panic in handler")
}

func main() {
	panicWithDefer()

	fileProcessing()

	defer transactionRollback()

	http.HandleFunc("/", handler)
	fmt.Println("Starting server at :8080")
	http.ListenAndServe(":8080", nil)
}
