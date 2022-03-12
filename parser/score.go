package parser

type Calculator struct{}

func (s Calculator) Score(lastName LastName, firstName FirstName) float64 {
	v := float64(len(lastName) - len(firstName))
	if v == 0 {
		return 1
	}

	return 1 / (v * v)
}
