package main

import "fmt"

func quickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	pivot := arr[len(arr) / 2]
	var less, equal, greater []int

	for _, value := range arr {
		if value < pivot {
			less = append(less, value)
		} else if value == pivot {
			equal = append(equal, value)
		} else {
			greater = append(greater, value)
		}
	}

	less = quickSort(less)
	greater = quickSort(greater)
	sorted := append(less, equal...)
	sorted = append(sorted, greater...)

	return sorted
}

func main() {
	arr := []int{5, 7, 4, 3, 9, 2, 1, 8, 6, 7}
	fmt.Println("before: ", arr)

	sortedArr := quickSort(arr)
	fmt.Println("after: ", sortedArr)
}
