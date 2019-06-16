package main

import (
	"flag"
	"fmt"
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
	largest := golib.ReadFile(*fi, &slice)
	if *verbose {
		fmt.Printf("Read: %s\n", time.Since(t))
		t = time.Now()
	}
	slice = radixSort(slice, largest)
	if *verbose {
		fmt.Printf("Sort: %s\n", time.Since(t))
		t = time.Now()
	}
	golib.WriteFile(*fo, slice, len(slice))
	if *verbose {
		fmt.Printf("Write: %s\n", time.Since(t))
	}
}

// Radix Sort
func radixSort(array []int, largest int) []int {
	// Base 10 is used
	size := len(array)
	significantDigit := 1
	semiSorted := make([]int, size, size)

	// Loop until we reach the largest significant digit
	for largest/significantDigit > 0 {
		bucket := [10]int{0}
		// Counts the number of "keys" or digits that will go into each bucket
		for i := 0; i < size; i++ {
			bucket[(array[i]/significantDigit)%10]++
		}

		// Add the count of the previous buckets
		// Acquires the indexes after the end of each bucket location in the array
		// Works similar to the count sort algorithm
		for i := 1; i < 10; i++ {
			bucket[i] += bucket[i-1]
		}

		// Use the bucket to fill a "semiSorted" array
		for i := size - 1; i >= 0; i-- {
			bucket[(array[i]/significantDigit)%10]--
			semiSorted[bucket[(array[i]/significantDigit)%10]] = array[i]
		}

		// Replace the current array with the semisorted array
		for i := 0; i < size; i++ {
			array[i] = semiSorted[i]
		}

		// Move to next significant digit
		significantDigit *= 10
	}

	return array
}
