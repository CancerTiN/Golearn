package main

import "fmt"

type myHeap struct {
	slice  []int
	length int
}

func (h *myHeap) Init(slice []int) {
	h.slice = slice[:]
	h.length = len(slice)
}

func (h *myHeap) Adjust(end int) {
	if !(end < h.length) {
		panic(fmt.Errorf("index out of range [%#v] with length %#v", end, h.length))
	}
	for i := len(h.slice[:end+1])/2 - 1; i >= 0; i-- {
		for j := 2*i + 1; j <= end; j = 2*j + 1 {
			if j+1 <= end && h.slice[j] < h.slice[j+1] {
				j++
			}
			if h.slice[i] < h.slice[j] {
				h.slice[i], h.slice[j] = h.slice[j], h.slice[i]
				i = j
			}
		}
	}
}

func (h *myHeap) Sort() {
	for end := h.length - 1; end > 0; end-- {
		h.Adjust(end)
		h.slice[0], h.slice[end] = h.slice[end], h.slice[0]
	}
}

func (h *myHeap) Push(i int) {
	h.slice = append(h.slice, i)
	h.length = len(h.slice)
	h.Adjust(h.length - 1)
}

func (h *myHeap) Pop() (i int) {
	h.Adjust(h.length - 1)
	i = h.slice[0]
	h.slice = h.slice[1:]
	h.length = len(h.slice)
	h.Adjust(h.length - 1)
	return
}

func main() {
	testHeap()
}

func testHeap() {
	var h myHeap
	h.Init([]int{4, 6, 8, 5, 9})
	fmt.Println(h)
	h.Sort()
	fmt.Println(h)
	h.Push(1000)
	fmt.Println(h)
	h.Sort()
	fmt.Println(h)
	fmt.Println(h.Pop())
	fmt.Println(h)
	h.Push(7)
	fmt.Println(h)
	fmt.Println(h.Pop())
	fmt.Println(h)
}
