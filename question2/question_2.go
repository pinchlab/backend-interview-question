package main

import "fmt"

func getNoZeroIntegers(n int) []int {
	// TODO: Write your code here
	// Given an integer n, return a list of two integers [a, b] where:
	// - a and b are No-Zero integers.
	// - a + b = n

	return []int{}
}

type Case struct {
	Input  int
	Output []int
}

func getCases() []Case {
	return []Case{
		{2, []int{1, 1}},
		{11, []int{2, 9}},
		{10000, []int{1, 9999}},
		{69, []int{1, 68}},
		{1010, []int{11, 999}},
		{100002349000, []int{2349111, 99999999889}},
	}
}

// CompareTwoSlices checks if two slices are equal
func CompareTwoSlices(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func main() {
	cases := getCases()
	for _, c := range cases {
		res := getNoZeroIntegers(c.Input)
		fmt.Printf("Input: %v, Output: %v, Expected: %v, Pass: %v\n", c.Input, res, c.Output, CompareTwoSlices(res, c.Output))
	}
}
