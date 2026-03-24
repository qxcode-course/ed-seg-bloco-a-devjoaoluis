package main

import "fmt"

func main() {
	var qtd, baruel int
	fmt.Scan(&qtd, &baruel)
	figurinhas := make([]int, baruel)
	var repetidas []int
	var faltantes []int
	//lendo os valores de entrada
	for i := range figurinhas {
		fmt.Scan(&figurinhas[i])
	}

	for i := 1; i < qtd+1; i++ {
		if contarNumero(figurinhas, i) == 0 {
			faltantes = append(faltantes, i)
		} else if contarNumero(figurinhas, i) > 1 {
			for range contarNumero(figurinhas, i) - 1 {
				repetidas = append(repetidas, i)
			}
		}
	}

	if len(repetidas) != 0 {
		printFormatado(repetidas)
	} else {
		fmt.Println("N")
	}
	if len(faltantes) != 0 {
		printFormatado(faltantes)
	} else {
		fmt.Println("N")
	}
}

func contarNumero(array []int, numero int) int {
	count := 0
	for _, num := range array {
		if num == numero {
			count++
		}
	}
	return count
}

func printFormatado(lista []int) {
	for index, num := range lista {
		if index != len(lista)-1 {
			fmt.Printf("%d ", num)
		} else {
			fmt.Printf("%d\n", num)
		}
	}
}
