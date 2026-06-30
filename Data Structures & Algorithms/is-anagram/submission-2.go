func isAnagram(s string, t string) bool {
	mp := make(map[rune]int, len(s))
	mp2 := make(map[rune]int, len(s))
	if len(s) != len(t) {
		return false
	}
	for _, r := range s {
		mp[r]++
	}
	for _, r := range t {
		mp2[r]++
	}

	for _, r := range t {
		if count, ok := mp[r]; !ok || count != mp2[r] {
			return false
		}
	}
	
	return true
}
