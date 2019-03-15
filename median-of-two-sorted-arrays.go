func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	i := 0
	j := 0
	len1 := len(nums1)
	len2 := len(nums2)
    count := len1 + len2
	if count % 2 == 1 {
		tmp := 0
		half := (count-1)/2
		for i+j<= half{if i < len1{
				if j < len2{
					if nums1[i] < nums2[j]{
						tmp = nums1[i]
						
						i++
					}else{
						tmp = nums2[j]
						
						j++
					}
				}else{
					tmp = nums1[i]
					
					i++
				}
								
			}else{
				tmp = nums2[j]
				j++
			}	
	//		fmt.Println("tmp:",tmp)	
		}	
		return float64(tmp)
	}else{
		tmp := 0
		tmp1 := 0
		half := count/2
		for i+j<= half{
			if i+j == half{
				tmp1 = tmp
//				fmt.Println("tmp1:",tmp1)
			}
			if i < len1{
				if j < len2{
					if nums1[i] < nums2[j]{
						tmp = nums1[i]
						
						i++
					}else{
						tmp = nums2[j]
						
						j++
					}
				}else{
					tmp = nums1[i]
					
					i++
				}
								
			}else{
				tmp = nums2[j]
				j++
			}	
//			fmt.Println("tmp:",tmp)		
		}	
		return (float64(tmp+tmp1))/2
	}	
}
