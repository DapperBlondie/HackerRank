package main

import "fmt"

func bonAppetit(bill []int32, k int32, b int32) {
	// Write your code here
	var sum int32 = 0

	for i := 0; i < len(bill); i += 1 {
		if i == int(k) {
			continue
		}
		sum += bill[i]
	}
	average := sum / 2

	if average == b {
		fmt.Println("Bon Appetit")
	} else {
		fmt.Println(b - average)
	}

}

func main()  {

	return
}
