package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func tostr(vet []int) string {
	var builder strings.Builder

	if len(vet) == 0 {
		return "[]"
	}

	if len(vet) == 1 {
		return fmt.Sprintf("[%d]", vet[0])
	}

	builder.WriteString("[")
	for i, val := range vet {
		if i != len(vet)-1 {
			builder.WriteString(fmt.Sprintf("%s, ", strconv.Itoa(val)))
		} else {
			builder.WriteString(strconv.Itoa(val))
		}
	}
	builder.WriteString("]")

	resultado := builder.String()
	return resultado
}

func tostrrev(vet []int) string {
	reverse(vet)
	return tostr(vet)
}

// reverse: inverte os elementos do slice
func reverse(vet []int) {
	for i, j := 0, len(vet)-1; i < j; i, j = i+1, j-1 {
		vet[i], vet[j] = vet[j], vet[i]
	}
}

// sum: soma dos elementos do slice
func sum(vet []int) int {
	if len(vet) == 0 {
		return 0
	}
	return vet[0] + sum(vet[1:])
}

// mult: produto dos elementos do slice
func mult(vet []int) int {
	if len(vet) == 0 {
		return 1
	}
	return vet[0] * mult(vet[1:])
}

// min: retorna o índice e valor do menor valor
// crie uma função recursiva interna do modelo
// var rec func(v []int) (int, int)
// para fazer uma recursão que retorna valor e índice
func min(vet []int) int {
	if len(vet) == 0 {
		return -1
	}

	var rec func(v []int, idxAtual int) (int, int)

	rec = func(v []int, idxAtual int) (int, int) {
		if len(v) == 1 {
			return v[0], idxAtual
		}

		minValorRestante, minIdxRestante := rec(v[1:], idxAtual+1)

		if v[0] < minValorRestante {
			return v[0], idxAtual
		}
		return minValorRestante, minIdxRestante
	}

	_, idx := rec(vet, 0)
	return idx
}

func main() {
	var vet []int
	scanner := bufio.NewScanner(os.Stdin)
	for {
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		args := strings.Fields(line)
		fmt.Println("$" + line)

		switch args[0] {
		case "end":
			return
		case "read":
			vet = nil
			for _, arg := range args[1:] {
				if val, err := strconv.Atoi(arg); err == nil {
					vet = append(vet, val)
				}
			}
		case "tostr":
			fmt.Println(tostr(vet))
		case "torev":
			fmt.Println(tostrrev(vet))
		case "reverse":
			reverse(vet)
		case "sum":
			fmt.Println(sum(vet))
		case "mult":
			fmt.Println(mult(vet))
		case "min":
			fmt.Println(min(vet))
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
