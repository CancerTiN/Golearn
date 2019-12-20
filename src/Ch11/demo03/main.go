package main

import "fmt"

type Student struct {
	Name  string
	Age   int
	Score int
}

func (s *Student) ShowInfo() {
	fmt.Printf("Name: %v; Age: %v; Score: %v.\n", s.Name, s.Age, s.Score)
}

func (s *Student) SetScore(score int) {
	s.Score = score
}

func (s *Student) GetSum(n1 int, n2 int) int {
	return n1 + n2
}

type Pupil struct {
	Student
}

func (p *Pupil) testing() {
	fmt.Println("Elementary school student are taking an exam ...")
}

type Graduate struct {
	Student
}

func (g *Graduate) testing() {
	fmt.Println("College student are taking an exam ...")
}

func main() {
	//var pupil Pupil
	pupil := &Pupil{}
	pupil.Student.Name = "tom"
	pupil.Student.Age = 8
	pupil.testing()
	pupil.Student.SetScore(70)
	pupil.Student.ShowInfo()
	fmt.Println(pupil.Student.GetSum(1, 2))

	graduate := &Graduate{}
	graduate.Student.Name = "mary"
	graduate.Student.Age = 28
	graduate.testing()
	graduate.Student.SetScore(90)
	graduate.ShowInfo()
	fmt.Println(graduate.Student.GetSum(10, 20))
}
