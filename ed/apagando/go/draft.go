package main

import "fmt"

func main() {
	var m, n int
	fmt.Scan(&m)
	numeros := make([]int, m)
	// le os numeros
	for i := range m {
		fmt.Scan(&numeros[i])
	}

	fmt.Scan(&n)
	apaga := make(map[int]int)

	for i := 1; i < n+1; i++ {
		fmt.Scan(&m)
		apaga[m] = m
	}

	for index, value := range numeros {

	}

	fmt.Println(numeros)
	fmt.Println(apaga)

}
