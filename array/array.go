package array

func Sum(arr []int) int {
	sum := 0
	for _, number := range arr {
		sum += number
	}
	return sum
}

func SumMultiArrs(arrs ...[]int) (sums []int) {
	countArrs := len(arrs)
	sums = make([]int, countArrs)
	for i, arr := range arrs {
		sums[i] = Sum(arr)
	}
	return
}
