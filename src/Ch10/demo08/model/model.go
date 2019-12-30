package model

type person struct {
	Name   string
	age    int
	salary float64
}

func NewPerson() *person {
	return &person{
		Name:   "WuSong",
		age:    30,
		salary: 8000,
	}
}

func (p *person) SetSalary(s float64) {
	p.salary = s
}
