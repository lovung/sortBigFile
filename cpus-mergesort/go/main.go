package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"

	"github.com/lovung/sortBigFile/cpus-mergesort/go/mergesort"
	"github.com/lovung/sortBigFile/lib/golib"
)

func main() {
	fi := flag.String("i", "resources/list.txt", "file path to read from")
	fo := flag.String("o", "out/out.txt", "file path to write from")
	verbose := flag.Bool("v", false, "verbose")
	debug := flag.Bool("debug", false, "debug")
	threshold := flag.Int("th", 10000, "threshold")
	flag.Parse()
	var slice []int

	t := time.Now()
	if *debug {
		slice = rand.Perm(10)
	} else {
		golib.ReadFile(*fi, &slice)
	}
	if *verbose {
		fmt.Printf("Read: %s\n", time.Since(t))
		t = time.Now()
		if *debug {
			fmt.Println(slice)
		}
	}
	slice = mergesort.MergeSort(slice, *threshold)
	if *verbose {
		fmt.Printf("Sort: %s\n", time.Since(t))
		t = time.Now()
	}

	if !(*debug) {
		golib.WriteFile(*fo, slice, len(slice))
	}
	if *verbose {
		fmt.Printf("Write: %s\n", time.Since(t))
		if *debug {
			fmt.Println(slice)
		}
	}
}
