package sort

import (
	"sync"
)

type Interface interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

func ParallelQuicksort(a Interface) {
	var wg sync.WaitGroup
	var qsort func(int, int)
	partition := func(left, right int) int {
		pivot := right
		storeIndex := left
		for i := left; i < right; i++ {
			if a.Less(i, pivot) {
				a.Swap(i, storeIndex)
				storeIndex++
			}
		}
		a.Swap(storeIndex, pivot)
		return storeIndex
	}
	qsort = func(p, r int) {
		defer wg.Done()
		if p < r {
			wg.Add(2)
			q := partition(p, r)
			go qsort(p, q-1)
			go qsort(q+1, r)
		}
	}
	wg.Add(1)
	qsort(0, a.Len()-1)
	wg.Wait()
}
