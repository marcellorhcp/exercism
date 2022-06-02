package prime

import "math"

//IsPrime check if a number is prime
func IsPrime(number int) bool {
	if number < 2 {
		return false
	}
	is := true
	max := math.Sqrt(float64(number))
	for i := 2; i <= int(max); i++ {
		if number%i == 0 {
			is = false
			break
		}
	}
	return is
}

func Factors(n int64) []int64 {
	primeFactors := make([]int64, 0)
	var i int64 = 2
	for n > 1 {
		if IsPrime(int(i)) && n%i == 0 {
			n /= i
			primeFactors = append(primeFactors, i)
		} else {
			i++
		}
	}
	return primeFactors
}
