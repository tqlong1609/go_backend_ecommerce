package tests

import "fmt"

func AddOne(a int) int {
	return a + 1
}

func BasicCoverage() int {
	if 1 == 2 {
		fmt.Println("Fail")
	}
	return 1
}
