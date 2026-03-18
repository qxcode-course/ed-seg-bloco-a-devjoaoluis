package main

import "fmt"

func main() {
	var qtd int
	count := 0
	fmt.Scanln(&qtd)
	animais := make([]int, qtd)
	for index, _ := range animais {
		fmt.Scan(&animais[index])
	}

	for ind, _ := range animais {
		if ind == 0 {
			if animais[ind] < 0 && animais[ind+1] > 0 {
				animais[ind] = 0
				animais[ind+1] = 0
				count++
			} else if animais[ind] > 0 && animais[ind+1] < 0 {
				animais[ind] = 0
				animais[ind+1] = 0
				count++
			}
		} else if ind == len(animais)-1 {
			if animais[ind] > 0 && animais[ind-1] < 0 {
				animais[ind] = 0
				animais[ind-1] = 0
				count++
			} else if animais[ind] < 0 && animais[ind-1] > 0 {
				animais[ind] = 0
				animais[ind-1] = 0
				count++
			}
		} else {
			if animais[ind] < 0 && animais[ind+1] > 0 {
				animais[ind] = 0
				animais[ind+1] = 0
				count++
			} else if animais[ind] > 0 && animais[ind+1] < 0 {
				animais[ind] = 0
				animais[ind+1] = 0
				count++
			}
		}
	}

	fmt.Println(count)
}
