package hamming

import "errors"

func Distance(a, b string) (int, error) {
	var hammingDistance int
	if len(a) == len(b) {
		for i := 0; i < len(a); i++ {
			if a[i] != b[i] {
				hammingDistance += 1
			}
		}
		return hammingDistance, nil
	}
	return 0, errors.New("a and b have different lengths")
}
