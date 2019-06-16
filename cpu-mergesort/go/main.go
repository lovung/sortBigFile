package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
	"github.com/lovung/sortBigFile/lib/golib"
)

func main() {
	fi := flag.String("i", "resources/list.txt", "file path to read from")
	fo := flag.String("o", "out/out.txt", "file path to write from")
	verbose := flag.Bool("v", false, "verbose")
	flag.Parse()
	var slice []int

	t := time.Now()
	golib.ReadFile(*fi, &slice)

	if *verbose {
		fmt.Printf("Read: %s\n", time.Since(t))
		t = time.Now()
	}
	// V2 gives the best result
	slice = mergeSortV2(slice)
	if *verbose {
		fmt.Printf("Sort: %s\n", time.Since(t))
		t = time.Now()
	}
	golib.WriteFile(*fo, slice, len(slice))
	if *verbose {
		fmt.Printf("Write: %s\n", time.Since(t))
	}
}

// Generates a slice of size, size filled with random numbers
func generateSlice(size int) []int {

	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(999) - rand.Intn(999)
	}
	return slice
}

func mergeV1(l, r []int) []int {
	ret := make([]int, 0, len(l)+len(r))
	for len(l) > 0 || len(r) > 0 {
		if len(l) == 0 {
			return append(ret, r...)
		}
		if len(r) == 0 {
			return append(ret, l...)
		}
		if l[0] <= r[0] {
			ret = append(ret, l[0])
			l = l[1:]
		} else {
			ret = append(ret, r[0])
			r = r[1:]
		}
	}
	return ret
}

func mergeSortV1(s []int) []int {
	if len(s) <= 1 {
		return s
	}
	n := len(s) / 2
	l := mergeSortV1(s[:n])
	r := mergeSortV1(s[n:])
	return mergeV1(l, r)
}

// MergeSort Runs MergeSort algorithm on a slice single
func mergeSortV2(slice []int) []int {

	if len(slice) < 2 {
		return slice
	}
	mid := (len(slice)) / 2
	return mergeV2(mergeSortV2(slice[:mid]), mergeSortV2(slice[mid:]))
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

func mergeSortV3(items []int) []int {
	var num = len(items)

	if num == 1 {
		return items
	}

	middle := int(num / 2)
	var (
		left  = make([]int, middle)
		right = make([]int, num-middle)
	)
	for i := 0; i < num; i++ {
		if i < middle {
			left[i] = items[i]
		} else {
			right[i-middle] = items[i]
		}
	}

	return mergeV3(mergeSortV3(left), mergeSortV3(right))
}

func mergeV3(left, right []int) (result []int) {
	result = make([]int, len(left)+len(right))

	i := 0
	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			result[i] = left[0]
			left = left[1:]
		} else {
			result[i] = right[0]
			right = right[1:]
		}
		i++
	}

	for j := 0; j < len(left); j++ {
		result[i] = left[j]
		i++
	}
	for j := 0; j < len(right); j++ {
		result[i] = right[j]
		i++
	}

	return
}
