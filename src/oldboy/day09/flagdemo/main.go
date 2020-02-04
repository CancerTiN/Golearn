package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {
	name := flag.String("name", "tin", "请输入名字")
	age := flag.Int("age", 27, "请输入年龄")
	married := flag.Bool("married", false, "是否已婚")
	delay := flag.Duration("delay", time.Second, "婚龄")
	var double float64
	flag.Float64Var(&double, "double", 0, "请输入浮点数")
	flag.Parse()
	fmt.Println("name:", *name)
	fmt.Println("age:", *age)
	fmt.Println("married:", *married)
	fmt.Println("delay:", *delay)
	fmt.Println("double:", double)
	fmt.Println(flag.Args())
	fmt.Println(flag.NFlag())
}
