package twopointers

import "strconv"

func Compress(chars []byte) int {
	indx := 0
	for i := 0; i < len(chars); {
		j := i
		for ; j < len(chars) && chars[j] == chars[i]; j++ {
		}

		chars[indx] = chars[i]
		indx++

		if j-i > 1 {
			if j-i > 9 {
				for _, r := range strconv.Itoa(j - i) {
					chars[indx] = byte(r)
					indx++
				}
				indx--
			} else {
				chars[indx] = []byte(strconv.Itoa(j - i))[0]
			}
			indx++
		}
		i = j
	}

	chars = chars[:indx]

	return indx
}
