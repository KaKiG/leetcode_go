package main

import "fmt"

func main() {
	s := ")((((((f"
	fmt.Println(removeInvalidParentheses(s))
}

func removeInvalidParentheses(s string) []string {
	lCount, rCount := 0, 0
	res := []string{}
	for i := range s {
		if s[i] == '(' {
			lCount++
		}
		if s[i] == ')' {
			if lCount > 0 {
				lCount--
			} else {
				rCount++
			}
		}
	}
	// counting the inbalance and using dfs (brute force) to solve
	solver(s, 0, lCount, rCount, &res)
	if len(res) == 0 {
		return []string{""}
	}
	return res
}

func solver(s string, start, l, r int, res *[]string) {
	// if s is valid and the string is valid. We need both conditions because )((), is also valid in count but not in string.
	if l == 0 && r == 0 && isValid(s) {
		*res = append(*res, s)
		return
	}

	// for every character in s, remove it the check rest of the string, and do dfs
	// when s[i] == s[i-1], it means there are two continuous same characters which will cause repetition if searched twice
	// begin with i = start is also to avoid duplication
	for i := start; i < len(s); i++ {
		if i != 0 && s[i] == s[i-1] {
			continue
		}
		newStr := s[:i] + s[i+1:]
		// here we only do substraction, no adding
		// because we only want the minimum removement, more removements will result in negative number of l and r
		// it will not be added to the solution set, exactly what we want
		if s[i] == '(' && l > 0 {
			solver(newStr, i, l-1, r, res)
		}
		if s[i] == ')' && r > 0 {
			solver(newStr, i, l, r-1, res)

		}

	}
}

func isValid(s string) bool {
	l := 0
	for i := range s {
		if s[i] == '(' {
			l++
		}
		if s[i] == ')' {
			l--
		}
		if l < 0 {
			return false
		}
	}
	return l == 0
}
