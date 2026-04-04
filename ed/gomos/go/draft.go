package main

import "fmt"

type Gomo struct {
	x int
	y int
}

func main() {
	var q int
	var d string
	fmt.Scan(&q)
	fmt.Scan(&d)
	var gomos []Gomo
	for range q {
		var a int
		fmt.Scan(&a)
		var b int
		fmt.Scan(&b)
		gomoAdd := Gomo{
			x: a,
			y: b,
		}
		gomos = append(gomos, gomoAdd)
	}

	for i := len(gomos) - 1; i > 0; i-- {
		gomos[i] = gomos[i-1]
	}

	//Modificando apenas a posição da cabeça da cobra
	switch d {
	case "L":
		gomos[0].x--
	case "R":
		gomos[0].x++
	case "U":
		gomos[0].y--
	case "D":
		gomos[0].y++
	}

	for _, gomo := range gomos {
		fmt.Printf("%d %d\n", gomo.x, gomo.y)
	}
}
