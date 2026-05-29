package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Set struct {
	data     []int
	size     int
	capacity int
}

func NewSet(capacity int) *Set {
	return &Set{
		data:     make([]int, capacity),
		size:     0,
		capacity: capacity,
	}
}

func (s *Set) binarySearch(value int) int {
	left := 0
	right := s.size - 1

	for left < right {
	mid := (right - left) / 2
		if s.data[mid] == value {
			return mid
		} else if s.data[mid] < value {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}

func (s *Set) Reserve(newCapacity int) {
	newVector := make([]int, newCapacity)
	for i, value := range s.data {
		newVector[i] = value
	}
	s.data = newVector
	s.capacity = newCapacity
}

func (s *Set) Insert(value int) {

	if s.capacity==0 {
		s.Reserve(1)
	}

	if s.size>=s.capacity {
		s.Reserve(s.size*2)
	}

	//fazer lógica de inserir em ordem qqi
	pos := s.size // posição onde vai inserir
    for i := 0; i < s.size; i++ {
        if s.data[i] == value {
            return 
        }
        if s.data[i] > value {
            pos = i 
            break
        }
    }

    for i := s.size; i > pos; i-- {
        s.data[i] = s.data[i-1]
    }

    s.data[pos] = value
    s.size++
}

func (s *Set) Erase(value int) error {
	if index>s.size {
		return errors.New("value not found")
	}

	for i := index; i<s.size; i++ {
		s.data[i] = 0
		s.data[i], s.data[i+1] = s.data[i+1], s.data[i]
	}

	return nil 
}

func (s *Set) Show() string {
	if s.size == 0 {
		return "[]"
	}
	return fmt.Sprintf("[%s]", Join(s.data[:s.size], ", "))
}

func Join(slice []int, sep string) string {
	if len(slice) == 0 {
		return ""
	}
	var result strings.Builder
	fmt.Fprintf(&result, "%d", slice[0])

	for _, value := range slice[1:] {
		fmt.Fprintf(&result, "%s%d", sep, value)
	}
	return result.String()
}



func main() {
	var line, cmd string
	scanner := bufio.NewScanner(os.Stdin)

	v := NewSet(0)
	for scanner.Scan() {
		fmt.Print("$")
		line = scanner.Text()
		fmt.Println(line)
		// fmt.Printf("Valor lido: %s", line)
		parts := strings.Fields(line)
		if len(parts) == 0 {
			continue
		}
		cmd = parts[0]

		switch cmd {
		case "end":
			return
		case "init":
			value, _ := strconv.Atoi(parts[1])
			v = NewSet(value)
		case "insert":
			for _, part := range parts[1:] {
				value, _ := strconv.Atoi(part)
				v.Insert(value)
			}
		case "show":
			fmt.Println(v.Show())
		case "erase":
			value, _ := strconv.Atoi(parts[1])
			v.Erase(value)
		case "contains":
			// value, _ := strconv.Atoi(parts[1])
		case "clear":
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
