package main

import (
	"fmt"
	"math/rand"
)

// make a new array
func make_array(n int) []int {
	data := make([]int, n)
	for i, _ := range data {
		data[i] = rand.Intn(100)
	}
	return data
}

//merge sort
func merge_sort(data []int) []int {
	if len(data) <= 1 {
		return data
	}
	m := len(data) / 2
	data1 := merge_sort(data[:m])
	data2 := merge_sort(data[m:])
	//merge
	d := []int{}
	i, j := 0, 0
	for i < len(data1) && j < len(data2) {
		if data1[i] >= data2[j] {
			d = append(d, data1[i])
			i += 1
		} else {
			d = append(d, data2[j])
			j += 1
		}
	}
	for ; i < len(data1); i += 1 {
		d = append(d, data1[i])
	}
	for ; j < len(data2); j += 1 {
		d = append(d, data2[j])
	}
	return d
}

//bubble sort
func bubble_sort(data []int) {
	if len(data) <= 1 {
		return
	}
	for i := len(data) - 1; i > 0; i -= 1 {
		for j := 0; j < i; j += 1 {
			if data[j] > data[j+1] {
				data[j], data[j+1] = data[j+1], data[j]
			}
		}
	}
}

//insert sort
func insert_sort(data []int) {
	for i := 1; i < len(data); i += 1 {
		d := data[i]
		j := 0
		for ; j <= i && data[j] < d; j += 1 {
		}
		for k := i; k > j; k -= 1 {
			data[k] = data[k-1]
		}
		data[j] = d
	}
}

//quick sort
func quick_sort(data []int) {
	q_sort(data, 0, len(data)-1)
}
func q_sort(data []int, s, e int) {
	if s >= e {
		return
	}
	p_index := rand.Intn(e-s+1) + s
	p := data[p_index]
	i := s
	j := e
	for i < j {
		for ; i < j && data[i] <= p; i += 1 {
		}
		for ; i < j && data[j] > p; j -= 1 {
		}
		if i < j {
			data[i], data[j] = data[j], data[i]
		}
	}
	data[j], data[p_index] = data[p_index], data[j]
	q_sort(data, s, j-1)
	q_sort(data, j+1, e)
}
func main() {
	n := 8
	data := make_array(n)
	data = merge_sort(data)
	fmt.Printf("%s\t%v\n", "merge_sort", data)
	data = make_array(n)
	bubble_sort(data)
	fmt.Printf("%s\t%v\n", "bubble_sort", data)
	data = make_array(n)
	insert_sort(data)
	fmt.Printf("%s\t%v\n", "insert_sort", data)
	data = make_array(n)
	quick_sort(data)
	fmt.Printf("%s\t%v\n", "quick_sort", data)
}
