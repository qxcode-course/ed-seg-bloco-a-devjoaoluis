package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func getMen(vet []int) []int {
	var men []int
	for _, val := range vet {
		if val > 0 {
			men = append(men, val)
		}
	}
	return men
}

func getCalmWomen(vet []int) []int {
	var calmWomen []int
	for _, val := range vet {
		if val < 0 && val > -10 {
			calmWomen = append(calmWomen, val)
		}
	}

	return calmWomen
}

func sortVet(vet []int) []int {
	for i := range vet {
		for j := 0; j < len(vet)-i-1; j++ {
			if vet[j] > vet[j+1] {
				vet[j], vet[j+1] = vet[j+1], vet[j]
			}
		}
	}
	return vet
}

func sortStress(vet []int) []int {
	for i := range vet {
		for j := 0; j < len(vet)-i-1; j++ {
			if math.Abs(float64(vet[j])) > math.Abs(float64(vet[j+1])) {
				vet[j], vet[j+1] = vet[j+1], vet[j]
			}
		}
	}
	return vet
}

func reverse(vet []int) []int {
	var reversedVet []int
	for i := len(vet) - 1; i > -1; i-- {
		reversedVet = append(reversedVet, vet[i])
	}
	return reversedVet
}

func unique(vet []int) []int {
	freq := make(map[int]int)
	var unicos []int
	for _, value := range vet {
		freq[value]++
		if freq[value] == 1 {
			unicos = append(unicos, value)
		}
	}
	return unicos
}

func repeated(vet []int) []int {
	freq := make(map[int]int)
	var repetidos []int
	for _, value := range vet {
		freq[value]++
		if freq[value] > 1 {
			repetidos = append(repetidos, value)
		}
	}
	return repetidos
}

// função auxiliar
// func count(vet []int, e int) int {
// 	count := 0
// 	for _, value := range vet {
// 		if value == e {
// 			count++
// 		}
// 	}
// 	return count
// }

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		if !scanner.Scan() {
			break
		}
		fmt.Print("$")
		line := scanner.Text()
		args := strings.Split(line, " ")
		fmt.Println(line)

		switch args[0] {
		case "get_men":
			printVec(getMen(str2vet(args[1])))
		case "get_calm_women":
			printVec(getCalmWomen(str2vet(args[1])))
		case "sort":
			printVec(sortVet(str2vet(args[1])))
		case "sort_stress":
			printVec(sortStress(str2vet(args[1])))
		case "reverse":
			array := str2vet(args[1])
			other := reverse(array)
			printVec(array)
			printVec(other)
		case "unique":
			printVec(unique(str2vet(args[1])))
		case "repeated":
			printVec(repeated(str2vet(args[1])))
		case "end":
			return
		}
	}
}

func printVec(vet []int) {
	fmt.Print("[")
	for i, val := range vet {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(val)
	}
	fmt.Println("]")
}

func str2vet(s string) []int {
	if s == "[]" {
		return nil
	}
	s = s[1 : len(s)-1]
	parts := strings.Split(s, ",")
	var vet []int
	for _, part := range parts {
		n, _ := strconv.Atoi(part)
		vet = append(vet, n)
	}
	return vet
}
