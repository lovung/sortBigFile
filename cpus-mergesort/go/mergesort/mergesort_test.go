package mergesort

import (
	"log"
	"math/rand"
	"testing"
	"time"
)

var testData = []struct {
	input     []int
	threshold int
}{
	//Simple Tests
	{[]int{}, -1},
	{[]int{42, 45, 78, 9, 5, 0, 0, 1, 1}, 0},
	{[]int{42, 23}, 0},

	//Tests with 1 million elements
	{rand.Perm(1000000), -1},
	{rand.Perm(1000000), 500000},
	{rand.Perm(1000000), 100000},
	{rand.Perm(1000000), 10000},
	{rand.Perm(1000000), 1000},
	{rand.Perm(1000000), 100},
	{rand.Perm(1000000), 10},

	//Tests with 100 million elements
	// {rand.Perm(100000000), -1},
	// {rand.Perm(100000000), 80000000},
	// {rand.Perm(100000000), 50000000},
	// {rand.Perm(100000000), 10000000},
}

func TestMergesort(t *testing.T) {
	for _, testCase := range testData {

		start := time.Now()

		actual := MergeSort(testCase.input, testCase.threshold)

		if !checkIfListLengthIsEqual(actual, testCase.input) || !checkIfListIsSorted(actual) {
			t.Fail()
		}

		log.Printf("Took %s to sort %d items (with threshold %d).", time.Since(start), len(testCase.input), testCase.threshold)
	}
}

func checkIfListLengthIsEqual(firstList, secondList []int) bool {

	if len(firstList) != len(secondList) {
		return false
	}
	return true
}

func checkIfListIsSorted(data []int) bool {
	last := 0
	length := len(data)

	for i := 0; i < length; i++ {

		if last > data[i] {
			return false
		}
		last = data[i]
	}

	return true
}
