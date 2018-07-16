package main

import (
	"fmt"
)

func main() {
	p, q := 5, 2
	fmt.Println(fmt.Sprintf("%d and %d produce: %d", p, q, mirrorReflection(p, q)))
}

func mirrorReflection(p int, q int) int {
	for p%2 == 0 && q%2 == 0 {
		p /= 2
		q /= 2
	}
	if p%2 == 0 {
		return 2
	} else if q%2 == 0 {
		return 0
	} else {
		return 1
	}
}
