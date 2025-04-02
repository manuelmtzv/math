package main

import "gomath/algebra"

func main() {
	proportion := &algebra.Proportion{}

	result, err := proportion.MissingValue(1, 2, 0, 26)
	if err != nil {
		panic(err)
	}

	println("Result:", result)
}
