package main

import "math"

func CatAndMouse(x int32, y int32, z int32) string {
	if math.Abs(float64(x-z)) < math.Abs(float64(y-z)) {
		return "Cat A"
	} else if math.Abs(float64(x-z)) > math.Abs(float64(y-z)) {
		return "Cat B"
	} else {
		return "Mouse C"
	}
}

func main() {

}
