package main

import (
	"fmt"
	"math"
	"sort"
)

func diagonalDifference(arr [][]int32) int32 {
	// Write your code here
	var sumRow int32 = 0
	var sumRow2 int32 = 0
	for x, l := range arr {
		for y, _ := range l{
			if x == y {
				sumRow += arr[x][y]
			}
			if x == len(l) - y - 1 {
				sumRow2 += arr[x][y]
			}
		}
	}

	return int32(math.Abs(float64(sumRow - sumRow2)))
}

func miniMaxSum(arr []int) {
	// Write your code here
	sort.Ints(arr)
	var minSum int64 = 0
	var maxSum int64 = 0

	for i := 0; i<4; i++ {
		minSum += int64(arr[i])
	}

	for i := 1; i < 5; i++ {
		maxSum += int64(arr[i])
	}

	fmt.Printf("%d %d", minSum, maxSum)
}

func main() {

}
