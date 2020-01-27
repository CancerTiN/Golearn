package main

func main() {

}

func document01() {
	/*
		关于同步锁，下面说法正确：
		1)当一个goroutine获得了Mutex后，其它goroutine就只能乖乖地等待，除非该goroutine释放这个Mutex
		2)RWMutex在读锁占用的情况下，会阻止写，但不阻止读
		3)RWMutex在写锁占用的情况下，会阻止任何其它goroutine进来，整个锁相当于由该goroutine独占
	*/
}
