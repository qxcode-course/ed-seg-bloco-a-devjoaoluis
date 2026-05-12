package main

import "fmt"

// x: número que está sendo testado
// div: divisor que está sendo testado
func eh_primo(x int, div int) bool {
	_, _ = x, div

	if div == 1 || x == 2 {
		return true
	}

	if x%div == 0 || x == 1 {
		return false
	}

	return eh_primo(x, div-1)
}

func main() {
	var x int
	fmt.Scan(&x)
	if eh_primo(x, 2) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
