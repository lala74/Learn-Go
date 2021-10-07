package main

import (
	"fmt"
)

type Triplet [3]int

func Range(min, max int) []Triplet {
	triplet := []Triplet{}

	for a := min; a <= max-2; a++ {
		for b := a + 1; b <= max-1; b++ {
			for c := b + 1; c <= max; c++ {
				if a*a+b*b == c*c {
					triplet = append(triplet, Triplet{a, b, c})
				}
			}
		}
	}
	return triplet
}

func Sum(p int) []Triplet {
	triplet := []Triplet{}

	if p < 4 {
		return triplet
	}

	for a := 1; a <= p/3; a++ {
		for b := a + 1; b <= (p-a)/2; b++ {
			c := p - a - b
			if a*a+b*b == c*c {
				triplet = append(triplet, Triplet{a, b, c})
			}
		}
	}
	return triplet
}

func main() {
	defer fmt.Println("Hello")
	defer fmt.Println("!")
	fmt.Println("World")
}
