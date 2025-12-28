package aho_corasick

import (
	"bytes"
	"fmt"
)

// getFailureFunction Knuth-Morris-Pratt algorithm
func getFailureFunction(str string) []int {
	n := len(str)
	f := make([]int, n)
	t := 0

	for s := 1; s < n; s++ {
		for t > 0 && str[s] != str[t] {
			t = f[t-1]
		}
		if str[s] == str[t] {
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
