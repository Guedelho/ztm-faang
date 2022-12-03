package removevalidparentheses

func MinRemoveToMakeValid(s string) string {
	stack := []int{}
	for i := 0; i < len(s); i++ {
		ch := s[i]
		if ch == '(' {
			stack = append(stack, i)
		} else if ch == ')' {
			if len(stack) == 0 {
				s = s[:i] + s[i+1:]
				i -= 1
			} else {
				stack = stack[:len(stack)-1]
			}
		}
	}
	for len(stack) != 0 {
		i := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		s = s[:i] + s[i+1:]
	}
	return s
}
