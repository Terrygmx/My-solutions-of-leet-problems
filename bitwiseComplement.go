func bitwiseComplement(N int) int {
	if N == 0 {
		return 1
	} else {
		c := 0
		p := N
		for N > 0 {
			N >>= 1
			c++
		}
		return (1 << uint(c)) - 1 - p
	}
}
