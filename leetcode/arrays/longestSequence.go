package arrays

func longestConsecutive(nums []int) int {
	hashset := make(map[int]bool)

	for _, n := range nums {
		hashset[n] = true
	}

	res := 0

	for _, n := range nums {
		if _, ok := hashset[n-1]; !ok {
			longest := 0

			for hashset[n+longest] {
				longest++
			}
			res = max(res, longest)
		}
	}
	return res

}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
