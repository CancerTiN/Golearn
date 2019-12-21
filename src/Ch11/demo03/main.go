package main

import "fmt"

type A struct {
	Name string
	age  int
}

type B struct {
	Name  string
	Score float64
}

type C struct {
	A
	B
	// Name string
}

type D struct {
	a A // Famous structure
}

type Goods struct {
	Name  string
	Price float64
}

type Brand struct {
	Name    string
	Address string
}

type TV struct {
	Goods
	Brand
}

type TVa struct {
	*Goods
	*Brand
}

type Monster struct {
	Name string
	Age  int
}

type E struct {
	Monster
	int
	n int
}

func main() {
	var c C
	// c.Name = "tom" // ambiguous selector c.Name
	c.A.Name = "tom"
	fmt.Println("c =", c)

	var d D
	// d.name = "jack" // d.name undefined (type D has no field or method name)
	d.a.Name = "jack"
	fmt.Println("d =", d)

	tv1 := TV{
		Goods{"television001", 5000.99},
		Brand{"Haier", "Shandong"},
	}

	fmt.Println("tv1.Goods.Name =", tv1.Goods.Name)
	fmt.Println("tv1.Price =", tv1.Price)

	tv2 := TV{
		Goods{
			Price: 5000.99,
			Name:  "television002",
		},
		Brand{
			"Sharp",
			"Beijing",
		},
	}
	fmt.Println("tv1 =", tv1)
	fmt.Println("tv2 =", tv2)

	tv3 := TVa{
		&Goods{"television003", 5000.99},
		&Brand{"Skyworth", "Henan"},
	}
	fmt.Println("tv3 =", tv3)
	fmt.Println("tv3.Goods =", *tv3.Goods, "tv3.Brand =", *tv3.Brand)

	tv4 := TVa{
		&Goods{
			Price: 5000.99,
			Name:  "television004",
		},
		&Brand{
			"Changhong",
			"Sichuan",
		},
	}
	fmt.Println("tv4.Goods =", *tv4.Goods, "tv4.Brand =", *tv4.Brand)

	var e E
	e.Name = "Vixen"
	e.Age = 300
	e.int = 200
	fmt.Println("e =", e)

}
