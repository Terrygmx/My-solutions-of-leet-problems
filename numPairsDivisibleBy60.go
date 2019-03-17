func numPairsDivisibleBy60(time []int) int {
	size := len(time)
	if size < 2 {
		return 0
	}
	r := 0
	for i := 0; i < size-1; i++ {
		for j := i + 1; j < size; j++ {
			if (time[i]+time[j])%60 == 0 {
				r++
			}
		}
	}
	return r
}
