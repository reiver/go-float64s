package float64s

import (
	"math"

	"testing"
)


func TestPSum(t *testing.T) {

	tests := []struct{
		Values []float64
		Expected float64
	}{
		{
			Values: []float64{},
			Expected: 0.0,
		},



		{
			Values: []float64{-math.MaxFloat64},
			Expected: -math.MaxFloat64,
		},
		{
			Values: []float64{-2.2},
			Expected: -2.2,
		},
		{
			Values: []float64{-1.1},
			Expected: -1.1,
		},
		{
			Values: []float64{-math.SmallestNonzeroFloat64},
			Expected: -math.SmallestNonzeroFloat64,
		},
		{
			Values: []float64{-0.0},
			Expected: -0.0,
		},
		{
			Values: []float64{0.0},
			Expected: 0.0,
		},
		{
			Values: []float64{math.SmallestNonzeroFloat64},
			Expected: math.SmallestNonzeroFloat64,
		},
		{
			Values: []float64{1.1},
			Expected: 1.1,
		},
		{
			Values: []float64{2.2},
			Expected: 2.2,
		},
		{
			Values: []float64{math.MaxFloat64,},
			Expected: math.MaxFloat64,
		},



		{
			Values: []float64{1.0, 10e100, 1.0, -10e100},
			Expected: 2.0,
		},
		{
			Values: []float64{-1.0, -10e100, -1.0, 10e100},
			Expected: -2.0,
		},
	}


	for testNumber, test := range tests {

		if expected, actual := test.Expected, PSum(test.Values...); expected != actual {
			t.Errorf("For test #%d, expected %g, but actually got %g.", testNumber, expected, actual)
			continue
		}
	}
}
