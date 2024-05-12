package twopointers

/* 6. https://leetcode.com/problems/container-with-most-water/
O(n), O(1)
*/

func MaxArea(height []int) int {
	maxArea := 0
	left, right := 0, len(height)-1
	x, y := 0, 0

	for left < right {
		x = right - left

		if height[right] < height[left] {
			y = height[right]
			right--
		} else {
			y = height[left]
			left++
		}

		if maxArea < x*y {
			maxArea = x * y
		}
	}

	return maxArea
}
