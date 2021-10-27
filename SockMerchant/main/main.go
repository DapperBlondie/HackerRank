package main

import "sort"

func sockMerchant(n int, ar []int) int {
	// Write your code here
	sort.Ints(ar)

	counter := 0
	j := 0
	for  {
		if j <= n - 1 {
			if ar[j] == ar[j + 1] {
				counter += 1
				j += 2
			} else {
				j += 1
			}
		} else {
			break
		}
	}

	return counter
}

func main() {

}