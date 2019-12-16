package main

import "fmt"

func main() {
	classSum := 0.0
	classNum := 3
	for j := 1; j <= classNum; j++ {
		classMean := classGrade()
		classSum += classMean
	}
	totalMean := classSum / float64(classNum)
	fmt.Printf("The average grade for the %d classes is %.2f\n", classNum, totalMean)
}

func classGrade() float64 {
	studentSum := 0.0
	studentNum := 5
	for i := 1; i <= studentNum; i++ {
		var score float64
		fmt.Printf("Please enter the grade of the student %d: ", i)
		fmt.Scanln(&score)
		studentSum += score
	}
	studentMean := studentSum / float64(studentNum)
	fmt.Printf("The average score for this class is %.2f\n", studentMean)
	return studentMean
}
