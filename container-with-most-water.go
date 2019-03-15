func maxArea(height []int) int {    max := 0
	i:= 0
	j := len(height)-1
	var tmp = 0
	for i<j{		
		if height[i]>height[j]{
			tmp = height[j] * (j - i)
			j--
		}else{
			tmp = height[i] * (j - i)
			i++
		}
		if tmp > max {
			max = tmp
		}
	}
	return max
}
