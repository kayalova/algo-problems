package twopointers

func Trap(height []int) int {
	suffixes := make([]int, len(height))
	prefixes := make([]int, len(height))

	suffixes[0] = 0
	prefixes[len(height)-1] = 0

	for i := 1; i < len(height); i++ {
		suffixes[i] = max(suffixes[i-1], height[i-1])
	}

	for i := len(height) - 2; i >= 0; i-- {
		prefixes[i] = max(prefixes[i+1], height[i+1])
	}

	sum, wh := 0, 0

	for i := 1; i < len(height)-1; i++ {
		if suffixes[i] == 0 || prefixes[i] == 0 {
			continue
		}

		wh = min(suffixes[i], prefixes[i]) - height[i]
		if wh > 0 {
			sum += wh
		}
	}

	return sum

}
