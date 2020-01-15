package main

func Select(slice []int) {
	for i := 0; i < len(slice); i++ {
		m := i
		for j := i + 1; j < len(slice); j++ {
			if slice[m] > slice[j] {
				m = j
			}
		}
		if m != i {
			slice[i], slice[m] = slice[m], slice[i]
		}
	}
}
