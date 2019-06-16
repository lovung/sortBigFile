package main

import (
	"flag"
	"fmt"
	"github.com/lovung/sortBigFile/lib/golib"
	"math/rand"
	"time"
)

func main() {
	fi := flag.String("i", "resource/list.txt", "file path to read from")
	fo := flag.String("o", "out/out.txt", "file path to write from")
	verbose := flag.Bool("v", false, "verbose")
	flag.Parse()
	var slice []int

	t := time.Now()
	golib.ReadFile(*fi, &slice)
	fmt.Printf("Read: %s\n", time.Since(t))
	t = time.Now()
	// slice = generateSlice(1000)
	if *verbose {
		fmt.Println("\n--- Unsorted --- \n\n", slice)
	}
	quicksort(slice)
	fmt.Printf("Sort: %s\n", time.Since(t))
	t = time.Now()
	if *verbose {
		fmt.Println("\n--- Sorted ---\n\n", slice)
	}
	golib.WriteFile(*fo, slice, len(slice))
	fmt.Printf("Write: %s\n", time.Since(t))
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

func quicksort(a []int) []int {
	if len(a) < 2 {
		return a
	}

	left, right := 0, len(a)-1

	pivot := rand.Int() % len(a)

	a[pivot], a[right] = a[right], a[pivot]

	for i := range a {
		if a[i] < a[right] {
			a[left], a[i] = a[i], a[left]
			left++
		}
	}

	a[left], a[right] = a[right], a[left]

	quicksort(a[:left])
	quicksort(a[left+1:])

	return a
}

func mergesort(s []int) {
	if len(s) > 1 {
		middle := len(s) / 2
		mergesort(s[:middle])
		mergesort(s[middle:])
		merge(s, middle)
	}
}

func mergesortv3(s []int) {
	len := len(s)

	if len > 1 {
		if len <= max { // Sequential
			mergesort(s)
		} else { // Parallel
			middle := len / 2

			var wg sync.WaitGroup
			wg.Add(1)

			go func() {
				defer wg.Done()
				mergesortv3(s[:middle])
			}()

			mergesortv3(s[middle:])

			wg.Wait()
			merge(s, middle)
		}
	}
}

func main() {

}
