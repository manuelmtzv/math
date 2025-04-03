package main

import (
	"fmt"
	"gomath/algebra"
	"gomath/models"
)

func main() {
	proportion := &algebra.Proportion{}

	leftSide := models.Fraction{
		Numerator:   340,
		Denominator: 3445,
	}

	rightSide := models.Fraction{
		Numerator:   20,
		Denominator: 0,
	}

	result, err := proportion.CalculateMissingVariable(leftSide, rightSide)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Result: %.8f \n", result)
}
