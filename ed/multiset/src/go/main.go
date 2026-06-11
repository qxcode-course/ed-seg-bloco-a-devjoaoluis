package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type MultiSet struct {
	data []int
	size int
	capacity int
}

func NewMultiSet(capacity int) *MultiSet {
	return &MultiSet{
		data: make([]int, capacity),
		size: 0,
		capacity: capacity,
	}
}

func (m *MultiSet) expand() {
	// if m.capacity==0 {
	// 	m.capacity++
	// 	return
	// }
	m.capacity = m.size * 2
	new := make([]int, m.capacity)
	for i, value := range m.data {
		new[i] = value
	}
	m.data = new
}

func (m *MultiSet) search(value int) (bool, int) {
	left := 0
	right := m.size

	for left <= right {
		mid := (left+right) / 2
		if m.data[mid]==value {
			return true, mid
		} else if m.data[mid] < value {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return false, -1
}

func (m *MultiSet) Insert(value int) {

	if m.size==m.capacity || m.capacity==0 {
		m.expand()
	}

	if m.size==0 {
		m.data[m.size] = value
		m.size++
		return
	}

	pos := m.size // posição onde vai inserir
    for i := 0; i < m.size; i++ {
        if m.data[i] > value {
            pos = i 
            break
        }
    }

    for i := m.size; i > pos; i-- {
        m.data[i] = m.data[i-1]
    }

    m.data[pos] = value
    m.size++

}

func (m *MultiSet) insert(value int, index int) error {
	return nil
}

func (m *MultiSet) Erase(value int) error {
	ok, index := m.search(value)
	if !ok {
		fmt.Println("value not found")
		return errors.New("value not found")
	}
	m.data[index] = 0
	for i := index; i < m.size; i++ {
		m.data[i], m.data[i+1] = m.data[i+1], m.data[i]
	}
	m.size--
	return nil
}

func (m *MultiSet) Contains(value int) bool {
	ok, _  := m.search(value)
	return ok
}

func (m *MultiSet) Count(value int) int {
	count := 0
	for _, val := range m.data[:m.size] {
		if val==value{
			count++
		}
	}	
	return count
}

func (m *MultiSet) Unique() int {
	if m.size==0 {
		return 0
	}
	count := 1
	for i := range m.size-1 {
		if m.data[i]!=m.data[i+1] {
			count++
		}
	}
	return count
}

func (m *MultiSet) Clear() {
	for i := 0; i < m.size; i++ {
		m.data[i] = 0
	}
	m.size = 0
}

func (m *MultiSet) String() string {
	return "[" + Join(m.data[:m.size], ", ") + "]"
}

func Join(slice []int, sep string) string {
	if len(slice) == 0 {
		return ""
	}
	result := fmt.Sprintf("%d", slice[0])
	for _, value := range slice[1:] {
		result += sep + fmt.Sprintf("%d", value)
	}
	return result
}

func main() {
	var line, cmd string
	scanner := bufio.NewScanner(os.Stdin)
	ms := NewMultiSet(0)

	for scanner.Scan() {
		fmt.Print("$")
		line = scanner.Text()
		args := strings.Fields(line)
		fmt.Println(line)
		if len(args) == 0 {
			continue
		}
		cmd = args[0]

		switch cmd {
		case "end":
			return
		case "init":
			value, _ := strconv.Atoi(args[1])
			ms = NewMultiSet(value)
		case "insert":
			for _, part := range args[1:] {
				value, _ := strconv.Atoi(part)
				ms.Insert(value)
			}
		case "show":
			fmt.Println(ms.String())
		case "erase":
			value, _ := strconv.Atoi(args[1])
			ms.Erase(value)
		case "contains":
			value, _ := strconv.Atoi(args[1])
			fmt.Println(ms.Contains(value))
		case "count":
			value, _ := strconv.Atoi(args[1])
			fmt.Println(ms.Count(value))
		case "unique":
			fmt.Println(ms.Unique())
		case "clear":
			ms.Clear()
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
