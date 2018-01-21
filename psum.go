package float64s

import (
	"sort"
)

// PSum is picky about how it does the sum, and does the sum in a way
// where it tries to minimize errors.
//
// float64s.PSum() is slower than float64s.Sum(), but float64s.PSum() tries
// to reduce the errors made (where float64s.Sum() doesn't).
func PSum(p ...float64) float64 {
	if 0 == len(p) {
		return 0.0
	}

	var negativesBootstrap [8]float64
	var negatives []float64 = negativesBootstrap[:0]

	var positivesBootstrap [8]float64
	var positives []float64 = positivesBootstrap[:0]

	for _, datum := range p {
		switch {
		case 0.0 == datum || -0.0 == datum:
			// Nothing here.
		case 0.0 < datum:
			positives = append(positives, datum)
		case 0.0 > datum:
			negatives = append(negatives, datum)
		}
	}


	sort.Float64s(negatives)
	sort.Float64s(positives)


	var negativesAccumulator float64
	var negativesMissing   []float64

	for _, datum := range negatives {
		x, y := negativesAccumulator, datum
		if x > y {
			x, y = y, x
		}

		newAccumulator := x + y
		newMissing     := y - (newAccumulator - x)

		negativesAccumulator = newAccumulator
		if 0.0 != newMissing {
			negativesMissing = append(negativesMissing, newMissing)
		}
	}


	var positivesAccumulator float64
	var positivesMissing   []float64

	for _, datum := range positives {
		x, y := positivesAccumulator, datum
		if x < y {
			x, y = y, x
		}

		newAccumulator := x + y
		newMissing     := y - (newAccumulator - x)

		positivesAccumulator = newAccumulator
		if 0.0 != newMissing {
			positivesMissing = append(positivesMissing, newMissing)
		}
	}

	sort.Float64s(negativesMissing)
	sort.Float64s(positivesMissing)

	result := (positivesAccumulator + negativesAccumulator) + (Sum(negativesMissing...) + Sum(positivesMissing...))

	return result
}
