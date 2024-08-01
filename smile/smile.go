package smile

import "regexp"

func Count(arr []string) int {
	re := regexp.MustCompile(`[:;][-~]?[)D]`)

	count := 0
	for _, face := range arr {
		if re.MatchString(face) {
			count++
		}
	}

	return count
}
