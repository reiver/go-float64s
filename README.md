# go-float64s

Package **float64s** provides basic math functions and Numerical Analysis helpers for `float64` data, for the Go programming language.

## Summation

For example, the Associative Property that hold for ℝ (real numbers) does NOT hold for floating point number types, such as `float64`.

The order that you add a (finite) series of floating point numbers — float64 numbers — can matter!

You can get different results depending on the order you add!


## Documention

Online documentation, which includes examples, can be found at: http://godoc.org/github.com/reiver/go-float64s

[![GoDoc](https://godoc.org/github.com/reiver/go-float64s?status.svg)](https://godoc.org/github.com/reiver/go-float64s)


## Example
```
	sum := float64s.PSum(1.0, 10e100, 1.0, -10e100)
	// sum = 2.0

	sum := float64s.Sum(1.0, 10e100, 1.0, -10e100)
	// sum = 0.0
``
