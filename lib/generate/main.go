package main

import (
	"flag"
	"math/rand"

	"github.com/lovung/sortBigFile/lib/golib"
)

func main() {
	fo := flag.String("o", "resources/list.txt", "file path to write to")
	num := flag.Int("n", 1000000, "number of items")
	flag.Parse()
	slice := rand.Perm(*num)

	golib.WriteFile(*fo, slice, len(slice))
}
