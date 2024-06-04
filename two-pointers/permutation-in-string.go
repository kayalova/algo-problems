package twopointers

/*  https://leetcode.com/problems/permutation-in-string/description/
* O(n*len(s2)), O(n*len(s1))
 */
func CheckInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	m := map[byte]int{}
	for i := 0; i < len(s1); i++ {
		m[s1[i]]++
	}

	for i := 0; i <= len(s2)-len(s1); i++ {
		if _, ok := m[s2[i]]; ok && areFullIntersect(m, s2[i:i+len(s1)]) {
			return true
		}
	}

	return false
}

func areFullIntersect(m map[byte]int, s string) bool {
	comp := map[byte]int{}
	for i := 0; i < len(s); i++ {
		comp[s[i]]++
	}

	for k, v := range m {
		if c, ok := comp[k]; !ok || v != c {
			return false
		}
	}

	return true
}
