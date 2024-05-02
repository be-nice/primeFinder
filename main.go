package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/fatih/color"
)

type Input struct {
	n   int
	seq bool
}

func main() {
	data := validateArgs()
	if !data.seq {
		isPrime := isPrime(data.n)
		if isPrime {
			fmt.Println(data.n, "is a prime number")
		} else {
			factors := findFactors(data.n)[2:]
			sort.Ints(factors)
			str := stringConversion(factors)
			cStr := color.GreenString(str)
			fmt.Println(data.n, "is not a prime number, its factors are", cStr)
		}
	} else {
		primes := sieve(data.n)
		str := stringConversion(primes)
		cStr := color.GreenString(str)
		fmt.Println("Prime numbers upto", data.n, "are", cStr)
	}
}

func validateArgs() Input {
	if len(os.Args) == 1 || len(os.Args) > 3 {
		fmt.Println("Please provide a number and an optional flag")
		os.Exit(0)
	}
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Please provide a valid number")
		os.Exit(0)
	}
	if len(os.Args) == 3 {
		if os.Args[2] != "--seq" {
			fmt.Println("Please provide a valid flag")
			os.Exit(0)
		}
		return Input{n, true}
	}
	return Input{n, false}
}

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	if n <= 3 {
		return true
	}
	if n%2 == 0 || n%3 == 0 {
		return false
	}

	for i := 5; i*i <= n; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}
	return true
}

func findFactors(number int) []int {
	factors := []int{}

	for i := 1; i*i <= number; i++ {
		if number%i == 0 {
			factors = append(factors, i)

			if i != number/i {
				factors = append(factors, number/i)
			}
		}
	}
	return factors
}

func sieve(n int) []int {
	prime := make([]bool, n+1)
	for i := 2; i <= n; i++ {
		prime[i] = true
	}

	for i := 2; i*i <= n; i++ {
		if prime[i] {
			for j := i * i; j <= n; j += i {
				prime[j] = false
			}
		}
	}

	primes := []int{}
	for i := 2; i <= n; i++ {
		if prime[i] {
			primes = append(primes, i)
		}
	}
	return primes
}

func stringConversion(factors []int) string {
	strArr := make([]string, len(factors))
	for i := 0; i < len(factors); i++ {
		strArr[i] = strconv.Itoa(factors[i])
	}
	return strings.Join(strArr, ", ")
}
