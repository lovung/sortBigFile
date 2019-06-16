package mergesort

import (
	"sync"
)

// MergeSort sort the list
func MergeSort(list []int, threshold int) []int {
	useThreshold := !(threshold < 0)
	size := len(list)
	middle := size / 2

	if size <= 1 {
		return list
	}

	var left, right []int
	sortInNewRoutine := size > threshold && useThreshold
	if !sortInNewRoutine {
		left = MergeSort(list[:middle], threshold)
		right = MergeSort(list[middle:], threshold)
	} else {
		var wg sync.WaitGroup
		wg.Add(1)

		go func() {
			defer func() { wg.Done() }()
			left = MergeSort(list[:middle], threshold)

		}()

		right = MergeSort(list[middle:], threshold)
		wg.Wait()
	}
	return mergeV2(left, right)

}

func merge(leftList, rightList []int) []int {
	size := len(leftList) + len(rightList)
	i, j := 0, 0
	slice := make([]int, size)

	for k := 0; k < size; k++ {
		if i > len(leftList)-1 && j <= len(rightList)-1 {
			slice[k] = rightList[j]
			j++
		} else if j > len(rightList)-1 && i <= len(leftList)-1 {
			slice[k] = leftList[i]
			i++
		} else if leftList[i] < rightList[j] {
			slice[k] = leftList[i]
			i++
		} else {
			slice[k] = rightList[j]
			j++
		}
	}
	return slice
}

// Merge left and right slice into newly created slice
func mergeV2(left, right []int) []int {
	size, i, j := len(left)+len(right), 0, 0
	slice := make([]int, size, size)
	count := 0

	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			slice[count] = left[i]
			count, i = count+1, i+1
		} else {
			slice[count] = right[j]
			count, j = count+1, j+1
		}
	}
	for i < len(left) {
		slice[count] = left[i]
		count, i = count+1, i+1
	}
	for j < len(right) {
		slice[count] = right[j]
		count, j = count+1, j+1
	}

	return slice
}
