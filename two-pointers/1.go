package twopointers

import (
	"sort"
	"strings"
	"unicode"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

/* 1. https://leetcode.com/problems/merge-two-sorted-lists/description/
* Временная сложность и пространственная сложность - O(n), O(n)
 */
func MergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil && list2 == nil {
		return nil
	}

	res := &ListNode{}
	head := res
	for list1 != nil || list2 != nil {
		if list1 == nil {
			res.Val = list2.Val
			res.Next = list2.Next
			break
		}

		if list2 == nil {
			res.Val = list1.Val
			res.Next = list1.Next
			break
		}

		if list1.Val <= list2.Val {
			res.Val = list1.Val
			list1 = list1.Next
		} else {
			res.Val = list2.Val
			list2 = list2.Next
		}

		res.Next = &ListNode{}
		res = res.Next
	}

	return head
}

/* 2. https://leetcode.com/problems/two-sum/
* O(nlogn), O(n)
* Выглядит как будто не очень, но тесты прошли
 */
func TwoSum(nums []int, target int) []int {
	m := map[int][]int{}

	for pos, n := range nums {
		m[n] = append(m[n], pos)
	}

	sort.Ints(nums)

	for i, j := 0, 1; i < len(nums); {
		if j == len(nums) {
			i++
			j = i + 1
		}

		if target == nums[i]+nums[j] {
			if nums[i] == nums[j] {
				return m[nums[i]]
			}
			return []int{m[nums[j]][0], m[nums[i]][0]}
		}

		if target > nums[i]+nums[j] {
			j++
		} else {
			i++
			j = i + 1
		}
	}

	return []int{}
}

/* 3. https://leetcode.com/problems/squares-of-a-sorted-array/
* O(n), O(n)
 */

func SortedSquares(nums []int) []int {
	i, j := 0, len(nums)-1
	resIdx := len(nums) - 1
	res := make([]int, len(nums))
	for i <= j {
		if square(nums[j]) > square(nums[i]) {
			res[resIdx] = square(nums[j])
			j--
		} else {
			res[resIdx] = square(nums[i])
			i++
		}
		resIdx--
	}

	return res
}

func square(a int) int {
	return a * a
}

/* 4. https://leetcode.com/problems/valid-palindrome/
* O(n), O(1)
 */
func IsPalindrome(s string) bool {
	i, j := 0, len(s)-1

	for i < j {
		if !isValid(rune(s[i])) {
			i++
			continue
		}

		if !isValid(rune(s[j])) {
			j--
			continue
		}

		if !equalStringsCaseInsensitive(rune(s[i]), rune(s[j])) {
			return false
		}

		i++
		j--
	}

	return true
}

func equalStringsCaseInsensitive(a, b rune) bool {
	return strings.EqualFold(string(a), string(b))
}

func isValid(r rune) bool {
	return unicode.IsDigit(r) || unicode.IsLetter(r)
}

/* 5. https://leetcode.com/problems/boats-to-save-people/
* O(nlogn), O(1)
 */
func NumRescueBoats(people []int, limit int) int {
	sort.Ints(people)

	count := 0

	for i, j := 0, len(people)-1; i <= j; j-- {
		if i == j {
			count++
			continue
		}

		if people[i]+people[j] <= limit {
			i++
		}

		count++
	}

	return count
}
