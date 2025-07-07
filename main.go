//Name: Eduardo Biscaia de Queiroz
//Date: 03/07/2025

//Objectives:
//Create a program that calculates the prime numbers under a given number
//Factorize the number, giving all the prime divisors, prime divisors and
//their repetition or all the divisors of a given number using golang flags

//Examples:
//--all-primes 50
//[2,3,5,7,11,13,17,19,23,29,31,37,41,43,47]

//--factors 50
//[2,5,5]

//--factor-pairs 50
//[(2,1),(5,2)]

//--divisors 50
//[1,2,5,10,50]

package main

import "fmt"

func main() {
	number := 50
	divisor := 11
	fmt.Printf("%v\n", number)
	checkPrime(divisor)
}

func checkPrime(number int) bool {
	//Check if a number is a prime number
	isPrime := false
	if number < 2 {
		fmt.Printf("%v is not a prime number\n", number)
	} else {
		for testNumber := 2; number >= testNumber; testNumber++ {
			if number%testNumber == 0 {
				if number == testNumber {
					fmt.Printf("%v is a prime number\n", number)
				} else {
					fmt.Printf("%v is not a prime number\n", number)
					break
				}
			}
		}
	}
}
