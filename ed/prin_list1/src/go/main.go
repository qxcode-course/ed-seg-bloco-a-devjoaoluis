package main

import (
	"fmt"
	"strings"
)

// mostra a lista com o elemento sword destacado
func ToStr(l *DList[int], sword *DNode[int]) string {
	var result strings.Builder
	result.WriteString("[ ")
	for node := l.Front(); node != l.End(); node = node.Next() {
		result.WriteString(fmt.Sprintf("%d", node.Value))

		if node == sword {
			result.WriteString(">")
		}
		
		if node.next != l.End() {
            result.WriteString(" ")
        }
	}
	result.WriteString(" ]")
	return result.String()
}

// move para frente na lista circular
func Next(l *DList[int], it *DNode[int]) *DNode[int] {
	if it.next != l.End() {
    	return it.next
	}	
	return l.Front()
}

func main() {
	var qtd, chosen int
	fmt.Scan(&qtd, &chosen)
	// fmt.Println(qtd, chosen)
	l := NewDList[int]()
	for i := 1; i <= qtd; i++ {
		l.PushBack(i)
	}
	sword := l.Front()
	for range chosen - 1 {
		sword = Next(l, sword)
	}
	for range qtd - 1 {
		fmt.Println(ToStr(l, sword))
		l.Erase(Next(l, sword))
		sword = Next(l, sword)
	}
	fmt.Println(ToStr(l, sword))
}
