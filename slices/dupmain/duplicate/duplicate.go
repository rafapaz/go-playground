package duplicate

func RemoveDup(s []string) []string {
	var i int
	var lens int

	lens = len(s)

	for i = 0; i < lens; i++ {
		if i == 0 {
			continue
		}

		if i >= len(s) {
			break
		}

		if s[i] == s[i-1] {
			s = append(s[:i], s[i+1:]...)
			i--
		}
	}

	return s
}
