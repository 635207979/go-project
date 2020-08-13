package integers

func Sum(ints []int) int {
	var sum int
	for i := 0; i < len(ints); i++ {
		sum += ints[i]
	}
	return sum
}
