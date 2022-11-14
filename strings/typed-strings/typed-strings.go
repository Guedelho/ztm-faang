package typedstrings

func BackspaceCompare(s, t string) bool {
	lastS := len(s) - 1
	lastT := len(t) - 1
	for lastS >= 0 || lastT >= 0 {
		lastS = getLast(s, lastS)
		lastT = getLast(t, lastT)
		if lastS == -1 && lastT == -1 {
			return true
		}
		if lastS == -1 || lastT == -1 {
			return false
		}
		if s[lastS] != t[lastT] {
			return false
		}
		lastS -= 1
		lastT -= 1
	}
	return true
}

func getLast(s string, last int) int {
	back := 0
	for last >= 0 {
		if s[last] == '#' {
			back += 1
		} else {
			if back == 0 {
				return last
			}
			back -= 1
		}
		last -= 1
	}
	return -1
}
