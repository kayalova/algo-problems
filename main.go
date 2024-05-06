package main

import (
	"fmt"
)

func main() {
	str := "ate"
	a := str[1:1]
	fmt.Println(len(a))
	// fmt.Println(strings.IndexRune(str, 'q'))

	// strRune := []rune(str)
	// fmt.Println(strRune)

	// strList := strings.Split(str, "")
	// sort.Strings(strList)
	// fmt.Println(strList)
	// // sort.Sort(strRune)
	// sort.Slice(strRune, func(i, j int) bool {
	// 	return strRune[i] < strRune[j]
	// })
	// fmt.Println(string(strRune))
	// sort.Ints(strRune)
	// list := []int{1, 2, 4, 21, 3}
	// l2 := []int{1}
	// l2 = append(l2, list...)
	// sort.Ints(list)
	// fmt.Println(list)
	// reversedM := map[int][]int{}
	// reversedM[1] = []int{1, 2}
	// fmt.Println(reversedM)
}
