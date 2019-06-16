package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"sync"
)

// ReadAndVerifyFile read and verify line by line a file which have integers in every lines
func ReadAndVerifyFile(filename string) bool {
	var once sync.Once
	var last int
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
		once.Do(func() {
			last = i
		})
		if i < last {
			return false
		}
		last = i
	}
	err = s.Err()
	if err != nil {
		fmt.Println(err)
	}
	return true
}

func main() {
	fi := flag.String("i", "resources/list.txt", "file path to read from")
	flag.Parse()
	if ReadAndVerifyFile(*fi) {
		fmt.Println("Valid ascending array")
		os.Exit(0)
	}
	fmt.Println("Invalid ascending array")
	os.Exit(1)
}
