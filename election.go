package main

func vote(a int, b int, c int) int {
	if (a == b) || (a == c) {
		return a
	}
	return b
}
