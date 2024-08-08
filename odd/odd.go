package odd

func Count(numbers []int) int {
	countMap := make(map[int]int)
	for _, num := range numbers {
		countMap[num]++
	}
	for num, count := range countMap {
		if count%2 != 0 {
			return num
		}
	}
	return 0
}
