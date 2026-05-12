package main

import "fmt"

func div(num int) {
	if num == 0 {
		return
	}
	//resultado, resto
	resultado := num / 2
	resto := num % 2
	div(resultado)
	fmt.Printf("%d %d\n", resultado, resto)
}

func main() {
	var entrada int
	fmt.Scan(&entrada)
	div(entrada)
}
