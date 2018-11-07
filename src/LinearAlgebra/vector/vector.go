package vector

import "errors"

//type Vector interface {
//	Add(va, vb Vectors) (Vectors, error)
//	Sub(va, vb Vectors) (Vectors, error)
//	Mul(k int, v Vectors) Vectors
//	Div(k int, v Vectors) Vectors
//}

type Vectors []int

func Add(va, vb Vectors) (Vectors, error) {
	if len(va) != len(vb) {
		return nil, errors.New("The sum of the two vectors must be the same length.")
	}

	var newVectors []int
	for i, _ := range va {
		newVectors = append(newVectors, va[i]+vb[i])
	}

	return newVectors, nil
}

func Sub(va, vb Vectors) (Vectors, error) {
	if len(va) != len(vb) {
		return nil, errors.New("The two vectors that you subtract have to be the same length.")
	}

	var newVectors Vectors
	for i, _ := range va {
		newVectors = append(newVectors, va[i]-vb[i])
	}

	return newVectors, nil
}

func Mul(k int, vec Vectors) Vectors {

	var newVectors Vectors
	for i, _ := range vec {
		newVectors = append(newVectors, k*vec[i])
	}

	return newVectors
}

func Div(k int, vec Vectors) Vectors {

	var newVectors Vectors
	for i, _ := range vec {
		newVectors = append(newVectors, vec[i]/k)
	}

	return newVectors
}
