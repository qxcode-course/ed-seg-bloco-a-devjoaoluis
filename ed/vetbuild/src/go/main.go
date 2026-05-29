package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Vector struct {
	data     []int
	size     int
	capacity int
}

func NewVector(capacity int) *Vector {
	return &Vector{
		data:     make([]int, capacity), // nunca use len(data) ou cap(data) ou qq método do go de manipulação de array
		size:     0,
		capacity: capacity,
	}
}

func (v *Vector) Reserve(newCapacity int) {
	// novo := make([]int, v.size, newCapacity)
	novo := NewVector(newCapacity)
	copy(novo.data, v.data)
	v.data = novo.data
	v.capacity = newCapacity
}

func (v *Vector) PushBack(value int) {
	if v.Size() == v.Capacity() {
		v.Reserve(v.Capacity() * 2)
	}
	v.data[v.Size()] = value
	v.size++
}

func (v *Vector) Size() int {
	return v.size
}

func (v *Vector) PopBack() (int, error) {
	if v.Size() == 0 {
		return -1, errors.New("vector is empty")
	}
	v.Set(v.Size(), 0)
	v.size--
	return -1, nil
}

func (v *Vector) Capacity() int {
	return v.capacity
}

func (v *Vector) Status() string {
	// cap := 0
	// for range v.data {
	// 	cap++
	// }
	// capString := strconv.Itoa(cap)
	return "size:" + strconv.Itoa(v.size) + " capacity:" + strconv.Itoa(v.capacity)
}

func (v *Vector) String() string {
	return v.Show()
}

func (v *Vector) Get(index int) int {
	return v.data[index]
}

func (v *Vector) At(index int) (int, error) {
	if index > v.Size()-1 {
		return -1, errors.New("index out of range")
	}
	return v.data[index], nil
}

func (v *Vector) Set(index int, value int) error {
	_, err := v.At(index)
	if err == nil {
		v.data[index] = value
	}
	//verificar como que se usa o error
	return err
}

func (v *Vector) Clear() {
	for index := range v.data[:v.Size()] {
		v.Set(index, 0)
	}
	v.size = 0
}

func (v *Vector) Insert(index int, value int) error {
	//o erro pode vir nil ou um error msm
	var aux []int
	temp := v.Get(index) //pega o valor antigo do index
	aux = append(aux, temp)
	aux = append(aux, v.data[index:]...) // REESCREVER SEM APPEND
	// if index != 0 {
	// 	aux = v.data[index-1:]
	// } else {
	// 	temp := v.Get(index)

	// }
	// v.Reserve(v.Capacity() + 1)
	if v.Size() == v.Capacity() {
		v.Reserve(v.Capacity() * 2)
	}
	err := v.Set(index, value)
	v.size++
	for i := index + 1; i < v.Size(); i++ {
		v.Set(i, aux[i-index])
	}

	return err
}

func (v *Vector) Erase(index int) error {
	_, err := v.At(index)

	if err != nil {
		return err
	}

	// if index >= v.Size() {
	// 	return err
	// }
	// v.size--

	//case correto para esse index
	// if index == 0 {
	// 	v.data = v.data[0:]
	// 	return nil
	// }

	//rebola o elemento que está no index lá pro final
	for i := index; i < v.Size()-1; i++ {
		v.data[i] = v.data[i+1]
	}
	v.data[v.Size()-1] = 0
	v.size--
	return nil
}

func (v *Vector) IndexOf(value int) int {
	for i, val := range v.data {
		if val == value {
			return i
		}
	}
	return -1
}

func (v *Vector) Contains(value int) bool {
	return v.IndexOf(value) != -1
}

func (v *Vector) Slice(start int, end int) *Vector {
	if start == -1 {
		result := NewVector(v.Size())
		for i := 0; i < v.Size(); i++ {
			result.data[i] = v.data[v.Size()-1-i]
		}
		result.size = v.Size()
		return &Vector{
			data:     result.data,
			size:     result.size,
			capacity: result.capacity,
		}
	}
	if end == -1 {
		end = v.Size() - 1
	}

	if start < 0 || end > v.Size() || start > end {
		return nil
	}

	tamanho := end - start

	result := NewVector(tamanho)
	for i := 0; i < tamanho; i++ {
		result.data[i] = v.data[start+i]
	}
	result.size = tamanho
	return &Vector{
		data:     result.data,
		size:     result.size,
		capacity: result.capacity,
	}
}

func (v *Vector) Show() string {
	// if v.Size() == 0 {
	// 	fmt.Println("[]")
	// } else {
	// 	fmt.Print("[")
	// 	for index, value := range v.data {
	// 		if value != 0 {
	// 			if index != v.Size() {
	// 				fmt.Printf("%d, ", value)
	// 			} else {
	// 				fmt.Print(value)
	// 			}
	// 		} else {
	// 			continue
	// 		}
	// 	}
	// 	fmt.Print("]\n")
	// }

	if v.Size() == 0 {
		return "[]"
	}
	return "[" + Join(v.data[:v.Size()], ", ") + "]"

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

	v := NewVector(0)
	for {
		fmt.Print("$")
		if !scanner.Scan() {
			break
		}
		line = scanner.Text()
		fmt.Println(line)
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
			v = NewVector(value)
		case "push":
			for _, part := range parts[1:] {
				value, _ := strconv.Atoi(part)
				v.PushBack(value)
			}
		case "show":
			// v.Show()
			fmt.Println(v.Show())
		case "status":
			fmt.Println(v.Status())
		case "pop":
			_, err := v.PopBack()
			if err != nil {
				fmt.Println(err)
			}
		case "insert":
			index, _ := strconv.Atoi(parts[1])
			value, _ := strconv.Atoi(parts[2])
			err := v.Insert(index, value)
			if err != nil {
				fmt.Println(err)
			}
		case "erase":
			index, _ := strconv.Atoi(parts[1])
			err := v.Erase(index)
			if err != nil {
				fmt.Println(err)
			}
		case "indexOf":
			value, _ := strconv.Atoi(parts[1])
			index := v.IndexOf(value)
			fmt.Println(index)
		case "contains":
			value, _ := strconv.Atoi(parts[1])
			if v.Contains(value) {
				fmt.Println("true")
			} else {
				fmt.Println("false")
			}
		case "clear":
			v.Clear()
		case "capacity":
			fmt.Println(v.Capacity())
		case "get":
			index, _ := strconv.Atoi(parts[1])
			value, err := v.At(index)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(value)
			}
		case "set":
			index, _ := strconv.Atoi(parts[1])
			value, _ := strconv.Atoi(parts[2])
			err := v.Set(index, value)
			if err != nil {
				fmt.Println(err)
			}

		case "reserve":
			newCapacity, _ := strconv.Atoi(parts[1])
			v.Reserve(newCapacity)
		case "slice":
			start, _ := strconv.Atoi(parts[1])
			end, _ := strconv.Atoi(parts[2])
			slice := v.Slice(start, end)
			fmt.Println(slice)
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
