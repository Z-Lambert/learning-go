package main

func SumAll(numbersToSum ...[]int) int {
	sumOfAll := 0
	for _, numbers := range numbersToSum {
		sumOfAll += Sum(numbers)
	}
	return sumOfAll
}
