package integers

func Sum(ints []int) int {
	var sum int
	for i := 0; i < len(ints); i++ {
		sum += ints[i]
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
