package golib

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// ReadFile read line by line a file which have integers in every lines
func ReadFile(filename string, arr *[]int) {
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
		*arr = append(*arr, i)
	}
	err = s.Err()
	if err != nil {
		fmt.Println(err)
	}
}

var concurrency = 100

// ConcurrentReadFile read line by line a file which have integers in every lines
func ConcurrentReadFile(filename string, arr *[]int) {
	workQueue := make(chan string)
	complete := make(chan bool)

	go func() {
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
			workQueue <- s.Text()
		}
		err = s.Err()
		if err != nil {
			fmt.Println(err)
		}
		close(workQueue)
	}()

	// Now read them all off, concurrently.
	for i := 0; i < concurrency; i++ {
		go startWorking(workQueue, complete, arr)
	}

	// Wait for everyone to finish.
	for i := 0; i < concurrency; i++ {
		<-complete
	}
}

func startWorking(queue chan string, complete chan bool, arr *[]int) {
	for line := range queue {
		// Do the work with the line.
		i, _ := strconv.Atoi(line)
		*arr = append(*arr, i)
	}

	// Let the main process know we're done.
	complete <- true
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
	for _, val := range arr {
		fmt.Fprintf(w, "%d\n", val)
	}
	err = w.Flush() // Don't forget to flush!
	if err != nil {
		fmt.Println(err)
	}
}
