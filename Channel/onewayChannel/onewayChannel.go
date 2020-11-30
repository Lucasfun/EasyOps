package main

import "fmt"

func counter(out chan<- int) { //只发送型Channel
	for x := 0;x <= 10;x ++ {
		out <- x
	}
	close(out)
}

func squarer(out chan<- int,in <-chan int) {
	for v := range in{
		out <- v * v
	}
	close(out)
}

func printer(in <-chan int) {
	for v := range in{
		fmt.Println(v)
	}
}

func main() {
	naturals := make(chan int)
	squarers := make(chan int)
	go counter(naturals) //naturals隐式转化为 chan<- int
	//任何双向Channel转单向都是隐形的，但是单向无论如何无法转换为双向
	go squarer(squarers,naturals)
	printer(squarers)

}
