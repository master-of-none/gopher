package arrays

func twoSum(nums []int, target int) []int {
	hashmap := make(map[int]int)

	for i, n := range nums {
		diff := target - n
		if value, ok := hashmap[diff]; ok {
			return []int{i, value}
		}
		hashmap[n] = i
	}
	return []int{}
}
