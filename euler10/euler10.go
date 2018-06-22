package main

import (
	"fmt"
	"math"
)

func isPrime(n int) bool {
	if n == 2 {
		return true
	}
	var ret bool = true
	for i := 2; i < int(math.Sqrt(float64(n)))+1; i++ {
		if n%i == 0 {
			ret = false
			break
		}
	}
	return ret
}

func main() {
	var sum int
	for i := 2; i < 2000000; i++ {
		if isPrime(i) {
			sum += i
		}
	}
	fmt.Println(sum)
}
