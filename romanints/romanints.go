package romanints

func RomanToInt(s string) int {

	romans := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	var result = romans[string(s[len(s)-1])]

	for i := len(s) - 1; i > 0; i-- {
		if romans[string(s[i-1])] >= romans[string(s[i])] {
			result += romans[string(s[i-1])]
		} else {
			result -= romans[string(s[i-1])]
		}
	}

	return result
}
