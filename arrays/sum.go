package main

func Sum(array []int) int {
	sum := 0
	for _, number := range array {
		sum += number
	}
	return sum
}
