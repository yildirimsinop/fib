/*
package fib

func Fib(n int) int {
	return 0
}


// Bør være ok for det fjerde leddet {3,2}
package fib

func Fib(n int) int {
	if n == 0 {
		return 0
	}
	if n <= 2 {
		return 1
	}
	return 2
}
*/
// 2 er egentlig 1 + 1
// fra formelen F(3) = F(3-1) + F(3-2)
// => F(3) = F(2) + F(1)
// => F(3) = 1 + 1 = 2
package fib

func Fib(n int) int {
	if n == 0 {
		return 0
	}
	if n <= 2 {
		return 1
	}
	return Fib(n-1) + Fib(n-2)
}
