package main

func minSubArrayLen(target int, nums []int) int {
	currentSumCount := 0
	sum, best, left := 0, 0, 0

	for i := 0; i < len(nums); i++ {
		currentSumCount += 1
		sum += nums[i]
		if sum >= target {
			if currentSumCount < best || best == 0 {
				best = currentSumCount
			}

			for j := left; j <= i; j++ {
				sum -= nums[j]
				currentSumCount--
				if sum < target {
					if currentSumCount+1 < best {
						best = currentSumCount + 1
					}
				}
			}
			left = i + 1
			sum, currentSumCount = 0, 0
		}
	}

	return best
}

func main() {
	minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3})
}

/*
You're given a method that returns a random number in a range [0, 4),
create a method that returns a random in [0,3) range
*/
