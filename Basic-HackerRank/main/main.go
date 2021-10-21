package main

import (
	"sync"
)

type in struct {
	num int32
	oddChan chan int32
	evenChan chan int32
}

func init() {

}

func server(serverChan chan in)  {
	for input := range serverChan {
		if input.num % 2 == 0 {
			input.evenChan <- input.num
		} else {
			input.oddChan <- input.num
		}
	}
}

func countBits(n int) int {
	count := 0
	for n != 0 {
		count += n&1
		n >>= 1
	}

	return count
}

func main()  {

	var arr []int32 = []int32{234, 56, 23, 78, 89}

	odds, evens := []int32{}, []int32{}
	wg := &sync.WaitGroup{}
	wg.Add(len(arr))

	serverChan := make(chan in)
	oddChan := make(chan int32)
	evenChan := make(chan int32)

	go server(serverChan)

	go func() {
		for newOdd := range oddChan {
			odds = append(odds, newOdd)
			wg.Done()
		}
	}()
	go func() {
		for newEven := range evenChan {
			evens = append(evens, newEven)
			wg.Done()
		}
	}()
	
	for idx := 0; idx < len(arr); idx++ {
		i := idx
		serverChan <- in{arr[i], oddChan, evenChan}
	}

	wg.Wait()
}
