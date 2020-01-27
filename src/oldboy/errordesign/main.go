package main

func main() {

}

func document01() {
	/*
		关于函数返回值的错误设计，下面说法正确：
		1)如果失败原因只有一个，则返回bool
		2)如果失败原因超过一个，则返回error
		3)如果重试几次可以避免失败，则不要立即返回bool或error
	*/
}
