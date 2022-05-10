package homework

import "sort"

func sortMapValues(input map[int]string) (result []string) {
	//Place your code here
	var (
		num []int
	)

	for k, _ := range input {
		num = append(num, k)
	}
	sort.Ints(num)

	for _, v := range num {
		result = append(result, input[v])
	}
	return result
}
