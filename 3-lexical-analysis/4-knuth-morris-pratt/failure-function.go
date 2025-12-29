package knuth_morris_pratt

import (
	"bytes"
	"fmt"
)

func search(str string, pattern string) bool {
	f := getFailureFunction(pattern)
	n := len(pattern)
	s := 0

	for i := 0; i < len(str); i++ {
		for s > 0 && str[i] != pattern[s] {
			s = f[s-1]
		}
		if str[i] == pattern[s] {
			s++
		}
		if s == n {
			return true
		}
	}

	return false
}

func getFailureFunction(pattern string) []int {
	n := len(pattern)
	f := make([]int, n)
	t := 0

	for s := 1; s < n; s++ {
		for t > 0 && pattern[s] != pattern[t] {
			t = f[t-1]
		}
		if pattern[s] == pattern[t] {
			t++
		}
		f[s] = t
	}

	return f
}

func printFailureFunction(str string, f []int) {
	bBuilder := bytes.NewBufferString("b: ")
	sBuilder := bytes.NewBufferString("s: ")
	fBuilder := bytes.NewBufferString("f: ")

	for i := 0; i < len(str); i++ {
		bBuilder.WriteString(fmt.Sprintf("%s ", string(str[i])))
		sBuilder.WriteString(fmt.Sprintf("%v ", i+1))
		fBuilder.WriteString(fmt.Sprintf("%v ", f[i]))
	}

	fmt.Println(bBuilder.String())
	fmt.Println(sBuilder.String())
	fmt.Println(fBuilder.String())
}

func getFibonacciString(n int) string {
	if n < 1 {
		return ""
	}

	f1, f2 := "b", "a"
	if n == 1 {
		return f1
	}

	for i := 2; i < n; i++ {
		f1, f2 = f2, f2+f1
	}

	return f2
}
