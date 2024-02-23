package arrays

func containsDuplicate(nums []int) bool {
	hashset := make(map[int]bool)

	for _, n := range nums {
		if _, ok := hashset[n]; ok {
			return true
		}
		hashset[n] = true
	}
	return false
}
