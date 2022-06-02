package reverse

func Reverse(input string) string {
	s := []rune(input)
	length := len(s)
	mid := length / 2
	for i, j := 0, length-1; i < mid; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return string(s)
}
