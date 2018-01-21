/*
Package float64s provides basic math functions and Numerical Analysis helpers for float64 data.


Summation

For example, the Associative Property that hold for ℝ (real numbers) does NOT hold for floating point number types, such as float64.

The order that you add a (finite) series of floating point numbers — float64 numbers — can matter!

You can get different results depending on the order you add!


Suggestion

You may find this pattern useful:

	var bootstrap [8]float64 // 64 bytes is the length of a cache line on some computer architectures.
	                         // 64 bytes = 8 × (64 bits ÷ 8 bits per byte) = 8 × (8 bytes)
	var f64s       []float64
	
	f64s = bootstrap[:0] // Make "bootstrap" (which is allocated on the stack) the backing array of "f64s".

It can allow you to try to allocate the backing array on the stack (rather than on the heap)
so as to try to reduce (or prevent) garbage collector (GC) pressure.
*/
package float64s

