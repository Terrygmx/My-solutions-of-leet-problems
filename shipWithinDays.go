func shipWithinDays(weights []int, D int) int {
	if D == len(weights) {
		return getMax(weights)
	} else if D == 1 {
		return getSum(weights)
	}
	average := getSum(weights) / D
	max := getMax(weights)
	var pivot int
	if average > max {
		pivot = average
	} else {
		pivot = max
	}
	fmt.Println("pivot:", pivot)
	flag := false
	for !flag {
		target := []int{}
		t := 0
		for _, v := range weights {
			//fmt.Println("i:", i)
			if t+v <= pivot {
				t += v
				//		fmt.Println("t:", t)
			} else {
				target = append(target, t)
				//		fmt.Println("target:", target)
				t = v
			}
		}
		if t > 0 {
			target = append(target, t)
		}
		//	fmt.Println("lenth of target", len(target))
		if len(target) <= D {
			flag = true
			return pivot
		} else {
			pivot++
			//	fmt.Println("pivot new:", pivot)
		}
	}
	return pivot

}

func getMax(w []int) int {
	m := 0
	for _, v := range w {
		if v > m {
			m = v
		}
	}
	return m
}
func getSum(w []int) int {
	m := 0
	for _, v := range w {

			m += v

	}
	return m
}
