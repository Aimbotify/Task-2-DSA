package main

import (
	"fmt"
	"sort"
)

func waveRearrange(arr []int, x int) {
	n := len(arr)
	blockSize := 2*x + 1

	for start := 0; start < n; start += blockSize {
		end := min(start+blockSize, n)
		block := arr[start:end]
		sort.Ints(block)
		peakIndex := x - 1
		if peakIndex < len(block) {
			peak := block[peakIndex]
			left := append([]int{}, block[:peakIndex]...)
			right := append([]int{}, block[peakIndex+1:]...)
			sort.Ints(left)
			sort.Sort(sort.Reverse(sort.IntSlice(right)))
			copy(arr[start:end], append(append(left, peak), right...))
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	arr := []int{10, 20, 15, 5, 23, 90, 67, 1, 44, 100, 34}
	x := 2
	waveRearrange(arr, x)
	fmt.Println("Wave array:", arr)
}
