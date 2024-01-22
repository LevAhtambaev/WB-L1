package main

import (
	"fmt"
)

func main() {
	arr := []int{9, 7, 5, 11, 12, 2, 14, 3, 10, 6, 0, 13131, -45}
	QuickSort(arr)
	fmt.Println(arr)
}

// QuickSort реализация быстрой сортировки - массив разделяется на две части: элементы, меньшие опорного, и элементы, большие опорного, затем рекурсивно применяется алгоритм к обеим половинам. В итоге, массив будет отсортирован в возрастающем порядке.
func QuickSort(arr []int) {
	if len(arr) <= 1 {
		return
	}

	pivot := arr[len(arr)-1]
	left := 0
	right := len(arr) - 2

	for left <= right {
		for left <= right && arr[left] < pivot {
			left++
		}

		for left <= right && arr[right] > pivot {
			right--
		}

		if left <= right {
			arr[left], arr[right] = arr[right], arr[left]
			left++
			right--
		}
	}

	arr[left], arr[len(arr)-1] = arr[len(arr)-1], arr[left]
	QuickSort(arr[:left])
	QuickSort(arr[left+1:])
}
