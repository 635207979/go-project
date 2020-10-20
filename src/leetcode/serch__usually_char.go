package main

import (
	"fmt"
	"strings"
)

//func compareMap(m1 *map[string]int, m2 map[string]int) {
//	for key, _ := range *m1 {
//		if m2[key] == 0 {
//			(*m1)[key] = 0
//		} else {
//			if (*m1)[key] >= m2[key] {
//				(*m1)[key] = m2[key]
//			}
//		}
//	}
//}
//
//func commonChars(A []string) []string {
//
//	result := make(map[string]int)
//	var res []string
//	for _, char := range A[0] {
//		result[string(char)] ++
//	}
//	for _, str := range A[1:] {
//		tmp := make(map[string]int)
//		for _, char := range str {
//			tmp[string(char)]++
//		}
//		compareMap(&result, tmp)
//	}
//	for key, value := range result {
//		if value >= 0 {
//			for i := 0; i < value; i++ {
//				res = append(res, key)
//			}
//
//		}
//	}
//	return res
//}
//func commonChars(A []string) []string {
//	min_str := A[0]
//	min_len := len(min_str)
//	for i := 1; i < len(A); i++ {
//		if len(A[i]) < min_len {
//			min_len = len(A[i])
//			min_str = A[i]
//		}
//	}
//	count := [26]int{}
//	for _, j := range min_str {
//		count[j-97] += 1
//	}
//	for _, k := range A {
//		tmp := [26]int{}
//		for _, v := range k {
//			tmp[v-97] += 1
//		}
//		for l := 0; l < 26; l++ {
//			if count[l] != 0 {
//				count[l] = min(count[l], tmp[l])
//			}
//		}
//	}
//	res := []string{}
//	for i := 0; i < 26; i++ {
//		for j := 0; j < count[i]; j++ {
//			res = append(res, string(i+97))
//		}
//	}
//	return res
//}

//func min(a int, b int) int {
//	if a < b {
//		return a
//	} else {
//		return b
//	}
//}

func commonChars(A []string) []string {
	res := []string{}
	mid := map[string]int{}
	for _, v := range A[0] {
		if _, ok := mid[string(v)]; !ok {
			mid[string(v)] = 1
		} else {
			mid[string(v)]++
		}
	}
	res = []string{}
	for j := 1; j < len(A); j++ {
		for k, v := range mid {
			if strings.Count(A[j], k) <= v {
				mid[k] = strings.Count(A[j], k)
			}
		}
	}
	for k, v := range mid {
		for i := 1; i <= v; i++ {
			res = append(res, k)
		}
	}
	return res

}
func main() {
	A := []string{"bella", "label", "roller"}
	res := commonChars(A)
	fmt.Printf("%v", res)

}
