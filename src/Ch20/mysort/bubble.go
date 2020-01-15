package main

func Bubble(slice []int) {
	var flag bool
	for i := 1; i < len(slice); i++ {
		flag = false
		for j := 0; j < len(slice)-i; j++ {
			if slice[j] > slice[j+1] {
				slice[j], slice[j+1] = slice[j+1], slice[j]
				flag = true
			}
		}
		if !flag {
			break
		}
	}
}
