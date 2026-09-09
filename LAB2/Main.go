package main

import (
	"Myproject/Mathutil"
	"Myproject/strop"
	"fmt"
)

func main() {
	// take all input from user
	var fact int
	fmt.Print("Enter a number to find factorial: ")
	fmt.Scan(&fact)
	fmt.Println("Factorial of", fact, "is", Mathutil.Factorial(fact))

	var base, exponent int
	fmt.Print("Enter base and exponent to find power: ")
	fmt.Scan(&base, &exponent)
	fmt.Println(base, "raised to the power", exponent, "is", Mathutil.Pow(exponent, base))

	var str string
	fmt.Print("Enter a string to reverse: ")
	fmt.Scan(&str)
	fmt.Println("Reversed string is:", strop.Reverse(str))

	var str1 string
	fmt.Print("Enter a string to count vowels: ")
	fmt.Scan(&str1)
	fmt.Println("Number of vowels in the string is:", strop.CountVowels(str1))
	// fmt.Println(Mathutil.factorial(5))
	// fmt.Println(Mathutil.power(3, 2))
	// fmt.Println(Mathutil.factorial(3, 2))
}
