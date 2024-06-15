package main

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
)

func dailyTemperatures(temperatures []int) []int {
	res := make([]int, len(temperatures))
	tmp := make([]int, 0)
	for index, v := range temperatures {
		if len(tmp) == 0 || v <= temperatures[tmp[len(tmp)-1]] {
			tmp = append(tmp, index)
		} else {
			for len(tmp) > 0 && v > temperatures[tmp[len(tmp)-1]] {
				pre := tmp[len(tmp)-1]
				tmp = tmp[:len(tmp)-1]
				res[pre] = index - pre
			}
			tmp = append(tmp, index)
		}
	}
	return res
}

func main() {
	data := []int{73, 74, 75, 71, 69, 72, 76, 73}
	res := dailyTemperatures(data)
	fmt.Println(res)
	t1()
}

func t1() {
	str := "{\"k1\": \"v1\", \"k2\": \"v2\"}"
	m := make(map[string]string)
	jsoniter.UnmarshalFromString(str, &m)
	fmt.Println(m)
}
