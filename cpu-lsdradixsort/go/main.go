package main

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"flag"
	"fmt"
	"os"
	"strconv"
)

// declarations for word size of data
// type word int64

const wordLen = 8
const highBit = -1 << 63

var data = []int64{170, 45, 75, -90, -802, 24, 2, 66}

// readFileToByteArray read line by line a file which have integers in every lines
func readFileToByteArray(filename string, arr *[][]byte) {
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}
	defer func() {
		if err = f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	buf := bytes.NewBuffer(nil)
	s := bufio.NewScanner(f)
	for s.Scan() {
		b := make([]byte, wordLen)
		x, _ := strconv.ParseInt(s.Text(), 10, 64)
		binary.Write(buf, binary.LittleEndian, x^highBit)
		buf.Read(b)
		*arr = append(*arr, b)
	}
	err = s.Err()
	if err != nil {
		fmt.Println(err)
	}
}

// writeFile write line by line to a file
func writeFile(filename string, arr [][]byte, len int) {
	f, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	var wo int64
	w := bufio.NewWriter(f)
	buf := bytes.NewBuffer(nil)
	for _, val := range arr {
		buf.Write(val)
		binary.Read(buf, binary.LittleEndian, &wo)
		fmt.Fprintf(w, "%d\n", wo^highBit)
	}
	err = w.Flush() // Don't forget to flush!
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	fi := flag.String("i", "resources/list.txt", "file path to read from")
	fo := flag.String("o", "out/out.txt", "file path to write from")
	// verbose := flag.Bool("v", false, "verbose")
	debug := flag.Bool("debug", false, "debug")
	flag.Parse()
	var ds [][]byte

	if *debug {
		buf := bytes.NewBuffer(nil)
		for _, x := range data {
			binary.Write(buf, binary.LittleEndian, x^highBit)
			b := make([]byte, wordLen)
			buf.Read(b)
			ds = append(ds, b)
		}
	} else {
		readFileToByteArray(*fi, &ds)
	}

	bins := make([][][]byte, 512)
	for i := 0; i < wordLen; i++ {
		for _, b := range ds {
			bins[b[i]] = append(bins[b[i]], b)
		}
		j := 0
		for k, bs := range bins {
			copy(ds[j:], bs)
			j += len(bs)
			bins[k] = bs[:0]
		}
	}

	if *debug {
		fmt.Println("original:", data)
		buf := bytes.NewBuffer(nil)
		var w int64
		for i, b := range ds {
			buf.Write(b)
			binary.Read(buf, binary.LittleEndian, &w)
			data[i] = w ^ highBit
		}
		fmt.Println("sorted:  ", data)
	} else {
		writeFile(*fo, ds, len(ds))
	}

}
