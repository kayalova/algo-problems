package twopointers

import (
	"math"
	"strings"
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
* Не понимаю как сделать через 2 указателя за O(n), были мысли отсортировать чтобы было хотя nlogn, но тогда идексы поменяются
* O(n^2), O(1)
 */
func TwoSum(nums []int, target int) []int {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[j] == target-nums[i] {
				return []int{i, j}
			}
		}
	}

	return []int{}
}

// O(n), O(n)
func TwoSum2(nums []int, target int) []int {
	m := map[int]int{}

	for pos, n := range nums {
		if v, ok := m[target-n]; ok {
			return []int{pos, v}
		}

		m[n] = pos
	}

	return []int{}
}

/* 3. https://leetcode.com/problems/squares-of-a-sorted-array/
* O(n), O(n)
 */

func SortedSquares(nums []int) []int {
	i, j := 0, len(nums)-1
	res := make([]int, len(nums))
	for i <= j {
		if square(nums[j]) > square(nums[i]) {
			res[j-i] = square(nums[j])
			j--
		} else {
			res[j-i] = square(nums[i])
			i++
		}
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
	m := createAndFillMap()
	i, j := 0, len(s)-1

	for i < j {
		if !isValid(m, string(rune(s[i]))) {
			i++
			continue
		}

		if !isValid(m, string(rune(s[j]))) {
			j--
			continue
		}

		if !strings.EqualFold(string(rune(s[i])), string(rune(s[j]))) { // можно ли как то сравнить символы, чтобы не делать столько преобразований
			return false
		}

		i++
		j--
	}

	return true
}

func isValid(m map[string]struct{}, s string) bool {
	_, ok := m[strings.ToLower(s)]
	return ok
}

func createAndFillMap() map[string]struct{} {
	m := map[string]struct{}{}
	for i := 'a'; i <= 'z'; i++ {
		m[string(i)] = struct{}{}
	}

	for i := '0'; i <= '9'; i++ {
		m[string(i)] = struct{}{}
	}

	return m
}

/* 5. https://acmp.ru/index.asp?main=task&id_task=869
* O(n), O(1)
 */
func MinimalCanoe(cap int, people []int) int {
	s := 0
	for _, p := range people {
		s += p
	}

	return int(math.Ceil(float64(s) / float64(cap)))
}
