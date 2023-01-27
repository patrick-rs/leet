package palindrom

func LongestPalindrome(str string) string {
	if len(str) == 1 {
		return str
	}

	lSI := 0
	lEI := 0

	for si := 0; si < len(str); si++ {

		if lEI-lSI > len(str)/2 {
			return str[lSI : lEI+1]
		}
		for ei := len(str) - 1; ei > si; ei-- {

			if IsPalindrome(str[si : ei+1]) {

				if ei-si > lEI-lSI {
					lSI = si
					lEI = ei
				}
			}

		}
	}

	return str[lSI : lEI+1]
}

func IsPalindrome(s string) bool {

	if len(s) == 1 {
		return true
	}

	if len(s) == 2 {
		return s[0] == s[1]
	}

	if s[0] != s[len(s)-1] {
		return false
	}

	stringCut := s[1 : len(s)-1]

	if len(stringCut) == 1 {
		return false
	}
	return IsPalindrome(stringCut)
}
