package odd

func Count(numbers []int) int {
	xor := 0
	for _, num := range numbers {
		xor ^= num
	}
	return xor
}
