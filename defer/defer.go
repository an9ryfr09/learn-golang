package deleyexec

func deferfunc(nums []int) []int {
	var newnums = make([]int, 5)
	for k, n := range nums {
		defer func(key, num int) {
			newnums[key] = num
		}(k, n)
	}
	return newnums
}
