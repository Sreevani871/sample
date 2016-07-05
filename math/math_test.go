package math

import "testing"

type Test struct {
	values  []float64
	average float64
}

var test = []Test{
	{values: []float64{1, 2}, average: 1.5},
	{values: []float64{1, 2, 3}, average: 2},
	{values: []float64{4, 6}, average: 5},
}

func TestAverage(t *testing.T) {
	for _, pair := range test {
		v := Average(pair.values)
		if v != pair.average {
			t.Error(
				"For", pair.values,
				"Expected", pair.average,
				"Got", v,
			)
		}
	}
}
