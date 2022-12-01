package utils

// StringContains returns true if s is contained in a, and false otherwise
func StringContains(a string, s rune) (found bool) {
	for _, x := range a {
		if x == s {
			found = true
		}
	}
	return found
}

func StringSliceContains(a []string, s string) (found bool) {
	for _, x := range a {
		if x == s {
			found = true
		}
	}
	return found
}

func GetStringMapKey(m map[string]string, value string) string {
	for k, v := range m {
		if v == value {
			return k
		}
	}
	return ""
}
