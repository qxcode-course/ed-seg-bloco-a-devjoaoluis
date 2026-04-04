package main

import "fmt"

func main() {
	var n int
	var e int
	var elementos []int
	fmt.Scan(&n, &e)

	//adição dos elementos
	for i := 1; i < n+1; i++ {
		elementos = append(elementos, i)
	}

	for {
		if len(elementos) == 1 {
			break
		} else {
			printFormatado(elementos, e)
			matarProximo(elementos, e)
			e = proxKiller(elementos, e)
			break
		}
	}
}

func proxKiller(arr []int, antigo int) int {
	var killer int
	for ind, value := range arr {
		if value == antigo {
			killer = ind + 1
			if killer != 0 {
				return killer % (len(arr))
			}
		}
	}
	return killer
}

// função feita para printar no formato que a questão pede
func printFormatado(arr []int, esc int) {
	fmt.Print("[ ")
	for _, num := range arr {
		if num == esc {
			fmt.Printf("%d> ", num)
		} else {
			fmt.Printf("%d ", num)
		}
	}
	fmt.Printf("]\n")
}

func matarProximo(arr []int, esc int) []int {
	var killerInd int
	matar := 0
	for index, value := range arr {
		if value == 0 {
			continue
		}

		if value == esc {
			killerInd = index
			matar = killerInd + 1
		}

		if matar >= len(arr)-1 {
			matar = matar % (len(arr) - 1)
		}
		arr[matar] = 0
	}
	return arr
	// index := 0
	// if len(arr) == 1 {
	// 	return arr
	// } else {
	// 	for ind, value := range arr {
	// 		if value == 0 {
	// 			continue
	// 		}
	// 		tam := len(arr)
	// 		if value == esc {
	// 			index = ind + 1
	// 			if index > len(arr)-1 {
	// 				index = index%tam - 1
	// 			}
	// 			arr[index] = 0
	// 			break
	// 		}
	// 	}
	// }
	// return arr
}
