package main

import (
	"fmt"
)

func main() {
	var a, b, c int
	found := 0
	for a = 1; a < 1000; a++ {
		for b = 1; b < 1000; b++ {
			c = 1000 - a - b
			if a*a+b*b == c*c {
				found = 1
				break
			}
		}
		if found == 1 {
			break
		}
	}
	fmt.Printf("a=%d, b=%d, c=%d, abc=%d\n", a, b, c, a*b*c)
}
