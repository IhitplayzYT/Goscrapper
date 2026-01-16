package main

import (
	"fmt"
	"os"
	"strings"
)

type i8 int8
type i16 int16
type i32 int32
type i64 int64
type s8 uint8
type s16 uint16
type s32 uint32
type s64 uint64
type f32 float32
type f64 float64

const True bool = true
const False bool = false

// Placeholder
var website_list []string = []string{"a", "b", "c"}

//  abc
// cabd
// 0123
// i = 3,cur_idx = 2

func match_longest_any(to_be_found, to_be_searched string) (int, int) {
	lf, ls := len(to_be_found), len(to_be_searched)
	mx_len, ed_idx := 0, -1
	dp := make([][]int, lf+1)
	for i := 0; i <= lf; i++ {
		dp[i] = make([]int, ls+1)
	}

	for i := 1; i <= lf; i++ {
		for j := 1; j <= ls; j++ {
			if to_be_found[i-1] == to_be_searched[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
				if dp[i][j] > mx_len {
					mx_len = dp[i][j]
					ed_idx = j
				}
			} else {
				dp[i][j] = dp[i-1][j-1]
			}
		}
	}
	return ed_idx - mx_len, mx_len
}

// :TODO
func match_longest(to_be_found, to_be_searched string) (int, int) {

}

// CAN FAIL WITH -1
func match_full(to_be_found, to_be_searched string) int {
	ls := len(to_be_searched)
	lf := len(to_be_found)
	idx := 0
	for i := 0; i < ls; i++ {
		if to_be_found[idx] == to_be_searched[i] {
			idx += 1
		}
		if idx == lf {
			return i - lf + 1
		}
	}

	return -1
}

func usage(str string) {
	fmt.Println(str, " [-d] <KEYWORD>")
	return
}

func print_map(s map[string][]string) {
	for k, v := range s {
		fmt.Print("[%v]:", k)
		z := len(k)
		for i := 0; i < len(v); i++ {
			fmt.Println(strings.Repeat(" ", z+2), v[i])
		}
	}
}

func search(str string, flag bool) map[string][]string {
	ret := make(map[string][]string)
	for _, v := range website_list {

	}

	return ret
}

func main() {
	args := os.Args
	keyword, deep_flag := "", false
	l := len(args)
	if l < 2 || l > 3 {
		usage(args[0])
		return
	} else if l == 2 {
		keyword = args[1]
	} else {
		if args[1] != "-d" {
			usage(args[0])
			return
		}
		deep_flag = true
		keyword = args[2]
	}

	ret := search(keyword, deep_flag)
	print_map(ret)
	return
}
