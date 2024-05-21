package twopointers

func MinWindow(s string, t string) string {
	/*
		1. Начиная с len(t) символа символа проходимся по строке s и проверяем содержится ли в ней t
		2. Как только содержится - запоминаем строку, начинаем срезать слева пока не перестанет содержаться
		3. Снова идем вправо и повторяем алгоритм пока не дойдем до конца строки
	*/
	if len(t) > len(s) {
		return ""

	}

	minimumW := ""
	flag := false
	mapTStr := getStringBytesMap(t)
	for left, right := 0, len(t)-1; right < len(s); {

		if flag { // понимаем, что нужно срезать слева
			substr := s[left+1 : right+1]
			if contains(substr, mapTStr) { // если содержится, значит все еще срезаем
				if len(minimumW) > len(substr) {
					minimumW = substr
				}
			} else {
				flag = false // перестаем срезать и идем искать новую длину подстроки содержащую t
			}
			left++
		} else {
			substr := s[left : right+1] // подстрока с которой работаем
			if contains(substr, mapTStr) {
				if len(minimumW) == 0 {
					minimumW = substr
				}

				if len(minimumW) > len(substr) {
					minimumW = substr
				}

				flag = true
			} else {
				right++
			}
		}
	}

	return minimumW
}

func contains(s string, t map[byte]int) bool {
	sMap := getStringBytesMap(s)
	for b, tCount := range t {
		if sCount, ok := sMap[b]; !ok || (sCount < tCount) {
			return false
		}
	}

	return true
}

func getStringBytesMap(s string) map[byte]int {
	m := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		m[s[i]]++
	}

	return m
}
