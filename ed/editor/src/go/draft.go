package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
    leitor := bufio.NewScanner(os.Stdin)
    leitor.Scan()
    linha := leitor.Text()
    // pointer := 0
    formatado := []rune(linha)
    pointer := 0
    for i, c := range formatado {
        if c=='<'|| c=='>' {
            if c=='<' {
                pointer--
            } else {
                pointer++
            }
            formatado = append(formatado[:i], formatado[i+1:]...)
            // continue
        // } else if c=='>' {
        //     pointer++
        // }
        }
        pointer++
    }

    //imprimir com enter
    for i, c := range linha {
        if c=='R' {
            fmt.Println("")
        } else if c=='<' || c=='>'{
            continue
        } else {
            if i==pointer {
                fmt.Print("|")
            } 
            fmt.Printf("%s", string(c))
        }
        if i==len(formatado)-1 {
            if pointer>=len(formatado) {
                fmt.Print("|")
            }
            fmt.Println()
        }
    }
    // fmt.Println(pointer)
    // fmt.Println("Hello, World!")
}