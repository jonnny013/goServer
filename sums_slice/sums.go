package main

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(arrays ...[]int) []int {
	var result []int
	for _, array := range arrays {
		result = append(result, Sum(array))
	}
	return result
}

func sumAllTails(arrays ...[]int) []int {
	var result []int
	for _, array := range arrays {
		if len(array) == 0 {
			result = append(result, 0)
		} else {
			headRemoved := array[1:]
			result = append(result, Sum(headRemoved))
		}
	}
	return result
}
