package main

import "fmt"

func main() {
	var h, p, f, d int32
	fmt.Scanln(&h)
	fmt.Scanln(&p)
	fmt.Scanln(&f)
	fmt.Scanln(&d)
	for {
		if f > 15 {
			f = f % 15
		} else if f < 0 {
			f = 15
		}

		if d == -1 {
			f--
		} else {
			f++
		}

		if f == p {
			fmt.Println("N")
			break
		}
		if f == h {
			fmt.Println("S")
			break
		}

	}

}
