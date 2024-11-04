package main

import "fmt"

func main() {
	var input int
	fmt.Printf("Masukkan angka: ")
	_, err := fmt.Scan(&input)

	if err != nil {
		fmt.Println("Input tidak valid")
		return
	} else if input == 42 {
		fmt.Println("Hello Universe")
	} else {
		fmt.Printf("Anda memasukkan angka: %d\n", input)
	}
}
