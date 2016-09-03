package main

import (
	"fmt"
)

// Prefix Function
func pi(s []byte) []int {
	n := len(s)
	pi := make([]int, n)
	pi[0] = 0
	for i := 1; i < n; i++ {
		j := pi[i-1]
		if s[i] == s[j] {
			pi[i] = j + 1
		} else {
			j = pi[j]
			for j > 0 {
				if s[i] == s[j] {
					pi[i] = j + 1
					break
				} else {
					j = pi[j]
				}
			}
		}
	}
	return pi
}

//check if long contain small
func contain(long []byte, small []byte) bool {
	pi := pi(small)
	i := 0
	j := 0
	for i < len(long) {
		if long[i] == small[j] {
			if j == len(small)-1 {
				return true
			}
			i += 1
			j += 1
		} else {
			for j > 0 && long[i] != small[j] {
				j = pi[j-1]
			}
			if j == 0 && long[i] != small[j] {
				i += 1
			}
		}
	}
	return false
}

func main() {
	s := "ababababca"
	pi := pi([]byte(s))
	for _, c := range pi {
		fmt.Print(c, " ")
	}
	fmt.Println()
	ss := "asdfaowabbabbababababcaabbababbc"
	if contain([]byte(ss), []byte(s)) {
		fmt.Println(ss, " contains ", s)
	} else {
		fmt.Println(ss, " doesn't contain ", s)
	}
}
