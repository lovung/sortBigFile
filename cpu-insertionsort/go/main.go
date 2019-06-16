package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func readFile(filename string, arr *[]int) {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = f.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	s := bufio.NewScanner(f)
	for s.Scan() {
		i, _ := strconv.Atoi(s.Text())
		*arr = append(*arr, i)
	}
	err = s.Err()
	if err != nil {
		log.Fatal(err)
	}
}

func writeFile(filename string, arr []int, len int) {
	f, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	w := bufio.NewWriter(f)
	for _, i := range arr {
		fmt.Fprintf(w, "%d\n", i)
	}
	err = w.Flush() // Don't forget to flush!
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	fi := flag.String("i", "resources/list.txt", "file path to read from")
	fo := flag.String("o", "out/out.txt", "file path to write from")
	verbose := flag.Bool("v", false, "verbose")
	flag.Parse()
	var slice []int

	t := time.Now()
	readFile(*fi, &slice)
	fmt.Printf("Read: %s\n", time.Since(t))
	t = time.Now()
	// slice = generateSlice(1000)
	if *verbose {
		fmt.Println("\n--- Unsorted --- \n\n", slice)
	}
	InsertionSort(slice)
	fmt.Printf("Sort: %s\n", time.Since(t))
	t = time.Now()
	if *verbose {
		fmt.Println("\n--- Sorted ---\n\n", slice)
	}
	writeFile(*fo, slice, len(slice))
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

// InsertionSort Difference with bubble sort, Here at any iteration of outerloop,
// the array to the left of the element will be sorted
func InsertionSort(numbers []int) []int {
	for i := 0; i < len(numbers); i++ {
		for j := 0; j < i+1; j++ {
			//compare element present at index i with every element present
			//  left of it place it in right place so that array on the
			//left remains   sorted
			if numbers[j] > numbers[i] {
				intermediate := numbers[j]
				numbers[j] = numbers[i]
				numbers[i] = intermediate
			}
		}
	}
	return numbers
}
