package main

import "fmt"

// 阶乘
// n!
// n=0 n!=1
// n>=1 n! = n*(n-1)!
// f(n) = n*f(n-1)
func factorial(n int) int {
	if n <= 1 {
		return 1
	} else {
		return factorial(n-1) * n
	}

}
func main() {
	fmt.Println(factorial(3))
}
