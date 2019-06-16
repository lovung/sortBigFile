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
	quicksort(slice)
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
