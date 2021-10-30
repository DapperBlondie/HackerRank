package main

func countingValleys(steps int, path string) int {
	// Write your code here
	sea := 0
	counter := 0
	var U uint8 = 'U'
	var D uint8 = 'D'

	for i := 0; i < steps; i++ {
		if sea == 0 {
			 counter += 1
		}
		if U == path[i] {
			sea += 1
		} else if D == path[i] {
			sea -= 1
		}
	}

	return counter - 1
}

func main()  {

}
