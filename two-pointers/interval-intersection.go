package twopointers

/* https://leetcode.com/problems/interval-list-intersections/
* O(n), O(n)
 */
func IntervalIntersection(firstList [][]int, secondList [][]int) [][]int {
	res := [][]int{}

	for i, j := 0, 0; i < len(firstList) && j < len(secondList); {
		a := max(firstList[i][0], secondList[j][0])
		b := min(firstList[i][1], secondList[j][1])
		if a > b {
			if a == firstList[i][0] {
				j++
			} else {
				i++
			}
			continue
		}

		res = append(res, []int{a, b})
		if b == firstList[i][1] {
			i++
		} else {
			j++
		}
	}

	return res
}
