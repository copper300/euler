package main

import (
	"fmt"
	"math"
)

func main() {
	var cnt int
	var i int
	for i = 2; ; i++ {
		if isPrime(i) {
			cnt++
			if cnt == 10001 {
				break
			}
		}
	}
	fmt.Println(i)
}

func isPrime(n int) bool {
	ret := true
	if n <= 1 {
		return false
	} else if n == 2 {
		return true
	} else {
		for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
			if n%i == 0 {
				ret = false
				break
			}
		}
	}
	return ret
}
