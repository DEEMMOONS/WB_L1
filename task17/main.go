package main

import (
	"fmt"
	"sort"
)

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	target := 7

	index := sort.Search(len(slice), func(i int) bool {
		return slice[i] >= target
	})

	if index < len(slice) && slice[index] == target {
		fmt.Printf("Значение %d найдено в срезе, индекс %d\n", target, index)
	} else {
		fmt.Printf("Значение %d не найдено в срезе\n", target)
	}
}
