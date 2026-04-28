package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Pair struct {
	One int
	Two int
}

func sortPair(vet []Pair) []Pair {
	for i := range vet {
		for j := 0; j < len(vet)-1-i; j++ {
			if vet[j].One > vet[j+1].One {
				vet[j], vet[j+1] = vet[j+1], vet[j]
			}
		}
	}
	return vet
}

func sort(vet []int) []int {
	for i := range vet {
		for j := 0; j < len(vet)-1-i; j++ {
			if vet[j] > vet[j+1] {
				vet[j], vet[j+1] = vet[j+1], vet[j]
			}
		}
	}
	return vet
}

func occurr(vet []int) []Pair {

	frequencia := make(map[int]int)

	//O mapa é usado para usar o valor único da chave e valor como frequência
	for _, value := range vet {
		valor := math.Abs(float64(value))
		valorInt := int(valor)
		frequencia[valorInt]++
	}

	ocorrencias := make([]Pair, 0)

	for chave, value := range frequencia {
		par := Pair{
			One: chave,
			Two: value,
		}
		ocorrencias = append(ocorrencias, par)
	}

	return sortPair(ocorrencias)
}

func teams(vet []int) []Pair {
	pares := make([]Pair, 0)

	if len(vet) == 1 {
		adicionar := Pair{
			One: vet[0],
			Two: 1,
		}
		pares = append(pares, adicionar)
	}

	for i := 0; i < len(vet)-1; i++ {

		// Conta até o indíce do último igual
		count := 0
		if vet[i] == vet[i+1] {
			count++
		}

	}

	return pares
}

func mnext(vet []int) []int {
	_ = vet
	return nil
}

func alone(vet []int) []int {
	_ = vet
	return nil
}

func couple(vet []int) int {
	_ = vet
	return 0
}

func hasSubseq(vet []int, seq []int, pos int) bool {
	_ = vet
	_ = seq
	_ = pos
	return false
}

func subseq(vet []int, seq []int) int {
	_ = vet
	_ = seq
	return -1
}

func erase(vet []int, posList []int) []int {
	newVet := make([]int, 0)
	posList = sort(posList)
	//Iterando sobre o maior vetor
	for i := range vet {

		if i+1 >= len(posList) {
			break
		}

		if i == posList[i] {
			continue
		} else {
			newVet = append(newVet, vet[i])
		}

	}

	//Iterando sobre o menor vetor
	// for i, value := range posList {
	// 	if
	// }
	return newVet
}

func clear(vet []int, value int) []int {
	newVet := make([]int, 0)
	for _, val := range vet {
		if val != value {
			newVet = append(newVet, val)
		}
	}
	return newVet
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("$")
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		args := strings.Split(line, " ")
		fmt.Println(line)

		switch args[0] {
		case "occurr":
			printSlice(occurr(str2vet(args[1])))
		case "teams":
			printSlice(teams(str2vet(args[1])))
		case "mnext":
			printSlice(mnext(str2vet(args[1])))
		case "alone":
			printSlice(alone(str2vet(args[1])))
		case "erase":
			printSlice(erase(str2vet(args[1]), str2vet(args[2])))
		case "clear":
			val, _ := strconv.Atoi(args[2])
			printSlice(clear(str2vet(args[1]), val))
		case "subseq":
			fmt.Println(subseq(str2vet(args[1]), str2vet(args[2])))
		case "couple":
			fmt.Println(couple(str2vet(args[1])))
		case "end":
			return
		default:
			fmt.Println("Invalid command")
		}
	}
}

// Funções auxiliares

func str2vet(str string) []int {
	if str == "[]" {
		return nil
	}
	str = str[1 : len(str)-1]
	parts := strings.Split(str, ",")
	var vet []int
	for _, part := range parts {
		num, _ := strconv.Atoi(strings.TrimSpace(part))
		vet = append(vet, num)
	}
	return vet
}

func printSlice[T any](vet []T) {
	fmt.Print("[")
	for i, x := range vet {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(x)
	}
	fmt.Println("]")
}

func (p Pair) String() string {
	return fmt.Sprintf("(%v, %v)", p.One, p.Two)
}
