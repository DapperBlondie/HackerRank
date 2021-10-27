package main

func pageCount(n int, p int) int {
	// Write your code here
	if n == p {
		return 0
	} else if 1 == p {
		return 0
	}

	fromFirst := p / 2
	fromEnd := (n / 2) - fromFirst

	if fromEnd < fromFirst {
		return fromEnd
	} else {
		return fromFirst
	}
}

func main()  {

}
