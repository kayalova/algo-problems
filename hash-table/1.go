package hashtable

import (
	"sort"
	"strings"
)

/* 1. https://leetcode.com/problems/contains-duplicate/description/
* O(n), O(n)
 */
func ContainsDuplicate(nums []int) bool {
	m := map[int]int{}

	for _, n := range nums {
		if _, ok := m[n]; ok {
			return true
		}

		m[n]++
	}

	return false
}

/* 2. https://leetcode.com/problems/valid-anagram/
* O(n), O(n)
 */
func IsAnagram(s string, t string) bool {
	sMap := make(map[rune]int, len(s))
	tMap := make(map[rune]int, len(s))

	for _, r := range s {
		sMap[r]++
	}

	for _, r := range t {
		tMap[r]++
	}

	for _, r := range s {
		v, ok := tMap[r]
		if !ok {
			return false
		}

		if v != sMap[r] {
			return false
		}
	}

	for _, r := range t {
		v, ok := sMap[r]
		if !ok {
			return false
		}

		if v != tMap[r] {
			return false
		}
	}

	return true
}

/* 3. https://leetcode.com/problems/group-anagrams/
* O(n)?, O(n)
 */
func GroupAnagrams(strs []string) [][]string {
	m := map[string][]string{}

	for _, str := range strs {
		sortedStr := sortStr(str)
		m[sortedStr] = append(m[sortedStr], str)
	}

	res := [][]string{}

	for _, v := range m {
		res = append(res, v)
	}

	return res
}

func sortStr(s string) string {
	runes := []rune(s)
	sort.Slice(runes, func(a, b int) bool {
		return runes[a] < runes[b]
	})

	return string(runes)
}

/* 4. https://leetcode.com/problems/longest-substring-without-repeating-characters/
O(n^2), O(n)
*/

func LengthOfLongestSubstring(s string) int {
	current := []rune{}
	longest := []rune{}

	for _, r := range s {
		if idx := strings.IndexRune(string(current), r); idx != -1 {
			if len(current) > len(longest) {
				longest = current
			}

			current = current[idx+1:]
		}

		current = append(current, r)
	}

	if len(current) > len(longest) {
		longest = current
	}

	return len(longest)
}

/* 5. https://leetcode.com/problems/top-k-frequent-elements/
* O(n), O(n)
* в описании было следующее замечание
* 	Follow up: Your algorithm's time complexity must be better than O(n log n), where n is the array's size.
* Не понимаю как сделать быстрее т.к нужна сортировка
 */
func TopKFrequent(nums []int, k int) []int {
	m := make(map[int]int)

	for _, n := range nums {
		m[n]++
	}

	reversedM := map[int][]int{} // ключ - это частота с которой встречаются элементы, а значения - сами элементы
	for el, freq := range m {
		if _, ok := reversedM[freq]; !ok {
			reversedM[freq] = []int{el}
		} else {
			reversedM[freq] = append(reversedM[freq], el)
		}
	}

	freqList := make([]int, 0) // создаем список на основе ключей и затем сортируем
	for fr := range reversedM {
		freqList = append(freqList, fr)
	}
	sort.Ints(freqList)

	res := make([]int, 0)
	for i := len(freqList) - 1; i >= 0 && len(res) < k; i-- {
		v := reversedM[freqList[i]]
		res = append(res, v...)
	}

	if len(res) > k {
		return res[:k]
	}

	return res
}
