package integers

func Sum(ints []int) int {
	var sum int
	for i := 0; i < len(ints); i++ {
		sum += ints[i]
	}
	return sum
}

func SumAll(numbersToSum ...[]int) (sums []int) {
	lengthOfNumbers := len(numbersToSum)
	sums = make([]int, lengthOfNumbers)

	for i, numbers := range numbersToSum {
		sums[i] = Sum(numbers)
	}

	return
}
