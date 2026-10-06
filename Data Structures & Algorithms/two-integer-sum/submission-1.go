func twoSum(nums []int, target int) []int {
	mp:= make(map[int]int)
	for index,a := range nums {
		mp[a] = index
	}

	for index,a := range nums {
		ost := target - a
		if indexMap , ok := mp[ost]; ok && indexMap != index {
			return []int{index, mp[ost]}
		}
	}
	return []int{0}
}
