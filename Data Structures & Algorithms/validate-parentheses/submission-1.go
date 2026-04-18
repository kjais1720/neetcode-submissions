func isValid(s string) bool {
    brackets := make([]string, 0, len(s))
	for _, _b := range s {
		b := string(_b)
		n := len(brackets)
		if strings.Contains("({[", b) {
			brackets = append(brackets, b)
		} else if b == ")" {
			if n==0 || brackets[n-1] != "(" {
				return false
			} else {
				brackets = brackets[:n-1]
			}
		} else if b == "}" {
			if n==0 || brackets[n-1] != "{" {
				return false
			} else {
				brackets = brackets[:n-1]
			}
		} else if b == "]" {
			if n==0 || brackets[n-1] != "[" {
				return false
			} else {
				brackets = brackets[:n-1]
			}
		}
	}
	if len(brackets) != 0 {
		return false
	}
	return true
}
