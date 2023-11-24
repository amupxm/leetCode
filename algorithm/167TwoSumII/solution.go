package twosumii

import "fmt"

func twoSum(numbers []int, target int) []int {
	hMap := make(map[int]int) // number[index]
	for i, number := range numbers {
		if index, ok := hMap[target-number]; ok {
			fmt.Println(hMap)
			return []int{index + 1, i + 1}
		}
		hMap[number] = i
	}
	return []int{}
}
