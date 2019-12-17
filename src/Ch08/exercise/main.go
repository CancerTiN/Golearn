package main

import "fmt"

func main() {
	var scores [3][5]float64
	for i := 0; i < len(scores); i++ {
		for j := 0; j < len(scores[i]); j++ {
			fmt.Printf("Please enter the grade of student %d in class %d: ", j+1, i+1)
			fmt.Scanln(&scores[i][j])
		}
	}
	fmt.Println(scores)

	totalSum := 0.0
	totalSize := 0.0
	for i := 0; i < len(scores); i++ {
		sum := 0.0
		size := float64(len(scores[i]))
		for j := 0; j < len(scores[i]); j++ {
			sum += scores[i][j]
		}
		totalSum += sum
		totalSize += size
		fmt.Printf("The class %d has a total score of %v and an average score of %v\n", i+1, sum, sum/size)
	}
	fmt.Printf("The total score for all classes is %v, with an average score of %v\n",
		totalSum, totalSum/totalSize)
}
