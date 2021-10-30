package main

func getMoneySpent(keyboards []int, drives []int, b int) int {
	/*
	 * Write your code here.
	 */

	maximum := 0
	sum := 0
	for i := 0; i < len(keyboards); i ++ {
		for j := 0; j < len(drives); j++ {
			sum = keyboards[i] + drives[j]
			if sum >= maximum && sum <= b {
				maximum = sum
			}
		}
	}
	if maximum == 0 {
		return -1
	}

	return maximum
}
