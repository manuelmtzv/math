package algebra

import (
	"fmt"
	"gomath/models"
)

type Proportion struct {
}

// CalculateMissingVariable solves the proportion equation a/b = c/d where one variable is zero.
// It returns the calculated value as float64 or an error if the calculation is not possible.
func (p *Proportion) CalculateMissingVariable(a, b models.Fraction) (float64, error) {
	zeroNumerators := 0
	zeroDenominators := 0

	if a.Numerator == 0 {
		zeroNumerators++
	}
	if b.Numerator == 0 {
		zeroNumerators++
	}
	if a.Denominator == 0 {
		zeroDenominators++
	}
	if b.Denominator == 0 {
		zeroDenominators++
	}

	if zeroNumerators > 1 {
		return 0, fmt.Errorf("cannot have more than one zero numerator")
	}
	if zeroDenominators > 1 {
		return 0, fmt.Errorf("cannot have more than one zero denominator")
	}
	if zeroNumerators == 0 && zeroDenominators == 0 {
		return 0, fmt.Errorf("no missing variable (zero value) found")
	}

	switch {
	case a.Numerator == 0:
		return a.Denominator * b.Numerator / b.Denominator, nil

	case b.Numerator == 0:
		return a.Numerator * b.Denominator / a.Denominator, nil

	case a.Denominator == 0:
		return a.Numerator * b.Denominator / b.Numerator, nil

	case b.Denominator == 0:
		return a.Denominator * b.Numerator / a.Numerator, nil

	default:
		return 0, fmt.Errorf("no missing variable found")
	}
}

func (p *Proportion) Demostration() {
	fmt.Println("Proportion Demonstration")
	fmt.Println("a/b = c/d")

	fmt.Println("Left Side: 1/1.29")
	leftSide := models.Fraction{
		Numerator:   1,
		Denominator: 1.29,
	}

	fmt.Println("Right Side: 0/1")
	rightSide := models.Fraction{
		Numerator:   0,
		Denominator: 1,
	}

	result, err := p.CalculateMissingVariable(leftSide, rightSide)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Result: %.8f \n", result)
}
