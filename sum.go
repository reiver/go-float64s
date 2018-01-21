package float64s

// Sum sums the floating point numbers in the []float64, in the order they are in.
//
// Note that although this is fast, there can be errors and loss of precision with this operation.
//
// In general, floating point number do not have the Associative Property that holds for ℝ (real numbers).
//
// For example:
//
//	var f64s []float64 = []float64{1.0, 10e100, 1.0, -10e100}
//	
//	sum := float64s..Sum(f64s...)
//
// The value of `sum` will be 0.0, not 2.0!
//
// If you prefer (potential) increased accuracy over speed, use the slower but more accurate float64s.PSum() func.
func Sum(p ...float64) float64 {

	var accumulator float64 = 0.0
	for _, datum := range p {
		accumulator += datum
	}

	return accumulator
}
