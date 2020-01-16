package main

func Quick(slice []int, left, right int) {
	if right < 1 {
		return
	}
	pivot := slice[(left+right)/2]
	l := left
	r := right
	for l < r {
		for slice[l] < pivot {
			l++
		}
		for pivot < slice[r] {
			r--
		}
		if l >= r {
			break
		}
		slice[l], slice[r] = slice[r], slice[l]
		if slice[l] == pivot {
			r--
		}
		if pivot == slice[r] {
			l++
		}
	}
	if l == r {
		l++
		r--
	}
	if l < right {
		Quick(slice, l, right)
	}
	if left < r {
		Quick(slice, left, r)
	}
}
