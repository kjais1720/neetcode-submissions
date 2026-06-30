func isAnagram(s string, t string) bool {
	mp := make(map[rune]int, len(s))
	mp2 := make(map[rune]int, len(s))
	
	for _, r := range s {
		mp[r]++
	}
	for _, r := range t {
		mp2[r]++
	}
	isAnagram := true
	for _, r := range t {
		if count, ok := mp[r]; !ok || count != mp2[r] {
			isAnagram = false
		}
	}
	for _, r := range s {
		if count, ok := mp2[r]; !ok || count != mp[r] {
			isAnagram = false
		}
	}
	return isAnagram
}
