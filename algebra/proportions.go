package algebra

import "fmt"

type Proportion struct {
}

func (p *Proportion) MissingValue(n1, d1, n2, d2 float64) (float64, error) {
	if n1 == 0 || d1 == 0 {
		return 0.0, fmt.Errorf("n1 and d1 cannot be 0")
	}

	if n2 == 0 {
		return d2 * n1 / d1, nil
	}

	if d2 == 0 {
		return n2 * d1 / n1, nil
	}

	return 0.0, nil
}
