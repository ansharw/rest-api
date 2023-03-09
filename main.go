package main

import "fmt"

func main() {
	fmt.Println("test ini punya master")

	GetNumber(2)
}

func GetNumber(number int) string {
	return fmt.Sprintf("ini angka: %v", number)
}