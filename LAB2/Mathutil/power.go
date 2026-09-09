// power.go
package Mathutil

func Pow(exponent, base int) int {
	result := 1
	for i := 0; i < exponent; i++ {
		result *= base
	}
	return result
}
