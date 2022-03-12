package parser

type Calculator struct{}

const orderOnlyScoreLength = 4

// Score referer: https://github.com/rskmoi/namedivider-python/blob/master/namedivider/name_divider.py#L206
func (s Calculator) Score(lastName LastName, firstName FirstName) float64 {
	fullname := JoinName(lastName, firstName)
	ols := s.orderScore(string(lastName), fullname.Length(), 0)
	ofs := s.orderScore(string(firstName), fullname.Length(), lastName.Length())
	os := (ols + ofs) / (float64(fullname.Length()) - minNameLength)

	// https://github.com/rskmoi/namedivider-python/blob/d87a488d4696bc26d2f6444ed399d83a6a1911a7/namedivider/name_divider.py#L219
	if fullname.Length() == orderOnlyScoreLength {
		return os
	}

	lls := s.lengthScore(string(lastName), fullname.Length(), 0)
	lfs := s.lengthScore(string(firstName), fullname.Length(), lastName.Length())
	ls := (lls + lfs) / float64(fullname.Length())

	return ls
}

// orderScore: patch work implementation.
func (s Calculator) orderScore(name string, fullNameLength, _ int) float64 {
	v := float64(len(name) - fullNameLength)
	if v == 0 {
		return 1
	}

	return 1 / (v * v)
}

// lengthScore: patch work implementation.
func (s Calculator) lengthScore(name string, fullNameLength, _ int) float64 {
	v := float64(len(name) - fullNameLength)
	if v == 0 {
		return 1
	}

	return 1 / (v * v)
}
