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

	//Como o
	p := e - 1

	//Abordagem 2
	for len(elementos) > 1 {
		printFormatado(elementos, p)
		matar := (p + 1) % len(elementos)
		elementos = append(elementos[:matar], elementos[matar+1:]...)
		p = matar % len(elementos)
	}

	printFormatado(elementos, p)

}

// func proxKiller(arr []int, antigo int) int {
// 	var killer int
// 	for ind, value := range arr {
// 		if value == antigo {
// 			killer = ind + 1
// 			if killer != 0 {
// 				return killer % (len(arr))
// 			}
// 		}
// 	}
// 	return killer
// }

// função feita para printar no formato que a questão pede
func printFormatado(arr []int, esc int) {
	fmt.Print("[ ")
	for i, num := range arr {
		if i == esc {
			fmt.Printf("%d> ", num)
		} else {
			fmt.Printf("%d ", num)
		}
	}
	fmt.Printf("]\n")
}

// func matarProximo(arr []int, esc int) []int {
// 	var killerInd int
// 	matar := 0
// 	for index, value := range arr {
// 		if value == 0 {
// 			continue
// 		}

// 		if value == esc {
// 			killerInd = index
// 			matar = killerInd + 1
// 		}

// 		if matar >= len(arr)-1 {
// 			matar = matar % (len(arr) - 1)
// 		}
// 		arr[matar] = 0
// 	}
// 	return arr

// }
