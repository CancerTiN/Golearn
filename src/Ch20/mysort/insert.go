package main

// 0 end [10, 34, 19, 99, 80]
// 1 end [10, 34, 19, 99, 80]
// 2 --- [10, 34, 34, 99, 80]
// 2 end [10, 19, 34, 99, 80]
// 3 end [10, 19, 34, 99, 80]
// 4 --- [10, 19, 34, 99, 99]
// 4 end [10, 19, 34, 80, 99]

func Insert(slice []int) {
	for i := 1; i < len(slice); i++ {
		insertIndex := i
		insertValue := slice[insertIndex]
		for j := i - 1; j >= 0; j-- {
			if slice[j] > insertValue {
				insertIndex = j
				slice[j+1] = slice[j]
			}
		}
		if insertIndex != i {
			slice[insertIndex] = insertValue
		}
	}
}
