package main

import "fmt"

func Sum(numbers []int) (sum int) {
	for i, number := range numbers {
		fmt.Println(i)
		sum += number
	}
	return sum
}
func SumAll(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}
	return sums
}
func SumTails(numbersToSum ...[]int) []int {
	var tails []int
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			tails = append(tails, 0)
		} else {
			tail := numbers[1:]
			tails = append(tails, Sum(tail))
		}
	}
	return tails
}
