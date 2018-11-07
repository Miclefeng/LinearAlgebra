package main

import (
	"LinearAlgebra/vector"
	"fmt"
)

func main() {
	va := vector.Vectors{4, 6, 2}
	vb := vector.Vectors{5, 2, 1}

	addVec, err := vector.Add(va, vb)
	if err != nil {
		fmt.Println(err)
		return
	}
	a, b := 3, 2
	fmt.Printf("(%v) + (%v) = (%v)\n", va, vb, addVec)
	subVec, _ := vector.Sub(va, vb)
	fmt.Printf("(%v) - (%v) = (%v)\n", va, vb, subVec)
	mulVec := vector.Mul(a, vb)
	fmt.Printf("(%v) * %d = (%v)\n", va, a, mulVec)
	divVec := vector.Div(b, va)
	fmt.Printf("(%v) / %d = (%v)\n", va, b, divVec)
}
