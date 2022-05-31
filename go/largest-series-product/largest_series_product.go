package lsproduct

import (
	"errors"
	"strconv"
)

func checkLargest(prod, largestProd int) int {
	if prod > largestProd {
		largestProd = prod
	}
	return largestProd
}

func LargestSeriesProduct(digits string, span int) (int64, error) {
	if span > len(digits) || span < 0 {
		return 0, errors.New("Span is bigger than digits length or is smaller than 0")
	}

	largestProduct := 0
	product := 1

	sliceOfIntDigits := make([]int, len(digits))

	for i, v := range digits {
		intNumber, err := strconv.Atoi(string(v))
		if err != nil {
			return 0, err
		}
		sliceOfIntDigits[i] = intNumber
	}

	if span == len(sliceOfIntDigits) {
		for i := 0; i < len(sliceOfIntDigits)-1; i++ {
			for j := 0; j < span; j++ {
				product *= sliceOfIntDigits[i+j]
			}
			largestProduct = checkLargest(product, largestProduct)
			product = 1
		}
	}

	for i := 0; i <= len(sliceOfIntDigits)-span; i++ {
		for j := 0; j < span; j++ {
			product *= sliceOfIntDigits[i+j]
		}
		largestProduct = checkLargest(product, largestProduct)
		product = 1
	}
	return int64(largestProduct), nil
}
