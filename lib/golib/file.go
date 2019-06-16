package golib

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// ReadFile read line by line a file which have integers in every lines
func ReadFile(filename string, arr *[]int) (largest int) {
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}
	defer func() {
		if err = f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	s := bufio.NewScanner(f)
	for s.Scan() {
		i, _ := strconv.Atoi(s.Text())
		if i > largest {
			largest = i
		}
		*arr = append(*arr, i)
	}
	err = s.Err()
	if err != nil {
		fmt.Println(err)
	}
	return largest
}

// WriteFile write line by line to a file
func WriteFile(filename string, arr []int, len int) {
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
		fmt.Println(err)
	}
}
