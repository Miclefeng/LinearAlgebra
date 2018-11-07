package main

import (
	"LinearAlgebra/vector"
	"fmt"
)

func main() {
	va := vector.Vectors{4, 3, 2}
	vb := vector.Vectors{5, 2, 1}

	newVec, err := vector.Add(va, vb)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("(%v) + (%v) = (%v)", va, vb, newVec)
}
