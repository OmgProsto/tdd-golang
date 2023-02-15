package golang_tdd

func Sum(numbers []int) int {
	var res int
	for _, v := range numbers {
		res += v
	}

	return res
}

func SumAll(nToSum ...[]int) []int {
	length := len(nToSum)
	sums := make([]int, length)

	for i, n := range nToSum {
		sums[i] = Sum(n)
	}

	return sums
}

func SumAllTails(nToSum ...[]int) []int {
	var sums []int
	for _, v := range nToSum {
		if len(v) == 0 {
			sums = append(sums, 0)
		} else {
			tail := v[1:]
			sums = append(sums, Sum(tail))
		}
	}
	return sums
}
