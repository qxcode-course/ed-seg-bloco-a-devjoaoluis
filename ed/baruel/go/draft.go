package main

import "fmt"

func main() {
	fmt.Println("Qx Code")
}

func buscarNumero(array []int, numero int) int {
	count := 0
	for _, num := range array {
		if num == numero {
			count++
		}
	}
	return count
}
