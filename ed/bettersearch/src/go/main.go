package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func BetterSearch(slice []int, value int) (bool, int) {
	// _, _ = slice, value
	// return false, 0
	left := 0
	right := len(slice)

	if right==0 {
		return false, 0
	}

	for left < right {
		mid := (right + left) / 2
		if slice[mid] == value {
			return true, mid
		} else if slice[mid] < value {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return false, left
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	parts := strings.Split(scanner.Text(), " ")
	slice := []int{}
	for _, elem := range parts[1 : len(parts)-1] {
		value, _ := strconv.Atoi(elem)
		slice = append(slice, value)
	}
	scanner.Scan()
	value, _ := strconv.Atoi(scanner.Text())
	found, result := BetterSearch(slice, value)
	if found {
		fmt.Println("V", result)
	} else {
		fmt.Println("F", result)
	}
}
