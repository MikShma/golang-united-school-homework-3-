package homework

import "sort"

func sortMapValues(input map[int]string) (result []string) {
	//Place your code here
	k := []int{}
	for i := range input {
		k = append(k, i)
	}
	sort.Ints(k)
	for i := range k {
		result = append(result, input[i])
	}
	return
}
