package validparentheses

func IsValid(s string) bool {
	brackets := map[rune]rune{'(': ')', '[': ']', '{': '}'}
	stack := []rune{}
	for _, ch := range s {
		if val, ok := brackets[ch]; ok {
			stack = append(stack, val)
		} else {
			if len(stack) == 0 {
				return false
			}
			if stack[len(stack)-1] != ch {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}
