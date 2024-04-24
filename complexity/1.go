package complexity

import (
	"os"
	"strconv"
	"strings"
)

/* 1. https://acmp.ru/index.asp?main=task&id_task=21
* В отделе работают 3 сотрудника, которые получают заработную плату в рублях.
* Требуется определить: на сколько зарплата самого высокооплачиваемого из них
* отличается от самого низкооплачиваемого
*
* O(n), O(1)
 */
func SalaryDiff(nums ...int) int {
	if len(nums) == 0 {
		return 0
	}

	min, max := nums[0], nums[0]

	for _, n := range nums {
		if n < min {
			min = n
		}

		if n > max {
			max = n
		}
	}

	return max - min
}

/*
 2. https://acmp.ru/index.asp?main=task&id_task=131

* Во входном файле find_oldest_man.txt в первой строке задано натуральное число N – количество жильцов (N ≤ 100).
* В последующих N строках располагается информация о всех жильцах: каждая строка содержит
* два целых числа: V и S – возраст и пол человека (1 ≤ V ≤ 100, S – 0 или 1).
* Мужскому полу соответствует значение S=1, а женскому – S=0.
* Требуется найти номер самого старшего жителя мужского пола.
*
* O(n), O(1)
*/
func FindOldestMan(fileName string) (int, error) {
	file, err := os.ReadFile(fileName)
	if err != nil {
		return 0, err
	}

	lines := strings.Split(string(file), "\n")

	peopleCount, err := strconv.Atoi(lines[0]) // почему можно иметь 2 переменых err на одном уровне
	if err != nil {
		return 0, err
	}

	var person []string
	maxAge := -1
	maxAgePosition := -1
	for p := 1; p <= peopleCount; p++ {
		person = strings.Split(lines[p], " ")
		if person[1] != "1" {
			continue
		}

		age, err := strconv.Atoi(person[0]) // можно ли как-то person[0] привести к uint8
		if err != nil {
			return 0, err
		}

		if maxAge < age {
			maxAge = age
			maxAgePosition = p
		}
	}

	return maxAgePosition, nil
}

/* 3. https://acmp.ru/index.asp?main=task&id_task=43
* Требуется найти самую длинную непрерывную цепочку нулей в последовательности нулей и единиц.
*
* O(n), O(1)
 */
func LongestZeroChain(s string) int {
	maxZeroCount, currentZeroCount := 0, 0
	TARGET := rune('0')

	for _, r := range s { // почему r = 48, это ascii код?
		if r == TARGET {
			currentZeroCount++
		} else {
			if maxZeroCount < currentZeroCount {
				maxZeroCount = currentZeroCount
			}
			currentZeroCount = 0
		}
	}

	return maxZeroCount
}

/* 4. https://acmp.ru/index.asp?main=task&id_task=46
* Выведите в выходной файл округленное до n знаков после десятичной точки число E.
* В данной задаче будем считать, что число Е в точности равно 2.7182818284590452353602875.
* O(n), O(n)
 */

func RoundEulersNumber(n int) string {
	e := "2.7182818284590452353602875"

	if n == 0 {
		return "3"
	}

	rounded := make([]byte, 0, len(e)-2) // сюда помещаем n+1 знаков после запятой
	for i := 0; i < 2+n+1; i++ {
		if i <= 1 {
			continue
		}

		rounded = append(rounded, e[i])
	}

	intPartStr := "2."
	toRoundStr := string(rounded[:len(rounded)-1]) // оставляем n знаков после запятой
	last := string(rounded[len(rounded)-1])        // на осоновании последней цифры решаем делать ли округление в большую сторону
	if last < "5" {
		return intPartStr + toRoundStr
	}

	toPlusNumber, _ := strconv.ParseInt(toRoundStr, 10, 64)
	roundedNumber := int(toPlusNumber + 1)

	return intPartStr + strconv.Itoa(roundedNumber)
}

/* 5. https://leetcode.com/problems/two-sum/
* O(n), O(n)
*
 */
func TwoSum(nums []int, target int) []int {
	s := map[int]int{}
	for pos, n := range nums {
		if i, ok := s[target-n]; ok {
			return []int{i, pos}
		}

		s[n] = pos
	}

	return []int{}
}

/* 6. https://leetcode.com/problems/remove-duplicates-from-sorted-array/
* O(n), O(1)
 */
func RemoveDuplicates(arr []int) int {
	i, k := 0, 1

	for j := 0; j < len(arr); j++ {
		if arr[j] == arr[i] {
			continue
		}
		arr[i+1], arr[j] = arr[j], arr[i+1]
		k++
		i++
	}

	return k
}

/*
7. https://leetcode.com/problems/remove-element/description/
* O(n), O(1)
*/
func RemoveElement(arr []int, val int) int {
	i, j, k := 0, 0, 0
	for j < len(arr) {
		if arr[j] == val {
			j++
			continue
		}

		arr[i], arr[j] = arr[j], arr[i]
		i++
		j++
		k++
	}

	return k
}

/* 8. https://leetcode.com/problems/search-insert-position/
* O(logn), O(1)
 */
func SearchInsert(nums []int, target int) int {
	left, right, midPos := 0, len(nums)-1, len(nums)/2
	for left <= right {
		if nums[left] > target {
			return left
		}

		if nums[right] < target {
			return right + 1
		}

		midPos = (left + right) / 2
		if nums[midPos] == target {
			return midPos
		}
		if nums[midPos] > target {
			right = midPos - 1
		} else {
			left = midPos + 1
		}
	}

	return midPos + 1
}
