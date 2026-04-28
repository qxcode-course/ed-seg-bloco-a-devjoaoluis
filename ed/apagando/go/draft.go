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

	//Mapa dos que serão apagados
	fmt.Scan(&n)
	apagar := make(map[int]int, n)
	for range n {
		var num int
		fmt.Scan(&num)
		apagar[num] = num
	}

	var apagados []int
	for _, val := range numeros {
		if val == apagar[val] {
			continue
		} else {
			apagados = append(apagados, val)
		}
	}
	printFormatado(apagados)
}

func printFormatado(arr []int) {
	for i, val := range arr {
		if i != len(arr)-1 {
			fmt.Printf("%d ", val)
		} else {
			fmt.Printf("%d \n", val)
		}
	}
}
