package permutations

import "fmt"

func Permutations(s string) []string {
	chars := []rune(s)
	res := make([]string, 0)
	generateStr(chars, 0, &res)

	checkDup := make(map[string]bool)
	uniqRes := make([]string, 0)
	for _, s := range res {
		if !checkDup[s] {
			checkDup[s] = true
			uniqRes = append(uniqRes, s)
		}
	}
	return uniqRes
}

func generateStr(chars []rune, start int, result *[]string) {
	if start == len(chars) {
		*result = append(*result, string(chars))
		return
	}

	for i := start; i < len(chars); i++ {
		fmt.Println(i, start, string(chars))
		chars[start], chars[i] = chars[i], chars[start]
		generateStr(chars, start+1, result)
		chars[start], chars[i] = chars[i], chars[start]
	}
}
