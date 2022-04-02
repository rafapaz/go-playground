package rotate

func Rotate(s []int, i int) []int {
	bef := s[:i]
	after := s[i:]
	return append(after, bef...)
}
